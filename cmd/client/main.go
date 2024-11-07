package main

import (
	"log"
	"net"

	"github.com/Zmey56/wisdom-tcp-server/internal/client"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	client.StartClient(conn)
}
