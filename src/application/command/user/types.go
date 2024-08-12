package command

import (
	"api.ddd/src/api/server"
	"api.ddd/src/domain/interfaces"
)

type CreateUserCommand struct {
	Name     string
	Lastname string
	Age      int32
	Email    string
}

type CreateUserCommandHandler struct {
	logger  *server.Logger
	service interfaces.IUserService
}

func NewCreateUserCommandHandler(
	logger *server.Logger,
	service interfaces.IUserService,
) *CreateUserCommandHandler {
	return &CreateUserCommandHandler{
		logger:  logger,
		service: service,
	}
}

type UpdateUserCommand struct {
	Id       string
	Name     string
	Lastname string
	Age      int32
	Email    string
}

type UpdateUserCommandHandler struct {
	logger  *server.Logger
	service interfaces.IUserService
}

func NewUpdateUserCommandHandler(
	logger *server.Logger,
	service interfaces.IUserService,
) *UpdateUserCommandHandler {
	return &UpdateUserCommandHandler{
		logger:  logger,
		service: service,
	}
}
