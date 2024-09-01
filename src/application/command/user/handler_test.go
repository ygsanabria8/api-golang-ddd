package command_test

import (
	kafkamocks "api.ddd/mocks/pkgs/message_bus"
	domainmocks "api.ddd/mocks/src/domain/interfaces"
	"api.ddd/pkgs/mediator"
	"api.ddd/src/api/server"
	command "api.ddd/src/application/command/user"
	"api.ddd/src/domain/aggregates"
	"api.ddd/src/domain/events"
	"errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

var KafkaMock = &kafkamocks.IMessageBusClient{}
var NoSqlRepositoryMock = &domainmocks.IUserRepository{}
var SqlRepositoryMock = &domainmocks.IUserRepository{}
var loggerMock = server.ProvideLogger()

func CleanMocks() {
	NoSqlRepositoryMock = &domainmocks.IUserRepository{}
	SqlRepositoryMock = &domainmocks.IUserRepository{}
	KafkaMock = &kafkamocks.IMessageBusClient{}
}

func TestWhenCallCreateUserCommandHandlerShouldReturnUser(t *testing.T) {
	// Arrange
	CleanMocks()
	mockedMethodStorage := "CreateUser"
	mockedMethodKafka := "SendMessage"
	commandMessage := &command.CreateUserCommand{
		Name:     "Clark",
		Lastname: "Tower",
		Email:    "clark@email.com",
		Age:      11,
	}
	expectedUser := &aggregates.User{
		Id:       "some-user-id",
		Name:     "Clark",
		Lastname: "Tower",
		Email:    "clark@email.com",
		Age:      11,
	}

	SqlRepositoryMock.On(mockedMethodStorage, mock.IsType(&aggregates.User{})).
		Return(expectedUser, nil)

	doneGoroutine := make(chan struct{})
	NoSqlRepositoryMock.On(mockedMethodStorage, mock.IsType(&aggregates.User{})).
		Run(func(args mock.Arguments) {
			defer close(doneGoroutine)
		}).
		Return(expectedUser, nil)

	KafkaMock.On(mockedMethodKafka, mock.IsType(&events.CreatedUserEvent{}))

	message := mediator.CreateMessage(commandMessage)
	commandHandler := command.NewCreateUserCommandHandler(command.CreateUserCommandHandler{
		Logger:          loggerMock,
		SqlRepository:   SqlRepositoryMock,
		NoSqlRepository: NoSqlRepositoryMock,
		Kafka:           KafkaMock,
	})

	// Act
	user, err := commandHandler.Handler(message)
	<-doneGoroutine

	// Assert
	assert.NotNil(t, user)
	assert.Nil(t, err)
	assert.IsType(t, &aggregates.User{}, user)
	assert.Equal(t, expectedUser.Id, user.(*aggregates.User).Id)
	SqlRepositoryMock.AssertNumberOfCalls(t, mockedMethodStorage, 1)
	NoSqlRepositoryMock.AssertNumberOfCalls(t, mockedMethodStorage, 1)
	KafkaMock.AssertNumberOfCalls(t, mockedMethodKafka, 1)
}

func TestWhenCallCreateUserCommandHandlerShouldReturnError(t *testing.T) {
	// Arrange
	CleanMocks()
	mockedMethodStorage := "CreateUser"
	mockedMethodKafka := "SendMessage"
	expectedError := errors.New("some error")
	commandMessage := &command.CreateUserCommand{
		Name:     "Clark",
		Lastname: "Tower",
		Email:    "clark@email.com",
		Age:      11,
	}

	SqlRepositoryMock.On(mockedMethodStorage, mock.IsType(&aggregates.User{})).
		Return(nil, expectedError)

	NoSqlRepositoryMock.On(mockedMethodStorage, mock.IsType(&aggregates.User{})).
		Return(nil, nil)

	KafkaMock.On(mockedMethodKafka, mock.IsType(&events.CreatedUserEvent{}))

	message := mediator.CreateMessage(commandMessage)
	commandHandler := command.NewCreateUserCommandHandler(command.CreateUserCommandHandler{
		Logger:          loggerMock,
		SqlRepository:   SqlRepositoryMock,
		NoSqlRepository: NoSqlRepositoryMock,
		Kafka:           KafkaMock,
	})

	// Act
	user, err := commandHandler.Handler(message)

	// Assert
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.Equal(t, expectedError, err)
	SqlRepositoryMock.AssertNumberOfCalls(t, mockedMethodStorage, 1)
	NoSqlRepositoryMock.AssertNumberOfCalls(t, mockedMethodStorage, 0)
	KafkaMock.AssertNumberOfCalls(t, mockedMethodKafka, 0)
}

func TestWhenCallUpdateUserCommandHandlerShouldReturnUser(t *testing.T) {
	// Arrange
	CleanMocks()
	mockedMethodStorage := "UpdateUser"
	mockedMethodKafka := "SendMessage"
	commandMessage := &command.UpdateUserCommand{
		Id:       "some-user-id",
		Name:     "Clark",
		Lastname: "Tower",
		Email:    "clark@email.com",
		Age:      11,
	}
	expectedUser := &aggregates.User{
		Id:       "some-user-id",
		Name:     "Clark",
		Lastname: "Tower",
		Email:    "clark@email.com",
		Age:      11,
	}

	SqlRepositoryMock.On(mockedMethodStorage, mock.IsType(&aggregates.User{})).
		Return(expectedUser, nil)

	doneGoroutine := make(chan struct{})
	NoSqlRepositoryMock.On(mockedMethodStorage, mock.IsType(&aggregates.User{})).
		Run(func(args mock.Arguments) {
			defer close(doneGoroutine)
		}).
		Return(expectedUser, nil)

	KafkaMock.On(mockedMethodKafka, mock.IsType(&events.CreatedUserEvent{}))

	message := mediator.CreateMessage(commandMessage)
	commandHandler := command.NewUpdateUserCommandHandler(command.UpdateUserCommandHandler{
		Logger:          loggerMock,
		SqlRepository:   SqlRepositoryMock,
		NoSqlRepository: NoSqlRepositoryMock,
		Kafka:           KafkaMock,
	})

	// Act
	user, err := commandHandler.Handler(message)
	<-doneGoroutine

	// Assert
	assert.NotNil(t, user)
	assert.Nil(t, err)
	assert.IsType(t, &aggregates.User{}, user)
	assert.Equal(t, commandMessage.Id, user.(*aggregates.User).Id)
	SqlRepositoryMock.AssertNumberOfCalls(t, mockedMethodStorage, 1)
	NoSqlRepositoryMock.AssertNumberOfCalls(t, mockedMethodStorage, 1)
	KafkaMock.AssertNumberOfCalls(t, mockedMethodKafka, 1)
}

func TestWhenCallUpdateUserCommandHandlerShouldReturnError(t *testing.T) {
	// Arrange
	CleanMocks()
	mockedMethodStorage := "UpdateUser"
	mockedMethodKafka := "SendMessage"
	expectedError := errors.New("some error")
	commandMessage := &command.UpdateUserCommand{
		Id:       "some-id",
		Name:     "Clark",
		Lastname: "Tower",
		Email:    "clark@email.com",
		Age:      11,
	}

	SqlRepositoryMock.On(mockedMethodStorage, mock.IsType(&aggregates.User{})).
		Return(nil, expectedError)

	NoSqlRepositoryMock.On(mockedMethodStorage, mock.IsType(&aggregates.User{})).
		Return(nil, nil)

	KafkaMock.On(mockedMethodKafka, mock.IsType(&events.UpdatedUserEvent{}))

	message := mediator.CreateMessage(commandMessage)
	commandHandler := command.NewUpdateUserCommandHandler(command.UpdateUserCommandHandler{
		Logger:          loggerMock,
		SqlRepository:   SqlRepositoryMock,
		NoSqlRepository: NoSqlRepositoryMock,
		Kafka:           KafkaMock,
	})

	// Act
	user, err := commandHandler.Handler(message)

	// Assert
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.Equal(t, expectedError, err)
	SqlRepositoryMock.AssertNumberOfCalls(t, mockedMethodStorage, 1)
	NoSqlRepositoryMock.AssertNumberOfCalls(t, mockedMethodStorage, 0)
	KafkaMock.AssertNumberOfCalls(t, mockedMethodKafka, 0)
}

func TestWhenCallDeleteUserCommandHandlerShouldReturnUserId(t *testing.T) {
	// Arrange
	CleanMocks()
	mockedMethodStorage := "DeleteUser"
	mockedMethodKafka := "SendMessage"
	userId := uuid.New().String()
	commandMessage := &command.DeleteUserCommand{
		Id: userId,
	}

	SqlRepositoryMock.On(mockedMethodStorage, mock.IsType(userId)).
		Return(nil)

	doneGoroutine := make(chan struct{})
	NoSqlRepositoryMock.On(mockedMethodStorage, mock.IsType(userId)).
		Run(func(args mock.Arguments) {
			defer close(doneGoroutine)
		}).
		Return(nil)

	KafkaMock.On(mockedMethodKafka, mock.IsType(&events.DeletedUserEvent{}))

	message := mediator.CreateMessage(commandMessage)
	commandHandler := command.NewDeleteUserCommandHandler(command.DeleteUserCommandHandler{
		Logger:          loggerMock,
		SqlRepository:   SqlRepositoryMock,
		NoSqlRepository: NoSqlRepositoryMock,
		Kafka:           KafkaMock,
	})

	// Act
	user, err := commandHandler.Handler(message)
	<-doneGoroutine

	// Assert
	assert.NotNil(t, user)
	assert.Nil(t, err)
	assert.IsType(t, &command.DeleteUserCommand{}, user)
	assert.Equal(t, commandMessage.Id, user.(*command.DeleteUserCommand).Id)
	SqlRepositoryMock.AssertNumberOfCalls(t, mockedMethodStorage, 1)
	NoSqlRepositoryMock.AssertNumberOfCalls(t, mockedMethodStorage, 1)
	KafkaMock.AssertNumberOfCalls(t, mockedMethodKafka, 1)
}

func TestWhenCallDeleteUserCommandHandlerShouldReturnError(t *testing.T) {
	// Arrange
	CleanMocks()
	mockedMethodStorage := "DeleteUser"
	mockedMethodKafka := "SendMessage"
	userId := uuid.New().String()
	expectedError := errors.New("some error")
	commandMessage := &command.DeleteUserCommand{
		Id: userId,
	}

	SqlRepositoryMock.On(mockedMethodStorage, mock.IsType(userId)).
		Return(expectedError)

	NoSqlRepositoryMock.On(mockedMethodStorage, mock.IsType(userId)).
		Return(expectedError)

	KafkaMock.On(mockedMethodKafka, mock.IsType(&events.DeletedUserEvent{}))

	message := mediator.CreateMessage(commandMessage)
	commandHandler := command.NewDeleteUserCommandHandler(command.DeleteUserCommandHandler{
		Logger:          loggerMock,
		SqlRepository:   SqlRepositoryMock,
		NoSqlRepository: NoSqlRepositoryMock,
		Kafka:           KafkaMock,
	})

	// Act
	user, err := commandHandler.Handler(message)

	// Assert
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.Equal(t, expectedError, err)
	SqlRepositoryMock.AssertNumberOfCalls(t, mockedMethodStorage, 1)
	NoSqlRepositoryMock.AssertNumberOfCalls(t, mockedMethodStorage, 0)
	KafkaMock.AssertNumberOfCalls(t, mockedMethodKafka, 0)
}
