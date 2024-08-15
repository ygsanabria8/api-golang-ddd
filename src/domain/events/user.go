package events

import (
	"api.ddd/pkgs/message_bus"
	"api.ddd/src/domain/aggregates"
)

func NewCreatedUserEvent(user *aggregates.User) *message_bus.Event {
	return message_bus.NewMessageBusEvent(&CreatedUserEvent{User: user})
}

func NewUpdatedUserEvent(user *aggregates.User) *message_bus.Event {
	return message_bus.NewMessageBusEvent(&UpdatedUserEvent{User: user})
}

func NewDeletedUserEvent(userId string) *message_bus.Event {
	return message_bus.NewMessageBusEvent(&DeletedUserEvent{UserId: userId})
}
