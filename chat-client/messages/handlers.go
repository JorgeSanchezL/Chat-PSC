package messages

import (
	"bufio"
	"fmt"
	"log"
	"strings"

	zmq "github.com/pebbe/zmq4"
)

func ReceiveMessages(socket *zmq.Socket, username string) {
	for {
		msg, err := socket.RecvMessage(0)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("\033[1G\033[K") // Clear the current line
		fmt.Printf("%s: %s\n", msg[0], msg[2])
		fmt.Printf("> ")
	}
}

func SendMessage(socket *zmq.Socket, username string, reader *bufio.Reader) {
	fmt.Printf("\033[1A\033[K") // Clear the last line
	fmt.Print("> Enter recipient: ")
	recipient, _ := reader.ReadString('\n')
	recipient = strings.TrimSpace(recipient)

	fmt.Print("> Enter message: ")
	message, _ := reader.ReadString('\n')
	message = strings.TrimSpace(message)

	fmt.Printf("\033[1A\033[K") // Clear the recipient line
	fmt.Printf("\033[1A\033[K") // Clear the message line

	_, err := socket.SendMessage(TYPEMESSAGE, DELIMITER, recipient, message)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("> You sent the following message to %s: %s\n", recipient, message)
}
