package command_test

import (
	mocks "api.ddd/mocks/src/domain/interfaces"
	"api.ddd/src/api/server"
	command "api.ddd/src/application/command/user"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWhenCallNewCreateUserCommandHandlerShouldReturnCreateUserCommandHandler(t *testing.T) {
	// Arrange
	logger := server.ProvideLogger()
	service := &mocks.IUserService{}

	// Act
	commandHandler := command.NewCreateUserCommandHandler(logger, service)

	// Assert
	assert.NotNil(t, commandHandler)
	assert.IsType(t, &command.CreateUserCommandHandler{}, commandHandler)
}

func TestWhenCallNewUpdateUserCommandHandlerShouldReturnUpdateUserCommandHandler(t *testing.T) {
	// Arrange
	logger := server.ProvideLogger()
	service := &mocks.IUserService{}

	// Act
	commandHandler := command.NewUpdateUserCommandHandler(logger, service)

	// Assert
	assert.NotNil(t, commandHandler)
	assert.IsType(t, &command.UpdateUserCommandHandler{}, commandHandler)
}

func TestWhenCallNewNewDeleteUserCommandHandlerShouldReturnDeleteUserCommandHandler(t *testing.T) {
	// Arrange
	logger := server.ProvideLogger()
	service := &mocks.IUserService{}

	// Act
	commandHandler := command.NewDeleteUserCommandHandler(logger, service)

	// Assert
	assert.NotNil(t, commandHandler)
	assert.IsType(t, &command.DeleteUserCommandHandler{}, commandHandler)
}
