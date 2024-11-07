package server

import (
	"bufio"
	"context"
	"net"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"

	"github.com/Zmey56/wisdom-tcp-server/pkg/pow"
)

func TestHandleConnection(t *testing.T) {
	ctx := context.Background()

	req := testcontainers.ContainerRequest{
		Image:        "word-of-wisdom-server",
		ExposedPorts: []string{"8080/tcp"},
		Env: map[string]string{
			"MOCK_WISDOM_QUOTE": "Mocked wisdom quote for testing.",
		},
		WaitingFor: wait.ForListeningPort("8080/tcp"),
	}
	serverContainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	assert.NoError(t, err)
	defer serverContainer.Terminate(ctx)

	host, err := serverContainer.Host(ctx)
	assert.NoError(t, err)
	port, err := serverContainer.MappedPort(ctx, "8080")
	assert.NoError(t, err)

	serverAddr := net.JoinHostPort(host, port.Port())
	t.Log("Connecting to server at:", serverAddr)

	conn, err := net.Dial("tcp", serverAddr)
	assert.NoError(t, err)
	defer conn.Close()

	reader := bufio.NewReader(conn)
	challengeMsg, err := reader.ReadString('\n')
	assert.NoError(t, err)
	t.Log("Received challenge from server:", challengeMsg)
	assert.Contains(t, challengeMsg, "Solve PoW:")

	seed := extractSeed(challengeMsg)
	t.Log("Extracted seed:", seed)

	proof := solvePoW(seed)
	t.Log("Generated proof:", proof)

	_, err = conn.Write([]byte(proof + "\n"))
	assert.NoError(t, err)

	response, err := reader.ReadString('\n')
	assert.NoError(t, err)
	t.Log("Received response from server:", response)
	assert.Equal(t, "Here is your wisdom: Mocked wisdom quote for testing.\n", response)
}

func extractSeed(msg string) string {
	parts := strings.Fields(msg)
	if len(parts) < 3 {
		return ""
	}
	return parts[2]
}

func solvePoW(seed string) string {
	var proof int
	powImpl := pow.PoWImpl{}

	for {
		proofStr := strconv.Itoa(proof)
		if powImpl.VerifyPoW(seed, proofStr) {
			return proofStr
		}
		proof++
	}
}
