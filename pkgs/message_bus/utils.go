package message_bus

import (
	"github.com/google/uuid"
)

// NewMessageBusEvent create an event based on an interface and topic
func NewMessageBusEvent(message interface{}) *Event {
	return &Event{
		Id:      uuid.New().String(),
		Message: message,
	}
}
