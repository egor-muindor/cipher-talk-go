package handler

import (
	"bufio"
	"fmt"
	"io"
	"net"

	"github.com/egor-muindor/cipher-talk-go/internal/repository"
)

type ConnectionHandler struct {
	clients *repository.Clients // clients is a pointer to a Clients struct.
}

// NewConnectionHandler is a constructor function that returns a new ConnectionHandler.
func NewConnectionHandler(clients *repository.Clients) ConnectionHandler {
	return ConnectionHandler{
		clients: clients,
	}
}

// HandleConnection is a method of ConnectionHandler that handles a client connection.
// It reads messages from the client and broadcasts them to other clients.
// If an error occurs or the client closes the connection, it removes the client from the Clients struct.
func (h ConnectionHandler) HandleConnection(c repository.Client) {
	defer func(Conn net.Conn) {
		err := Conn.Close()
		if err != nil {
			fmt.Println("Ошибка при закрытии соединения:", err)
		}
	}(c.Conn)
	reader := bufio.NewReader(c.Conn)
	fmt.Println("Клиент открыл соединение")
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Println("Клиент закрыл соединение")
			} else {
				fmt.Println("Ошибка при чтении данных:", err)
			}
			h.clients.Remove(c.ID)
			return
		}
		fmt.Print("Получено сообщение:", message)
		h.Broadcast(c.ID, message)
	}
}

// Broadcast is a method of ConnectionHandler that sends a message to all clients except the sender.
// It takes the sender's ID and the message as arguments.
func (h ConnectionHandler) Broadcast(senderID, message string) {
	clients := h.clients.GetAll()
	for id, c := range clients {
		if id != senderID {
			go func(c repository.Client, message string) {
				_, err := c.Conn.Write([]byte(message))
				if err != nil {
					fmt.Println("Ошибка при отправке сообщения:", err)
				}
			}(c, message)
		}
	}
}
