package users

import (
	"sync"
)

var (
	userList = []string{}
	mu       sync.Mutex
)

func AddUser(clientName string) {
	mu.Lock()
	defer mu.Unlock()

	if clientName == "" {
		return
	}

	if userList == nil {
		userList = []string{}
	}

	for _, c := range userList {
		if c == clientName {
			return
		}
	}

	userList = append(userList, clientName)
}

func GetUsers() []string {
	mu.Lock()
	defer mu.Unlock()
	return append([]string(nil), userList...)
}
