package client

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"

	"github.com/Zmey56/wisdom-tcp-server/pkg/pow"
)

func StartClient(conn net.Conn) {
	message, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Println("Received:", message)

	parts := strings.Fields(message)
	if len(parts) < 3 {
		fmt.Println("Invalid challenge format received")
		return
	}
	seed := parts[2]

	proof := solvePoW(seed)

	_, _ = conn.Write([]byte(proof + "\n"))
	time.Sleep(100 * time.Millisecond) // Небольшая задержка для синхронизации

	response, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Println("Server response:", response)
}

func solvePoW(seed string) string {
	var proof int
	powImpl := pow.PoWImpl{}

	for {
		proofStr := strconv.Itoa(proof)
		if powImpl.VerifyPoW(seed, proofStr) {
			fmt.Printf("Proof found: %s\n", proofStr)
			return proofStr
		}
		proof++
	}
}
