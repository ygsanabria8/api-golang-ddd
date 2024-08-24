package utils

import (
	"github.com/IBM/sarama"
)

type IEvent interface {
	GetMessage() []byte
	GetTopic() string
	GetHeader() map[string]string
}

// Event basic structure that contains the topic to send event and event's data
type Event struct {
	Message []byte
	Topic   string
	Headers []*sarama.RecordHeader
}

// GetMessage return message bytes
func (e *Event) GetMessage() []byte {
	return e.Message
}

// GetTopic return the topic's name
func (e *Event) GetTopic() string {
	return e.Topic
}

// GetHeader return map from each header into the message
func (e *Event) GetHeader() map[string]string {
	headers := make(map[string]string)
	for _, header := range e.Headers {
		headers[string(header.Key)] = string(header.Value)
	}
	return headers
}
