package command_test

import (
	mocks "api.ddd/mocks/src/domain/interfaces"
	"api.ddd/pkgs/mediator"
	"api.ddd/src/api/server"
	command "api.ddd/src/application/command/user"
	"api.ddd/src/domain/aggregates"
	"errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestWhenCallCreateUserCommandHandlerShouldReturnUser(t *testing.T) {
	// Arrange
	logger := server.ProvideLogger()
	service := &mocks.IUserService{}
	methodMocker := "CreateUser"
	commandMessage := &command.CreateUserCommand{
		Name:     "",
		Lastname: "",
		Age:      1,
		Email:    "",
	}
	expectedUser := &aggregates.User{
		Id:       uuid.New().String(),
		Name:     "",
		Lastname: "",
		Age:      1,
		Email:    "",
	}

	message := mediator.CreateMessage(commandMessage)
	service.On(methodMocker, mock.Anything).Return(expectedUser, nil)
	commandHandler := command.NewCreateUserCommandHandler(logger, service)

	// Act
	user, err := commandHandler.Handler(message)

	// Assert
	assert.NotNil(t, user)
	assert.Nil(t, err)
	assert.IsType(t, &aggregates.User{}, user)
	assert.Equal(t, expectedUser.Id, user.(*aggregates.User).Id)
	service.AssertNumberOfCalls(t, methodMocker, 1)
}

func TestWhenCallCreateUserCommandHandlerShouldReturnError(t *testing.T) {
	// Arrange
	logger := server.ProvideLogger()
	service := &mocks.IUserService{}
	methodMocker := "CreateUser"
	expectedError := errors.New("some-error")
	commandMessage := &command.CreateUserCommand{
		Name:     "",
		Lastname: "",
		Age:      1,
		Email:    "",
	}

	message := mediator.CreateMessage(commandMessage)
	service.On(methodMocker, mock.Anything).Return(nil, expectedError)
	commandHandler := command.NewCreateUserCommandHandler(logger, service)

	// Act
	user, err := commandHandler.Handler(message)

	// Assert
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.Equal(t, expectedError, err)
	service.AssertNumberOfCalls(t, methodMocker, 1)
}

func TestWhenCallUpdateUserCommandHandlerShouldReturnUser(t *testing.T) {
	// Arrange
	logger := server.ProvideLogger()
	service := &mocks.IUserService{}
	methodMocker := "UpdateUser"
	userId := uuid.New().String()
	commandMessage := &command.UpdateUserCommand{
		Id:       userId,
		Name:     "",
		Lastname: "",
		Age:      1,
		Email:    "",
	}
	expectedUser := &aggregates.User{
		Id:       userId,
		Name:     "",
		Lastname: "",
		Age:      1,
		Email:    "",
	}

	message := mediator.CreateMessage(commandMessage)
	service.On(methodMocker, mock.Anything).Return(expectedUser, nil)
	commandHandler := command.NewUpdateUserCommandHandler(logger, service)

	// Act
	user, err := commandHandler.Handler(message)

	// Assert
	assert.NotNil(t, user)
	assert.Nil(t, err)
	assert.IsType(t, &aggregates.User{}, user)
	assert.Equal(t, userId, user.(*aggregates.User).Id)
	service.AssertNumberOfCalls(t, methodMocker, 1)
}

func TestWhenCallUpdateUserCommandHandlerShouldReturnError(t *testing.T) {
	// Arrange
	logger := server.ProvideLogger()
	service := &mocks.IUserService{}
	methodMocker := "UpdateUser"
	expectedError := errors.New("some-error")
	commandMessage := &command.UpdateUserCommand{
		Id:       uuid.New().String(),
		Name:     "",
		Lastname: "",
		Age:      1,
		Email:    "",
	}

	message := mediator.CreateMessage(commandMessage)
	service.On(methodMocker, mock.Anything).Return(nil, expectedError)
	commandHandler := command.NewUpdateUserCommandHandler(logger, service)

	// Act
	user, err := commandHandler.Handler(message)

	// Assert
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.Equal(t, expectedError, err)
	service.AssertNumberOfCalls(t, methodMocker, 1)
}

func TestWhenCallDeleteUserCommandHandlerShouldReturnUserId(t *testing.T) {
	// Arrange
	logger := server.ProvideLogger()
	service := &mocks.IUserService{}
	methodMocker := "DeleteUser"
	commandMessage := &command.DeleteUserCommand{
		Id: uuid.New().String(),
	}

	message := mediator.CreateMessage(commandMessage)
	service.On(methodMocker, mock.Anything).Return(nil)
	commandHandler := command.NewDeleteUserCommandHandler(logger, service)

	// Act
	user, err := commandHandler.Handler(message)

	// Assert
	assert.NotNil(t, user)
	assert.Nil(t, err)
	assert.IsType(t, &command.DeleteUserCommand{}, user)
	assert.Equal(t, commandMessage.Id, user.(*command.DeleteUserCommand).Id)
	service.AssertNumberOfCalls(t, methodMocker, 1)
}

func TestWhenCallDeleteUserCommandHandlerShouldReturnError(t *testing.T) {
	// Arrange
	logger := server.ProvideLogger()
	service := &mocks.IUserService{}
	methodMocker := "DeleteUser"
	expectedError := errors.New("some-error")

	commandMessage := &command.DeleteUserCommand{
		Id: uuid.New().String(),
	}

	message := mediator.CreateMessage(commandMessage)
	service.On(methodMocker, mock.Anything).Return(expectedError)
	commandHandler := command.NewDeleteUserCommandHandler(logger, service)

	// Act
	user, err := commandHandler.Handler(message)

	// Assert
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.Equal(t, expectedError, err)
	service.AssertNumberOfCalls(t, methodMocker, 1)
}
