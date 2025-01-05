package messages

type Message struct {
	Time      string `json:"time"`
	Sender    string `json:"sender"`
	Recipient string `json:"recipient"`
	Content   string `json:"content"`
}

type Connection struct {
	Sender string `json:"sender"`
}
