package flags

import "sync"

var (
	serverRunning bool
	mu            sync.Mutex
)

func SetServerRunning(value bool) {
	mu.Lock()
	defer mu.Unlock()
	serverRunning = value
}

func GetServerRunning() bool {
	mu.Lock()
	defer mu.Unlock()
	return serverRunning
}
