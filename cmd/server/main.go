package main

import (
	"log"
	"net"

	"github.com/Zmey56/wisdom-tcp-server/internal/server"
	"github.com/Zmey56/wisdom-tcp-server/pkg/pow"
	"github.com/Zmey56/wisdom-tcp-server/pkg/wisdom"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
	defer listener.Close()
	log.Println("Server is listening on port 8080...")

	// Создаем сервер с реальными реализациями PoW и Wisdom
	srv := server.NewServer(pow.PoWImpl{}, wisdom.WisdomImpl{})

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go srv.HandleConnection(conn)
	}
}
