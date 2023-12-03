package repository

import (
	"net"

	storage "github.com/egor-muindor/cipher-talk-go/internal/deps"
)

// TODO Зарефакторить весь файл, разделить на хранилище и dto

type Client struct {
	Conn net.Conn
	ID   string
}

type Clients struct {
	mutex storage.Mutex
	list  map[string]Client
}

// NewClients is a constructor function that returns a new Clients struct.
func NewClients(mu storage.Mutex) Clients {
	return Clients{
		list:  make(map[string]Client),
		mutex: mu,
	}
}

// Add adds a client to the Clients struct.
func (cs *Clients) Add(conn net.Conn) Client {
	cs.mutex.Lock()
	id := conn.RemoteAddr().String()
	c := Client{
		ID:   id,
		Conn: conn,
	}
	cs.list[id] = c
	cs.mutex.Unlock()

	return c
}

// Remove removes a client from the Clients struct.
func (cs *Clients) Remove(id string) {
	cs.mutex.Lock()
	delete(cs.list, id)
	cs.mutex.Unlock()
}

// GetAll returns a map of all clients.
func (cs *Clients) GetAll() map[string]Client {
	return cs.list
}
