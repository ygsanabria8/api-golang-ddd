package command

import (
	"api.ddd/pkgs/message_bus"
	"api.ddd/src/api/server"
	"api.ddd/src/domain/interfaces"
	"go.uber.org/fx"
)

type Params struct {
	fx.In
	Logger          *server.Logger
	NoSqlRepository interfaces.IUserRepository `name:"NoSqlRepository"`
	SqlRepository   interfaces.IUserRepository `name:"SqlRepository"`
	Kafka           message_bus.IMessageBusClient
}

type CreateUserCommand struct {
	Name     string
	Lastname string
	Age      int32
	Email    string
}

type CreateUserCommandHandler struct {
	logger          *server.Logger
	noSqlRepository interfaces.IUserRepository `name:"NoSqlRepository"`
	sqlRepository   interfaces.IUserRepository `name:"SqlRepository"`
	kafka           message_bus.IMessageBusClient
}

func NewCreateUserCommandHandler(
	params Params,
) *CreateUserCommandHandler {
	return &CreateUserCommandHandler{
		logger:          params.Logger,
		sqlRepository:   params.SqlRepository,
		noSqlRepository: params.NoSqlRepository,
		kafka:           params.Kafka,
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
	logger          *server.Logger
	noSqlRepository interfaces.IUserRepository `name:"NoSqlRepository"`
	sqlRepository   interfaces.IUserRepository `name:"SqlRepository"`
	kafka           message_bus.IMessageBusClient
}

func NewUpdateUserCommandHandler(
	params Params,
) *UpdateUserCommandHandler {
	return &UpdateUserCommandHandler{
		logger:          params.Logger,
		sqlRepository:   params.SqlRepository,
		noSqlRepository: params.NoSqlRepository,
		kafka:           params.Kafka,
	}
}

type DeleteUserCommand struct {
	Id string
}

type DeleteUserCommandHandler struct {
	logger          *server.Logger
	noSqlRepository interfaces.IUserRepository `name:"NoSqlRepository"`
	sqlRepository   interfaces.IUserRepository `name:"SqlRepository"`
	kafka           message_bus.IMessageBusClient
}

func NewDeleteUserCommandHandler(
	params Params,
) *DeleteUserCommandHandler {
	return &DeleteUserCommandHandler{
		logger:          params.Logger,
		sqlRepository:   params.SqlRepository,
		noSqlRepository: params.NoSqlRepository,
		kafka:           params.Kafka,
	}
}
