package server

import (
	"bufio"
	"fmt"
	"net"
	"strings"

	"github.com/Zmey56/wisdom-tcp-server/pkg/pow"
	"github.com/Zmey56/wisdom-tcp-server/pkg/wisdom"
)

type Server struct {
	PoW    pow.PoW
	Wisdom wisdom.Wisdom
}

func NewServer(pow pow.PoW, wisdom wisdom.Wisdom) *Server {
	return &Server{
		PoW:    pow,
		Wisdom: wisdom,
	}
}

func (s *Server) HandleConnection(conn net.Conn) {
	defer conn.Close()

	seed, prefix := s.PoW.GenerateChallenge()
	_, _ = conn.Write([]byte("Solve PoW: " + seed + " with prefix " + prefix + "\n"))

	reader := bufio.NewReader(conn)
	proof, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading proof:", err)
		return
	}

	proof = strings.TrimSpace(proof)
	fmt.Printf("Received proof: %s\n", proof)

	if s.PoW.VerifyPoW(seed, proof) {
		quote := s.Wisdom.GetRandomQuote()
		_, _ = conn.Write([]byte("Here is your wisdom: " + quote + "\n"))
	} else {
		_, _ = conn.Write([]byte("Invalid PoW.\n"))
	}
}
