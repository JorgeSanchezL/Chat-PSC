package messages

import (
	"fmt"
	"log"
	"server/users"

	zmq "github.com/pebbe/zmq4"
)

func Process(socket *zmq.Socket, storage *Storage, msg []string) {
	// The message always contains the message type, client identity, and an empty delimiter
	// If it is a message for another client, it will also contain the recipient's name and the message content
	fmt.Printf("Processing message: %s\n", msg)

	sender := msg[0]
	messageType := msg[1]

	switch messageType {
	case TYPECONNECT:
		handleConnect(socket, storage, sender)
	case TYPEMESSAGE:
		messageRecipient := msg[3]
		messageContent := msg[4]
		message := Message{Sender: sender, Recipient: messageRecipient, Content: messageContent}
		handleClientMessage(socket, storage, message)
	default:
		log.Printf("Unknown message type: %s\n", messageType)
	}
}

func handleConnect(socket *zmq.Socket, storage *Storage, user string) {
	users.AddUser(user)

	message := Message{
		Sender:    "server",
		Recipient: user,
		Content:   "Welcome to the chat server! Press 's' and 'ENTER' to send a message.",
	}

	sendMessage(socket, message)

	storedMessages := storage.GetMessages(user)
	for _, message := range storedMessages {
		sendMessage(socket, message)
	}

	fmt.Printf("Client connected: %s\n", user)
}

func handleClientMessage(socket *zmq.Socket, storage *Storage, message Message) {
	fmt.Printf("Received message from client %s for %s: %s\n", message.Sender, message.Recipient, message.Content)

	storage.Store(message)

	sendMessage(socket, message)
}

func sendMessage(routerSocket *zmq.Socket, message Message) {
	if !users.UserExists(message.Recipient) {
		fmt.Printf("Recipient not found: %s\n", message.Recipient)
		routerSocket.SendMessage(message.Sender, "server", DELIMITER, "The recipient does not exist, so the message could not be sent.")
		return
	}

	_, err := routerSocket.SendMessage(message.Recipient, message.Sender, DELIMITER, message.Content)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s sent a message to %s: %s\n", message.Sender, message.Recipient, message.Content)
}
