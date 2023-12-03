package server

import (
	"fmt"
	"net"
	"os"

	"github.com/egor-muindor/cipher-talk-go/internal/repository"
	"github.com/egor-muindor/cipher-talk-go/internal/server/handler"
)

type Server struct {
	clients  repository.Clients
	handler  handler.ConnectionHandler
	listener net.Listener
}

// New is a constructor function that returns a new Server.
func New(address string, clients repository.Clients, connectionHandler handler.ConnectionHandler) Server {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Println("Ошибка при создании сервера:", err)
		os.Exit(1)
	}
	return Server{
		clients:  clients,
		handler:  connectionHandler,
		listener: listener,
	}
}

// Start is a method of Server that starts the server.
func (s Server) Start() {
	fmt.Println("Сервер запущен и слушает на адресе", s.listener.Addr().String())
	defer s.listener.Close()
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			fmt.Println("Ошибка при подключении:", err)
			continue
		}
		c := s.clients.Add(conn)
		go s.handler.HandleConnection(c)
	}
}
