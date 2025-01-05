package main

import (
	"bufio"
	"client/messages"
	"fmt"
	"log"
	"os"
	"strings"

	zmq "github.com/pebbe/zmq4"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: ./client <username>")
		return
	}
	username := os.Args[1]

	socket, err := zmq.NewSocket(zmq.DEALER)
	if err != nil {
		log.Fatal(err)
	}
	defer socket.Close()

	socket.SetIdentity(username)

	err = socket.Connect("tcp://localhost:5555")
	if err != nil {
		log.Fatal(err)
	}

	_, err = socket.SendMessage(messages.TYPECONNECT)
	if err != nil {
		log.Fatal(err)
	}

	go messages.ReceiveMessages(socket, username)

	reader := bufio.NewReader(os.Stdin)
	for {
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "s":
			messages.SendMessage(socket, username, reader)
		default:
			fmt.Println("> Unknown command. Press 's' and 'ENTER' to send a message.")
		}
	}
}
