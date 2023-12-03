package main

import (
	"bufio"
	"fmt"
	"net"

	"github.com/chzyer/readline"
)

func main() {
	conn, _ := net.Dial("tcp", "localhost:8080")
	defer conn.Close()

	messages := make(chan string)

	go readMessages(conn, messages)
	go readUserInput(conn)

	for message := range messages {
		fmt.Println("Получено сообщение:", message)
	}
}

func readMessages(conn net.Conn, messages chan string) {
	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			close(messages)
			return
		}
		messages <- message
	}
}

func readUserInput(conn net.Conn) {
	rl, _ := readline.New("Введите сообщение: ")
	defer rl.Close()

	for {
		line, _ := rl.Readline()
		fmt.Fprintf(conn, line+"\n")
	}
}
