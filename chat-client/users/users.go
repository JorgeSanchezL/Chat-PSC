package users

import (
	"client/flags"
	"encoding/json"
	"fmt"
	"log"
	"sync"

	"github.com/nats-io/nats.go"
)

var (
	userList []string
	mu       sync.Mutex
)

func SyncServerUsers(nc *nats.Conn) {
	// Request existing users from the server, if we dont use a server, the request will simply timeout
	msg, err := nc.Request("connections-srv", nil, nats.DefaultTimeout)
	if err != nil {
		if err == nats.ErrTimeout {
			fmt.Println("Request timed out, no server response. The chat will still work as it if was ran without a server.")
			flags.SetServerRunning(false)
		} else if err == nats.ErrNoResponders {
			fmt.Println("No servers found. The chat will still work with limitations. If you are running the chat without a server, you can ignore this message.")
			flags.SetServerRunning(false)
		} else {
			log.Fatal(err)
		}
	} else {
		flags.SetServerRunning(true)
		var users []string
		err = json.Unmarshal(msg.Data, &users)
		if err != nil {
			log.Fatal(err)
		}

		mu.Lock()
		userList = users
		mu.Unlock()
	}
}

func AddUser(user string) {
	mu.Lock()
	defer mu.Unlock()
	userList = append(userList, user)
}

func GetUsers() []string {
	mu.Lock()
	defer mu.Unlock()
	return append([]string(nil), userList...)
}
