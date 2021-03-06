package pubsub

import "github.com/goccy/go-json"

// Message represents a single message with payload bound to a particular topic
type Message struct {
	// Topic (channel in terms of redis)
	Topic string `json:"topic"`
	// Payload (on some decode stages might be represented as base64 string)
	Payload []byte `json:"payload"`
}

func (m *Message) MarshalBinary() (data []byte, err error) {
	return json.Marshal(m)
}
