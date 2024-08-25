package consumers

import (
	"api.ddd/pkgs/message_bus/utils"
	"api.ddd/src/domain/events"
	constants "api.ddd/src/utils"
	"encoding/json"
)

func (e *CreatedUserEventMessage) OnMessage(event utils.IEvent) error {
	e.logger.Infof("Invoked CreatedUserEventMessage Consumer")

	var message *events.CreatedUserEvent
	err := json.Unmarshal(event.GetMessage(), &message)
	if err != nil {
		e.logger.Infow(constants.ErrorDeserializerMessageBus, err)
		return err
	}

	e.logger.Infoln(message.User)
	return nil
}

func (e *UpdatedUserEventMessage) OnMessage(event utils.IEvent) error {
	e.logger.Infof("Invoked UpdatedUserEventMessage Consumer")

	var message *events.UpdatedUserEvent
	err := json.Unmarshal(event.GetMessage(), &message)
	if err != nil {
		e.logger.Infow(constants.ErrorDeserializerMessageBus, err)
		return err
	}

	e.logger.Infoln(message.User)
	return nil
}

func (e *DeletedUserEventMessage) OnMessage(event utils.IEvent) error {
	e.logger.Infof("Invoked DeletedUserEventMessage Consumer")

	var message *events.DeletedUserEvent
	err := json.Unmarshal(event.GetMessage(), &message)
	if err != nil {
		e.logger.Infow(constants.ErrorDeserializerMessageBus, err)
		return err
	}

	e.logger.Infoln(message.UserId)
	return nil
}
