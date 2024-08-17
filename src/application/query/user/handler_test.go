package query_test

import (
	mocks "api.ddd/mocks/src/domain/interfaces"
	"api.ddd/pkgs/mediator"
	"api.ddd/src/api/server"
	query "api.ddd/src/application/query/user"
	"api.ddd/src/domain/aggregates"
	"errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestWhenCallGetUserByIdQueryHandlerShouldReturnUser(t *testing.T) {
	// Arrange
	logger := server.ProvideLogger()
	service := &mocks.IUserService{}
	userId := uuid.New().String()
	methodMocker := "GetUserById"
	queryMessage := &query.GetUserByIdQuery{
		Id: userId,
	}
	expectedUser := &aggregates.User{
		Id: userId,
	}

	message := mediator.CreateMessage(queryMessage)
	service.On(methodMocker, mock.Anything).Return(expectedUser, nil)
	queryHandler := query.NewGetUserByIdQueryHandler(logger, service)

	// Act
	user, err := queryHandler.Handler(message)

	// Assert
	assert.NotNil(t, user)
	assert.Nil(t, err)
	assert.IsType(t, &aggregates.User{}, user)
	assert.Equal(t, userId, user.(*aggregates.User).Id)
	service.AssertNumberOfCalls(t, methodMocker, 1)
}

func TestWhenCallGetUserByIdQueryHandlerShouldReturnError(t *testing.T) {
	// Arrange
	logger := server.ProvideLogger()
	service := &mocks.IUserService{}
	methodMocker := "GetUserById"
	expectedError := errors.New("not-found")
	userId := uuid.New().String()
	queryMessage := &query.GetUserByIdQuery{
		Id: userId,
	}

	message := mediator.CreateMessage(queryMessage)
	service.On(methodMocker, mock.Anything).Return(nil, expectedError)
	queryHandler := query.NewGetUserByIdQueryHandler(logger, service)

	// Act
	user, err := queryHandler.Handler(message)

	// Assert
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.Equal(t, expectedError, err)
	service.AssertNumberOfCalls(t, methodMocker, 1)
}

func TestWhenCallGetAllUsersQueryHandlerShouldReturnUsers(t *testing.T) {
	// Arrange
	logger := server.ProvideLogger()
	service := &mocks.IUserService{}
	methodMocker := "GetAllUsers"
	queryMessage := &query.GetAllUsersQuery{}
	expectedUser := []*aggregates.User{
		{
			Id: uuid.New().String(),
		},
		{
			Id: uuid.New().String(),
		},
	}

	message := mediator.CreateMessage(queryMessage)
	service.On(methodMocker, mock.Anything).Return(expectedUser, nil)
	queryHandler := query.NewGetAllUsersQueryHandler(logger, service)

	// Act
	users, err := queryHandler.Handler(message)

	// Assert
	assert.NotNil(t, users)
	assert.Nil(t, err)
	assert.IsType(t, []*aggregates.User{}, users)
	assert.Equal(t, len(expectedUser), len(users.([]*aggregates.User)))
	service.AssertNumberOfCalls(t, methodMocker, 1)
}

func TestWhenCallGetAllUsersQueryHandlerShouldReturnError(t *testing.T) {
	// Arrange
	logger := server.ProvideLogger()
	service := &mocks.IUserService{}
	methodMocker := "GetAllUsers"
	expectedError := errors.New("not-found")
	queryMessage := &query.GetAllUsersQuery{}

	message := mediator.CreateMessage(queryMessage)
	service.On(methodMocker, mock.Anything).Return(nil, expectedError)
	queryHandler := query.NewGetAllUsersQueryHandler(logger, service)

	// Act
	user, err := queryHandler.Handler(message)

	// Assert
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.Equal(t, expectedError, err)
	service.AssertNumberOfCalls(t, methodMocker, 1)
}
