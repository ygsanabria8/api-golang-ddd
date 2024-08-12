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

func (h *UpdateUserCommandHandler) Handler(message *mediator.Message) (interface{}, error) {

	commandRequest := message.GetMessage().(*UpdateUserCommand)
	commandJson, _ := json.Marshal(commandRequest)
	h.logger.Info("Invoked UpdateUserCommandHandler: " + string(commandJson))

	newUser := &aggregates.User{
		Id:       commandRequest.Id,
		Name:     commandRequest.Name,
		Lastname: commandRequest.Lastname,
		Age:      commandRequest.Age,
		Email:    commandRequest.Email,
	}

	updatedUser, err := h.service.UpdateUser(newUser)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}
