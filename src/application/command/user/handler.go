package command

import (
	"api.ddd/pkgs/mediator"
	"api.ddd/src/domain/aggregates"
	"encoding/json"
)

func (h *CreateUserCommandHandler) Handler(message *mediator.Message) (interface{}, error) {

	command := message.GetMessage().(*CreateUserCommand)

	commandJson, _ := json.Marshal(command)
	h.logger.Info("Invoked CreateUserCommandHandler: " + string(commandJson))

	newUser := &aggregates.User{
		Name:     command.Name,
		Lastname: command.Lastname,
		Age:      command.Age,
		Email:    command.Email,
	}

	savedUser, err := h.service.CreateUser(newUser)
	if err != nil {
		return nil, err
	}

	return savedUser, nil
}
