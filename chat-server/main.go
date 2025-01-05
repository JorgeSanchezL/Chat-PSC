package main

import (
	"fmt"
	"log"
	"server/messages"

	zmq "github.com/pebbe/zmq4"
)

func main() {
	routerSocket, err := zmq.NewSocket(zmq.ROUTER)
	if err != nil {
		log.Fatal(err)
	}
	defer routerSocket.Close()

	err = routerSocket.Bind("tcp://*:5555")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Server is running on tcp://*:5555")

	storage := messages.NewStorage()

	for {
		msg, err := routerSocket.RecvMessage(0)
		if err != nil {
			log.Fatal(err)
		}

		go messages.Process(routerSocket, storage, msg)
	}
}
