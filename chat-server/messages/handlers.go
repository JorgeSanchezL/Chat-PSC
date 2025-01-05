package messages

import (
	"encoding/json"
	"fmt"
	"log"
	"server/users"

	"github.com/nats-io/nats.go"
)

func HandleConnect(nc *nats.Conn, storage *Storage, user string) {
	users.AddUser(user)

	sendStoredMessages(nc, storage, user)

	fmt.Printf("Client connected: %s\n", user)
}

func sendStoredMessages(nc *nats.Conn, storage *Storage, user string) {
	storedMessages := storage.GetMessages(user)
	for _, message := range storedMessages {
		msgBytes, err := json.Marshal(message)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Sending old message to %s: %s\n", fmt.Sprintf("%s.old", message.Recipient), message.Content)
		err = nc.Publish(fmt.Sprintf("%s.old", message.Recipient), msgBytes)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func CreateMessageSubscription(nc *nats.Conn, storage *Storage, username string) (*nats.Subscription, error) {
	return nc.Subscribe(fmt.Sprintf("%s.new", username), func(m *nats.Msg) {
		var msg Message
		err := json.Unmarshal(m.Data, &msg)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("The following message was sent by %s to %s: %s\n", msg.Sender, msg.Recipient, msg.Content)
		storage.Store(msg)
	})
}
