package main

import (
	"bufio"
	"client/messages"
	"client/users"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/nats-io/nats.go"
)

func main() {
	var natsURL string
	if url := os.Getenv("NATS_URL"); url != "" {
		natsURL = url
	} else {
		natsURL = nats.DefaultURL
	}

	if len(os.Args) < 2 {
		fmt.Println("Usage: ./client <username>")
		return
	}
	username := os.Args[1]

	nc, err := nats.Connect(natsURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	_, err = nc.Subscribe(fmt.Sprintf("%s.*", username), func(m *nats.Msg) {
		var msg messages.Message
		err := json.Unmarshal(m.Data, &msg)
		if err != nil {
			log.Fatal(err)
		}

		if msg.Recipient != username {
			log.Fatal("Received a message that wasn't meant for this user")
		}

		fmt.Printf("%s\t%s: %s\n", msg.Time, msg.Sender, msg.Content)
	})

	if err != nil {
		log.Fatal(err)
	}

	connection := messages.Connection{Sender: username}
	connectionMessage, err := json.Marshal(connection)
	if err != nil {
		log.Fatal(err)
	}

	err = nc.Publish("connections", connectionMessage)
	if err != nil {
		log.Fatal(err)
	}

	// Gets the users that are already connected to the server. If we have no server, this won't do anything
	users.SyncServerUsers(nc)

	fmt.Printf("> Welcome to the chat server, %s! Press 's' and 'ENTER' to send a message.\n", username)

	_, err = nc.Subscribe("connections", func(m *nats.Msg) {
		var msg messages.Connection
		err := json.Unmarshal(m.Data, &msg)
		if err != nil {
			log.Fatal(err)
		}

		// Receiving this message means a user has logged in, so we can print it
		fmt.Printf("> %s has logged in\n", msg.Sender)

		users.AddUser(msg.Sender)
	})

	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "s":
			messages.SendMessage(nc, username, reader)
		default:
			fmt.Println("> Unknown command. Press 's' and 'ENTER' to send a message.")
		}
	}
}
