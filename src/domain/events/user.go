package events

import (
	"api.ddd/pkgs/message_bus/utils"
	"api.ddd/src/domain/aggregates"
)

func NewCreatedUserEvent(user *aggregates.User) *utils.Event {
	return utils.NewMessageBusEvent(&CreatedUserEvent{User: user})
}

func NewUpdatedUserEvent(user *aggregates.User) *utils.Event {
	return utils.NewMessageBusEvent(&UpdatedUserEvent{User: user})
}

func NewDeletedUserEvent(userId string) *utils.Event {
	return utils.NewMessageBusEvent(&DeletedUserEvent{UserId: userId})
}
