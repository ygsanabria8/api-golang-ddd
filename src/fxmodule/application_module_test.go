package fxmodule_test

import (
	kafkamocks "api.ddd/mocks/pkgs/message_bus"
	domainmocks "api.ddd/mocks/src/domain/interfaces"
	"api.ddd/pkgs/mediator"
	"api.ddd/src/api/server"
	command "api.ddd/src/application/command/user"
	query "api.ddd/src/application/query/user"
	"api.ddd/src/fxmodule"
	"github.com/stretchr/testify/assert"
	"testing"
)

var loggerMock = server.ProvideLogger()

var commandParams = command.Params{
	Logger:          loggerMock,
	SqlRepository:   &domainmocks.IUserRepository{},
	NoSqlRepository: &domainmocks.IUserRepository{},
	Kafka:           &kafkamocks.IMessageBusClient{},
}

var queryParams = query.Params{
	Logger:      loggerMock,
	NoSqlFinder: &domainmocks.IUserFinder{},
}

func TestGivenCommandHandlersWhenCallRegisterHandlersShouldNotReturnError(t *testing.T) {
	// Arrange
	dispatcher := mediator.NewDispatcherMemory(server.ProvideLogger())
	createUserCommandHandler := command.NewCreateUserCommandHandler(commandParams)
	updateUserCommandHandler := command.NewUpdateUserCommandHandler(commandParams)
	deleteUserCommandHandler := command.NewDeleteUserCommandHandler(commandParams)
	getUserByIdQueryHandler := query.NewGetUserByIdQueryHandler(queryParams)
	getAllUsersQueryHandler := query.NewGetAllUsersQueryHandler(queryParams)

	// Act
	err := fxmodule.RegisterHandlers(
		dispatcher, createUserCommandHandler, updateUserCommandHandler, deleteUserCommandHandler,
		getUserByIdQueryHandler, getAllUsersQueryHandler,
	)

	// Assert
	assert.Nil(t, err)
}
