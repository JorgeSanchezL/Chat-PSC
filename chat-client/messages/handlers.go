package messages

import (
	"bufio"
	"client/flags"
	"client/users"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/nats-io/nats.go"
)

func SendMessage(nc *nats.Conn, username string, reader *bufio.Reader) {
	fmt.Printf("\033[1A\033[K") // Clear the last line

	fmt.Print("> Enter recipient: ")
	recipient, _ := reader.ReadString('\n')
	recipient = strings.TrimSpace(recipient)

	fmt.Print("> Enter message: ")
	content, _ := reader.ReadString('\n')
	content = strings.TrimSpace(content)

	fmt.Printf("\033[1A\033[K") // Clear the recipient line
	fmt.Printf("\033[1A\033[K") // Clear the message line

	if flags.GetServerRunning() && username != recipient {
		usersList := users.GetUsers()
		userExists := false
		for _, user := range usersList {
			if user == recipient {
				userExists = true
				break
			}
		}

		if !userExists {
			fmt.Printf("> User %s does not exist\n", recipient)
			return
		}
	}

	message := Message{
		Time:      time.Now().Format("02-01-2006 15:04:05"),
		Recipient: recipient,
		Sender:    username,
		Content:   content,
	}
	var msgBytes []byte
	msgBytes, err := json.Marshal(message)
	if err != nil {
		log.Fatal(err)
	}

	err = nc.Publish(fmt.Sprintf("%s.new", recipient), msgBytes)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("> You sent the following message to %s: %s\n", recipient, message.Content)
}
