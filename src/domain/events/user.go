package events

import (
	"api.ddd/src/domain/aggregates"
)

func NewCreatedUserEvent(user *aggregates.User) *CreatedUserEvent {
	return &CreatedUserEvent{User: user}
}

func NewUpdatedUserEvent(user *aggregates.User) *UpdatedUserEvent {
	return &UpdatedUserEvent{User: user}
}

func NewDeletedUserEvent(userId string) *DeletedUserEvent {
	return &DeletedUserEvent{UserId: userId}
}
