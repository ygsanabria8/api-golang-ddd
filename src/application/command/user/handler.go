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

	command := message.GetMessage().(*UpdateUserCommand)
	commandJson, _ := json.Marshal(command)
	h.logger.Info("Invoked UpdateUserCommandHandler: " + string(commandJson))

	newUser := &aggregates.User{
		Id:       command.Id,
		Name:     command.Name,
		Lastname: command.Lastname,
		Age:      command.Age,
		Email:    command.Email,
	}

	updatedUser, err := h.service.UpdateUser(newUser)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}

func (h *DeleteUserCommandHandler) Handler(message *mediator.Message) (interface{}, error) {

	command := message.GetMessage().(*DeleteUserCommand)

	commandJson, _ := json.Marshal(command)
	h.logger.Info("Invoked DeleteUserCommandHandler: " + string(commandJson))

	err := h.service.DeleteUser(command.Id)
	if err != nil {
		return nil, err
	}

	return command, nil
}
