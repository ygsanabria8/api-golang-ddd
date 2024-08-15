package consumers

import (
	"api.ddd/pkgs/message_bus"
	"api.ddd/src/domain/events"
)

func (e *CreatedUserEventMessage) OnMessage(event *message_bus.Event) {
	e.logger.Infof("Invoked CreatedUserEventMessage Consumer")

	var message *events.CreatedUserEvent
	err := event.GetMessage(&message)
	if err != nil {
		e.logger.Infow("Error Getting Message", err)
		return
	}

	e.logger.Infow("CreatedUserEventMessage", message.User)
}

func (e *UpdatedUserEventMessage) OnMessage(event *message_bus.Event) {
	e.logger.Infof("Invoked CreatedUserEventMessage Consumer")

	var message *events.UpdatedUserEvent
	err := event.GetMessage(&message)
	if err != nil {
		e.logger.Infow("Error Getting Message", err)
		return
	}

	e.logger.Infow("UpdatedUserEventMessage", message.User)
}

func (e *DeletedUserEventMessage) OnMessage(event *message_bus.Event) {
	e.logger.Infof("Invoked CreatedUserEventMessage Consumer")

	var message *events.DeletedUserEvent
	err := event.GetMessage(&message)
	if err != nil {
		e.logger.Infow("Error Getting Message", err)
		return
	}

	e.logger.Infow("DeletedUserEventMessage", message.UserId)
}
