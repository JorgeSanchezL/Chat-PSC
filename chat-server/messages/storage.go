package messages

import (
	"sync"
)

type Storage struct {
	messages map[string][]Message
	mu       sync.Mutex
}

func NewStorage() *Storage {
	return &Storage{
		messages: make(map[string][]Message),
	}
}

func (s *Storage) Store(msg Message) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.messages[msg.Recipient] = append(s.messages[msg.Recipient], msg)
}

func (s *Storage) GetMessages(recipient string) []Message {
	s.mu.Lock()
	defer s.mu.Unlock()
	queue := s.messages[recipient]
	copyQueue := make([]Message, len(queue))
	copy(copyQueue, queue)
	return copyQueue
}
