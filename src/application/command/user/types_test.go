package command_test

import (
	command "api.ddd/src/application/command/user"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWhenCallNewCreateUserCommandHandlerShouldReturnCreateUserCommandHandler(t *testing.T) {
	// Act
	commandHandler := command.NewCreateUserCommandHandler(command.Params{
		Logger:          loggerMock,
		SqlRepository:   SqlRepositoryMock,
		NoSqlRepository: NoSqlRepositoryMock,
		Kafka:           KafkaMock,
	})

	// Assert
	assert.NotNil(t, commandHandler)
	assert.IsType(t, &command.CreateUserCommandHandler{}, commandHandler)
}

func TestWhenCallNewUpdateUserCommandHandlerShouldReturnUpdateUserCommandHandler(t *testing.T) {
	// Act
	commandHandler := command.NewUpdateUserCommandHandler(command.Params{
		Logger:          loggerMock,
		SqlRepository:   SqlRepositoryMock,
		NoSqlRepository: NoSqlRepositoryMock,
		Kafka:           KafkaMock,
	})

	// Assert
	assert.NotNil(t, commandHandler)
	assert.IsType(t, &command.UpdateUserCommandHandler{}, commandHandler)
}

func TestWhenCallNewNewDeleteUserCommandHandlerShouldReturnDeleteUserCommandHandler(t *testing.T) {
	// Act
	commandHandler := command.NewDeleteUserCommandHandler(command.Params{
		Logger:          loggerMock,
		SqlRepository:   SqlRepositoryMock,
		NoSqlRepository: NoSqlRepositoryMock,
		Kafka:           KafkaMock,
	})

	// Assert
	assert.NotNil(t, commandHandler)
	assert.IsType(t, &command.DeleteUserCommandHandler{}, commandHandler)
}
