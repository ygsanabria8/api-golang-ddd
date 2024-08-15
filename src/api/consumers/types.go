package consumers

import (
	"api.ddd/src/api/server"
)

type CreatedUserEventMessage struct {
	logger *server.Logger
}

func NewCreatedUserEventMessage(
	logger *server.Logger,
) *CreatedUserEventMessage {
	return &CreatedUserEventMessage{
		logger: logger,
	}
}

type UpdatedUserEventMessage struct {
	logger *server.Logger
}

func NewUpdatedUserEventMessage(
	logger *server.Logger,
) *UpdatedUserEventMessage {
	return &UpdatedUserEventMessage{
		logger: logger,
	}
}

type DeletedUserEventMessage struct {
	logger *server.Logger
}

func NewDeletedUserEventMessage(
	logger *server.Logger,
) *DeletedUserEventMessage {
	return &DeletedUserEventMessage{
		logger: logger,
	}
}
