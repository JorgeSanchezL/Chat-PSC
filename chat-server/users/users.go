package users

import (
	"sync"
)

var (
	clientMap = []string{}
	mu        sync.Mutex
)

func AddUser(clientName string) {
	mu.Lock()
	defer mu.Unlock()

	if clientName == "" {
		return
	}

	if clientMap == nil {
		clientMap = []string{}
	}

	for _, c := range clientMap {
		if c == clientName {
			return
		}
	}

	clientMap = append(clientMap, clientName)
}

func UserExists(clientName string) bool {
	mu.Lock()
	defer mu.Unlock()

	for _, c := range clientMap {
		if c == clientName {
			return true
		}
	}

	return false
}
