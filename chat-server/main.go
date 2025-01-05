package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"server/messages"
	"server/users"

	"github.com/nats-io/nats.go"
)

func main() {
	var natsURL string
	if url := os.Getenv("NATS_URL"); url != "" {
		natsURL = url
	} else {
		natsURL = nats.DefaultURL
	}

	nc, err := nats.Connect(natsURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	fmt.Println("Server is running on NATS")

	storage := messages.NewStorage()

	_, err = nc.Subscribe("connections", func(m *nats.Msg) {
		var msg messages.Connection
		err := json.Unmarshal(m.Data, &msg)
		if err != nil {
			log.Fatal(err)
		}

		messages.HandleConnect(nc, storage, msg.Sender)
		messages.CreateMessageSubscription(nc, storage, msg.Sender)
	})

	nc.Subscribe("connections-srv", func(m *nats.Msg) {
		UserList := users.GetUsers()

		response, err := json.Marshal(UserList)
		if err != nil {
			log.Printf("Failed to marshal users: %v", err)
			return
		}

		m.Respond(response)
	})

	if err != nil {
		log.Fatal(err)
	}

	select {} // Keep the server running
}
