package services_test

import (
	kafkamocks "api.ddd/mocks/pkgs/message_bus"
	domainmocks "api.ddd/mocks/src/domain/interfaces"
	"api.ddd/src/api/server"
	"api.ddd/src/domain/aggregates"
	"api.ddd/src/domain/events"
	services "api.ddd/src/domain/services/user"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

var KafkaMock = &kafkamocks.IMessageBusClient{}
var NoSqlFinderMock = &domainmocks.IUserFinder{}
var NoSqlRepositoryMock = &domainmocks.IUserRepository{}
var SqlFinderMock = &domainmocks.IUserFinder{}
var SqlRepositoryMock = &domainmocks.IUserRepository{}
var loggerMock = server.ProvideLogger()

func CleanMocks() {
	NoSqlFinderMock = &domainmocks.IUserFinder{}
	NoSqlRepositoryMock = &domainmocks.IUserRepository{}
	SqlFinderMock = &domainmocks.IUserFinder{}
	SqlRepositoryMock = &domainmocks.IUserRepository{}
	KafkaMock = &kafkamocks.IMessageBusClient{}
}

func TestGivenUserIdWhenCallGetUserByIdShouldReturnUser(t *testing.T) {
	// Arrange
	CleanMocks()
	userId := "some-user-id"
	expectedUser := &aggregates.User{
		Id: userId,
	}
	mockedMethod := "GetUserById"
	NoSqlFinderMock.On(mockedMethod, mock.AnythingOfType("string")).
		Return(expectedUser, nil)

	service := &services.UserService{
		NoSqlRepository: NoSqlRepositoryMock,
		SqlFinder:       SqlFinderMock,
		SqlRepository:   SqlRepositoryMock,
		NoSqlFinder:     NoSqlFinderMock,
		Kafka:           KafkaMock,
		Logger:          loggerMock,
	}

	// Act
	user, err := service.GetUserById(userId)

	// Assert
	assert.NotNil(t, user)
	assert.Nil(t, err)
	assert.Equal(t, userId, user.Id)
	NoSqlFinderMock.AssertNumberOfCalls(t, mockedMethod, 1)
}

func TestGivenUserIdWhenCallGetUserByIdShouldReturnError(t *testing.T) {
	// Arrange
	CleanMocks()
	userId := "some-user-id"
	expectedError := errors.New("some error")
	mockedMethod := "GetUserById"
	NoSqlFinderMock.On(mockedMethod, mock.AnythingOfType("string")).
		Return(nil, expectedError)

	service := &services.UserService{
		NoSqlRepository: NoSqlRepositoryMock,
		SqlFinder:       SqlFinderMock,
		SqlRepository:   SqlRepositoryMock,
		NoSqlFinder:     NoSqlFinderMock,
		Kafka:           KafkaMock,
		Logger:          loggerMock,
	}

	// Act
	user, err := service.GetUserById(userId)

	// Assert
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.Equal(t, expectedError, err)
	NoSqlFinderMock.AssertNumberOfCalls(t, mockedMethod, 1)
}

func TestWhenCallGetAllUsersShouldReturnUsers(t *testing.T) {
	// Arrange
	CleanMocks()
	expectedUsers := []*aggregates.User{
		{
			Id: "1",
		},
		{
			Id: "2",
		},
	}
	mockedMethod := "GetAllUsers"
	NoSqlFinderMock.On(mockedMethod, mock.Anything).
		Return(expectedUsers, nil)

	service := &services.UserService{
		NoSqlRepository: NoSqlRepositoryMock,
		SqlFinder:       SqlFinderMock,
		SqlRepository:   SqlRepositoryMock,
		NoSqlFinder:     NoSqlFinderMock,
		Kafka:           KafkaMock,
		Logger:          loggerMock,
	}

	// Act
	users, err := service.GetAllUsers()

	// Assert
	assert.NotNil(t, users)
	assert.Nil(t, err)
	assert.Equal(t, len(expectedUsers), len(users))
	NoSqlFinderMock.AssertNumberOfCalls(t, mockedMethod, 1)
}

func TestWhenCallGetAllUsersShouldReturnError(t *testing.T) {
	// Arrange
	CleanMocks()
	expectedError := errors.New("some error")
	mockedMethod := "GetAllUsers"
	NoSqlFinderMock.On(mockedMethod, mock.Anything).
		Return(nil, expectedError)

	service := &services.UserService{
		NoSqlRepository: NoSqlRepositoryMock,
		SqlFinder:       SqlFinderMock,
		SqlRepository:   SqlRepositoryMock,
		NoSqlFinder:     NoSqlFinderMock,
		Kafka:           KafkaMock,
		Logger:          loggerMock,
	}

	// Act
	users, err := service.GetAllUsers()

	// Assert
	assert.Nil(t, users)
	assert.NotNil(t, err)
	assert.Equal(t, expectedError, err)
	NoSqlFinderMock.AssertNumberOfCalls(t, mockedMethod, 1)
}

func TestGivenUserWhenCallCreateUserShouldReturnUser(t *testing.T) {
	// Arrange
	CleanMocks()
	mockedMethodStorage := "CreateUser"
	mockedMethodKafka := "SendMessage"
	expectedUser := &aggregates.User{
		Id:       "some-user-id",
		Name:     "Clark",
		Lastname: "Tower",
		Email:    "clark@email.com",
		Age:      11,
	}

	doneGoroutine := make(chan struct{})
	SqlRepositoryMock.On(mockedMethodStorage, mock.IsType(&aggregates.User{})).
		Return(expectedUser, nil)

	NoSqlRepositoryMock.On(mockedMethodStorage, mock.IsType(&aggregates.User{})).
		Run(func(args mock.Arguments) {
			defer close(doneGoroutine)
		}).
		Return(expectedUser, nil)

	KafkaMock.On(mockedMethodKafka, mock.IsType(&events.CreatedUserEvent{}))

	service := &services.UserService{
		NoSqlRepository: NoSqlRepositoryMock,
		SqlFinder:       SqlFinderMock,
		SqlRepository:   SqlRepositoryMock,
		NoSqlFinder:     NoSqlFinderMock,
		Kafka:           KafkaMock,
		Logger:          loggerMock,
	}

	// Act
	user, err := service.CreateUser(expectedUser)
	<-doneGoroutine

	// Assert
	assert.NotNil(t, user)
	assert.Nil(t, err)
	assert.Equal(t, expectedUser.Id, user.Id)
	SqlRepositoryMock.AssertNumberOfCalls(t, mockedMethodStorage, 1)
	NoSqlRepositoryMock.AssertNumberOfCalls(t, mockedMethodStorage, 1)
	KafkaMock.AssertNumberOfCalls(t, mockedMethodKafka, 1)
}

func TestGivenUserWhenCallCreateUserShouldReturnError(t *testing.T) {
	// Arrange
	CleanMocks()
	mockedMethodStorage := "CreateUser"
	mockedMethodKafka := "SendMessage"
	expectedError := errors.New("some error")

	SqlRepositoryMock.On(mockedMethodStorage, mock.IsType(&aggregates.User{})).
		Return(nil, expectedError)

	NoSqlRepositoryMock.On(mockedMethodStorage, mock.IsType(&aggregates.User{})).
		Return(nil, nil)

	KafkaMock.On(mockedMethodKafka, mock.IsType(&events.CreatedUserEvent{}))

	service := &services.UserService{
		NoSqlRepository: NoSqlRepositoryMock,
		SqlFinder:       SqlFinderMock,
		SqlRepository:   SqlRepositoryMock,
		NoSqlFinder:     NoSqlFinderMock,
		Kafka:           KafkaMock,
		Logger:          loggerMock,
	}

	// Act
	user, err := service.CreateUser(&aggregates.User{})

	// Assert
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.Equal(t, expectedError, err)
	SqlRepositoryMock.AssertNumberOfCalls(t, mockedMethodStorage, 1)
	NoSqlRepositoryMock.AssertNumberOfCalls(t, mockedMethodStorage, 0)
	KafkaMock.AssertNumberOfCalls(t, mockedMethodKafka, 0)
}

func TestGivenUserWhenCallUpdateUserShouldReturnUser(t *testing.T) {
	// Arrange
	CleanMocks()
	mockedMethodStorage := "UpdateUser"
	mockedMethodKafka := "SendMessage"
	expectedUser := &aggregates.User{
		Id:       "some-user-id",
		Name:     "Clark",
		Lastname: "Tower",
		Email:    "clark@email.com",
		Age:      11,
	}

	doneGoroutine := make(chan struct{})
	SqlRepositoryMock.On(mockedMethodStorage, mock.IsType(&aggregates.User{})).
		Return(expectedUser, nil)

	NoSqlRepositoryMock.On(mockedMethodStorage, mock.IsType(&aggregates.User{})).
		Run(func(args mock.Arguments) {
			defer close(doneGoroutine)
		}).
		Return(expectedUser, nil)

	KafkaMock.On(mockedMethodKafka, mock.IsType(&events.UpdatedUserEvent{}))

	service := &services.UserService{
		NoSqlRepository: NoSqlRepositoryMock,
		SqlFinder:       SqlFinderMock,
		SqlRepository:   SqlRepositoryMock,
		NoSqlFinder:     NoSqlFinderMock,
		Kafka:           KafkaMock,
		Logger:          loggerMock,
	}

	// Act
	user, err := service.UpdateUser(expectedUser)
	<-doneGoroutine

	// Assert
	assert.NotNil(t, user)
	assert.Nil(t, err)
	assert.Equal(t, expectedUser.Id, user.Id)
	SqlRepositoryMock.AssertNumberOfCalls(t, mockedMethodStorage, 1)
	NoSqlRepositoryMock.AssertNumberOfCalls(t, mockedMethodStorage, 1)
	KafkaMock.AssertNumberOfCalls(t, mockedMethodKafka, 1)
}

func TestGivenUserWhenCallUpdateUserShouldReturnError(t *testing.T) {
	// Arrange
	CleanMocks()
	mockedMethodStorage := "UpdateUser"
	mockedMethodKafka := "SendMessage"
	expectedError := errors.New("some error")

	SqlRepositoryMock.On(mockedMethodStorage, mock.IsType(&aggregates.User{})).
		Return(nil, expectedError)

	NoSqlRepositoryMock.On(mockedMethodStorage, mock.IsType(&aggregates.User{})).
		Return(nil, nil)

	KafkaMock.On(mockedMethodKafka, mock.IsType(&events.UpdatedUserEvent{}))

	service := &services.UserService{
		NoSqlRepository: NoSqlRepositoryMock,
		SqlFinder:       SqlFinderMock,
		SqlRepository:   SqlRepositoryMock,
		NoSqlFinder:     NoSqlFinderMock,
		Kafka:           KafkaMock,
		Logger:          loggerMock,
	}

	// Act
	user, err := service.UpdateUser(&aggregates.User{})

	// Assert
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.Equal(t, expectedError, err)
	SqlRepositoryMock.AssertNumberOfCalls(t, mockedMethodStorage, 1)
	NoSqlRepositoryMock.AssertNumberOfCalls(t, mockedMethodStorage, 0)
	KafkaMock.AssertNumberOfCalls(t, mockedMethodKafka, 0)
}

func TestGivenUserWhenCallDeleteUserShouldReturnNilError(t *testing.T) {
	// Arrange
	CleanMocks()
	mockedMethodStorage := "DeleteUser"
	mockedMethodKafka := "SendMessage"
	userId := "some-user-id"

	doneGoroutine := make(chan struct{})
	SqlRepositoryMock.On(mockedMethodStorage, mock.IsType(userId)).
		Return(nil)

	NoSqlRepositoryMock.On(mockedMethodStorage, mock.IsType(userId)).
		Run(func(args mock.Arguments) {
			defer close(doneGoroutine)
		}).
		Return(nil)

	KafkaMock.On(mockedMethodKafka, mock.IsType(&events.DeletedUserEvent{}))

	service := &services.UserService{
		NoSqlRepository: NoSqlRepositoryMock,
		SqlFinder:       SqlFinderMock,
		SqlRepository:   SqlRepositoryMock,
		NoSqlFinder:     NoSqlFinderMock,
		Kafka:           KafkaMock,
		Logger:          loggerMock,
	}

	// Act
	err := service.DeleteUser(userId)
	<-doneGoroutine

	// Assert
	assert.Nil(t, err)
	SqlRepositoryMock.AssertNumberOfCalls(t, mockedMethodStorage, 1)
	NoSqlRepositoryMock.AssertNumberOfCalls(t, mockedMethodStorage, 1)
	KafkaMock.AssertNumberOfCalls(t, mockedMethodKafka, 1)
}

func TestGivenUserWhenCallDeleteUserShouldReturnError(t *testing.T) {
	// Arrange
	CleanMocks()
	mockedMethodStorage := "DeleteUser"
	mockedMethodKafka := "SendMessage"
	userId := "some-user-id"
	expectedError := errors.New("some error")

	SqlRepositoryMock.On(mockedMethodStorage, mock.IsType(userId)).
		Return(expectedError)

	NoSqlRepositoryMock.On(mockedMethodStorage, mock.IsType(userId)).
		Return(expectedError)

	KafkaMock.On(mockedMethodKafka, mock.IsType(&events.DeletedUserEvent{}))

	service := &services.UserService{
		NoSqlRepository: NoSqlRepositoryMock,
		SqlFinder:       SqlFinderMock,
		SqlRepository:   SqlRepositoryMock,
		NoSqlFinder:     NoSqlFinderMock,
		Kafka:           KafkaMock,
		Logger:          loggerMock,
	}

	// Act
	err := service.DeleteUser(userId)

	// Assert
	assert.NotNil(t, err)
	assert.Equal(t, expectedError, err)
	SqlRepositoryMock.AssertNumberOfCalls(t, mockedMethodStorage, 1)
	NoSqlRepositoryMock.AssertNumberOfCalls(t, mockedMethodStorage, 0)
	KafkaMock.AssertNumberOfCalls(t, mockedMethodKafka, 0)
}
