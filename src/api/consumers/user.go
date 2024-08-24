package consumers

import (
	"api.ddd/pkgs/message_bus/utils"
	"api.ddd/src/domain/events"
	"encoding/json"
)

func (e *CreatedUserEventMessage) OnMessage(event utils.IEvent) {
	e.logger.Infof("Invoked CreatedUserEventMessage Consumer")

	var message *events.CreatedUserEvent
	err := json.Unmarshal(event.GetMessage(), &message)
	if err != nil {
		e.logger.Infow("Error Getting Message", err)
		return
	}

	e.logger.Infow("CreatedUserEventMessage", message.User)
}

func (e *UpdatedUserEventMessage) OnMessage(event utils.IEvent) {
	e.logger.Infof("Invoked CreatedUserEventMessage Consumer")

	var message *events.UpdatedUserEvent
	err := json.Unmarshal(event.GetMessage(), &message)
	if err != nil {
		e.logger.Infow("Error Getting Message", err)
		return
	}

	e.logger.Infow("UpdatedUserEventMessage", message.User)
}

func (e *DeletedUserEventMessage) OnMessage(event utils.IEvent) {
	e.logger.Infof("Invoked CreatedUserEventMessage Consumer")

	var message *events.DeletedUserEvent
	err := json.Unmarshal(event.GetMessage(), &message)
	if err != nil {
		e.logger.Infow("Error Getting Message", err)
		return
	}

	e.logger.Infow("DeletedUserEventMessage", message.UserId)
}
