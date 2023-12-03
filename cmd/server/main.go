package main

import (
	"sync"

	"github.com/egor-muindor/cipher-talk-go/internal/repository"
	"github.com/egor-muindor/cipher-talk-go/internal/server"
	"github.com/egor-muindor/cipher-talk-go/internal/server/handler"
)

func main() {
	mu := sync.Mutex{}
	clients := repository.NewClients(&mu)
	connectionHandler := handler.NewConnectionHandler(&clients)
	server.New(":8080", clients, connectionHandler).Start()
}
