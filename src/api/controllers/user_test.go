package controllers_test

import (
	mocks "api.ddd/mocks/pkgs/mediator"
	"api.ddd/src/api/controllers"
	"api.ddd/src/api/server"
	command "api.ddd/src/application/command/user"
	"api.ddd/src/domain/aggregates"
	"bytes"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http/httptest"
	"testing"
)

var Id = uuid.New().String()
var request = &command.CreateUserCommand{
	Name:     "Clark",
	Lastname: "Tower",
	Email:    "clark@email.com",
	Age:      11,
}

var expectedResponse = &aggregates.User{
	Id:       Id,
	Name:     request.Name,
	Lastname: request.Lastname,
	Email:    request.Email,
	Age:      request.Age,
}

func TestWhenCreateUserShouldReturnUser(t *testing.T) {
	// Arrange
	methodMock := "Send"
	logger := server.ProvideLogger()
	mediator := &mocks.IDispatcher{}
	controller := controllers.NewController(logger, mediator)

	w := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
	ctx, _ := gin.CreateTestContext(w)
	requestBytes, _ := json.Marshal(request)
	ctx.Request = httptest.NewRequest(
		"POST", "/user",
		bytes.NewBuffer(requestBytes),
	)

	mediator.On(methodMock, mock.Anything).Return(expectedResponse, nil)

	// Act
	controller.CreateUser(ctx)
	var response *aggregates.User
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Fatal(err)
	}

	// Assert
	assert.NotNil(t, response)
	assert.IsType(t, &aggregates.User{}, response)
	assert.NotNil(t, Id, response.Id)
	mediator.AssertNumberOfCalls(t, methodMock, 1)
}

func TestWhenCreateUserShouldReturnErrorBody(t *testing.T) {
	// Arrange
	expectedErrorLen := 1
	expectedErrorMessage := "please check your body request"
	methodMock := "Send"
	logger := server.ProvideLogger()
	mediator := &mocks.IDispatcher{}
	controller := controllers.NewController(logger, mediator)

	w := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
	ctx, _ := gin.CreateTestContext(w)
	requestBytes, _ := json.Marshal("")
	ctx.Request = httptest.NewRequest(
		"POST", "/user",
		bytes.NewBuffer(requestBytes),
	)

	mediator.On(methodMock, mock.Anything).Return(nil, nil)

	// Act
	controller.CreateUser(ctx)

	// Assert
	assert.Equal(t, expectedErrorLen, len(ctx.Errors))
	assert.Equal(t, expectedErrorMessage, ctx.Errors[0].Error())
	mediator.AssertNumberOfCalls(t, methodMock, 0)
}

func TestWhenCreateUserShouldReturnErrorMediator(t *testing.T) {
	// Arrange
	methodMock := "Send"
	expectedErrorLen := 1
	expectedError := errors.New("error-mediator")
	logger := server.ProvideLogger()
	mediator := &mocks.IDispatcher{}
	controller := controllers.NewController(logger, mediator)

	w := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
	ctx, _ := gin.CreateTestContext(w)
	requestBytes, _ := json.Marshal(request)
	ctx.Request = httptest.NewRequest(
		"POST", "/user",
		bytes.NewBuffer(requestBytes),
	)

	mediator.On(methodMock, mock.Anything).Return(nil, expectedError)

	// Act
	controller.CreateUser(ctx)

	// Assert
	assert.Equal(t, expectedErrorLen, len(ctx.Errors))
	assert.Equal(t, expectedError.Error(), ctx.Errors[0].Error())
	mediator.AssertNumberOfCalls(t, methodMock, 1)
}

func TestWhenUpdateUserShouldReturnUser(t *testing.T) {
	// Arrange
	methodMock := "Send"
	logger := server.ProvideLogger()
	mediator := &mocks.IDispatcher{}
	controller := controllers.NewController(logger, mediator)

	w := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
	ctx, _ := gin.CreateTestContext(w)
	requestBytes, _ := json.Marshal(request)
	ctx.Request = httptest.NewRequest(
		"PATCH", "/user",
		bytes.NewBuffer(requestBytes),
	)

	mediator.On(methodMock, mock.Anything).Return(expectedResponse, nil)

	// Act
	controller.UpdateUser(ctx)
	var response *aggregates.User
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Fatal(err)
	}

	// Assert
	assert.NotNil(t, response)
	assert.IsType(t, &aggregates.User{}, response)
	assert.NotNil(t, Id, response.Id)
	mediator.AssertNumberOfCalls(t, methodMock, 1)
}

func TestWhenUpdateUserShouldReturnErrorBody(t *testing.T) {
	// Arrange
	expectedErrorLen := 1
	expectedErrorMessage := "please check your body request"
	methodMock := "Send"
	logger := server.ProvideLogger()
	mediator := &mocks.IDispatcher{}
	controller := controllers.NewController(logger, mediator)

	w := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
	ctx, _ := gin.CreateTestContext(w)
	requestBytes, _ := json.Marshal("")
	ctx.Request = httptest.NewRequest(
		"PATCH", "/user",
		bytes.NewBuffer(requestBytes),
	)

	mediator.On(methodMock, mock.Anything).Return(nil, nil)

	// Act
	controller.UpdateUser(ctx)

	// Assert
	assert.Equal(t, expectedErrorLen, len(ctx.Errors))
	assert.Equal(t, expectedErrorMessage, ctx.Errors[0].Error())
	mediator.AssertNumberOfCalls(t, methodMock, 0)
}

func TestWhenUpdateUserShouldReturnErrorMediator(t *testing.T) {
	// Arrange
	methodMock := "Send"
	expectedErrorLen := 1
	expectedError := errors.New("error-mediator")
	logger := server.ProvideLogger()
	mediator := &mocks.IDispatcher{}
	controller := controllers.NewController(logger, mediator)

	w := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
	ctx, _ := gin.CreateTestContext(w)
	requestBytes, _ := json.Marshal(request)
	ctx.Request = httptest.NewRequest(
		"PATCH", "/user",
		bytes.NewBuffer(requestBytes),
	)

	mediator.On(methodMock, mock.Anything).Return(nil, expectedError)

	// Act
	controller.UpdateUser(ctx)

	// Assert
	assert.Equal(t, expectedErrorLen, len(ctx.Errors))
	assert.Equal(t, expectedError.Error(), ctx.Errors[0].Error())
	mediator.AssertNumberOfCalls(t, methodMock, 1)
}

func TestWhenDeleteUserShouldReturnUser(t *testing.T) {
	// Arrange
	methodMock := "Send"
	logger := server.ProvideLogger()
	mediator := &mocks.IDispatcher{}
	controller := controllers.NewController(logger, mediator)
	expectedResponse := &command.UpdateUserCommand{
		Id: Id,
	}

	w := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
	ctx, _ := gin.CreateTestContext(w)
	ctx.Params = gin.Params{gin.Param{Key: "userId", Value: Id}}

	mediator.On(methodMock, mock.Anything).Return(expectedResponse, nil)

	// Act
	controller.DeleteUser(ctx)
	var response *command.UpdateUserCommand
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Fatal(err)
	}

	// Assert
	assert.NotNil(t, response)
	assert.IsType(t, &command.UpdateUserCommand{}, response)
	assert.NotNil(t, Id, response.Id)
	mediator.AssertNumberOfCalls(t, methodMock, 1)
}

func TestWhenDeleteUserShouldReturnErrorParams(t *testing.T) {
	// Arrange
	methodMock := "Send"
	expectedErrorLen := 1
	expectedErrorMessage := "please check your path request"
	logger := server.ProvideLogger()
	mediator := &mocks.IDispatcher{}
	controller := controllers.NewController(logger, mediator)

	w := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
	ctx, _ := gin.CreateTestContext(w)

	mediator.On(methodMock, mock.Anything).Return(nil, nil)

	// Act
	controller.DeleteUser(ctx)

	// Assert
	assert.Equal(t, expectedErrorLen, len(ctx.Errors))
	assert.Equal(t, expectedErrorMessage, ctx.Errors[0].Error())
	mediator.AssertNumberOfCalls(t, methodMock, 0)
}

func TestWhenDeleteUserShouldReturnErrorMediator(t *testing.T) {
	// Arrange
	methodMock := "Send"
	expectedErrorLen := 1
	expectedError := errors.New("error-error")
	logger := server.ProvideLogger()
	mediator := &mocks.IDispatcher{}
	controller := controllers.NewController(logger, mediator)

	w := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
	ctx, _ := gin.CreateTestContext(w)
	ctx.Params = gin.Params{gin.Param{Key: "userId", Value: Id}}

	mediator.On(methodMock, mock.Anything).Return(nil, expectedError)

	// Act
	controller.DeleteUser(ctx)

	// Assert
	assert.Equal(t, expectedErrorLen, len(ctx.Errors))
	assert.Equal(t, expectedError.Error(), ctx.Errors[0].Error())
	mediator.AssertNumberOfCalls(t, methodMock, 1)
}

func TestWhenGetUserByIdShouldReturnUser(t *testing.T) {
	// Arrange
	methodMock := "Send"
	logger := server.ProvideLogger()
	mediator := &mocks.IDispatcher{}
	controller := controllers.NewController(logger, mediator)

	w := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
	ctx, _ := gin.CreateTestContext(w)
	ctx.Params = gin.Params{gin.Param{Key: "userId", Value: Id}}

	mediator.On(methodMock, mock.Anything).Return(expectedResponse, nil)

	// Act
	controller.GetUserById(ctx)
	var response *command.UpdateUserCommand
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Fatal(err)
	}

	// Assert
	assert.NotNil(t, response)
	assert.IsType(t, &command.UpdateUserCommand{}, response)
	assert.NotNil(t, Id, response.Id)
	mediator.AssertNumberOfCalls(t, methodMock, 1)
}

func TestWhenGetUserByIdShouldReturnErrorParams(t *testing.T) {
	// Arrange
	methodMock := "Send"
	expectedErrorLen := 1
	expectedErrorMessage := "please check your path request"
	logger := server.ProvideLogger()
	mediator := &mocks.IDispatcher{}
	controller := controllers.NewController(logger, mediator)

	w := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
	ctx, _ := gin.CreateTestContext(w)

	mediator.On(methodMock, mock.Anything).Return(nil, nil)

	// Act
	controller.GetUserById(ctx)

	// Assert
	assert.Equal(t, expectedErrorLen, len(ctx.Errors))
	assert.Equal(t, expectedErrorMessage, ctx.Errors[0].Error())
	mediator.AssertNumberOfCalls(t, methodMock, 0)
}

func TestWhenGetUserByIdShouldReturnErrorMediator(t *testing.T) {
	// Arrange
	methodMock := "Send"
	expectedErrorLen := 1
	expectedError := errors.New("error-error")
	logger := server.ProvideLogger()
	mediator := &mocks.IDispatcher{}
	controller := controllers.NewController(logger, mediator)

	w := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
	ctx, _ := gin.CreateTestContext(w)
	ctx.Params = gin.Params{gin.Param{Key: "userId", Value: Id}}

	mediator.On(methodMock, mock.Anything).Return(nil, expectedError)

	// Act
	controller.GetUserById(ctx)

	// Assert
	assert.Equal(t, expectedErrorLen, len(ctx.Errors))
	assert.Equal(t, expectedError.Error(), ctx.Errors[0].Error())
	mediator.AssertNumberOfCalls(t, methodMock, 1)
}

func TestWhenGetAllUsersIdShouldReturnUser(t *testing.T) {
	// Arrange
	methodMock := "Send"
	logger := server.ProvideLogger()
	mediator := &mocks.IDispatcher{}
	controller := controllers.NewController(logger, mediator)

	w := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
	ctx, _ := gin.CreateTestContext(w)

	expectedResponse := &[]aggregates.User{
		{
			Id:       "1234",
			Name:     "John",
			Lastname: "Doe",
			Email:    "john.doe@example.com",
			Age:      30,
		},
	}

	mediator.On(methodMock, mock.Anything).Return(expectedResponse, nil)

	// Act
	controller.GetAllUsers(ctx)
	var response *[]aggregates.User
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Fatal(err)
	}

	// Assert
	assert.NotNil(t, response)
	assert.IsType(t, &[]aggregates.User{}, response)
	assert.Equal(t, len(*expectedResponse), len(*response))
	mediator.AssertNumberOfCalls(t, methodMock, 1)
}

func TestWhenGetAllUsersShouldReturnErrorMediator(t *testing.T) {
	// Arrange
	methodMock := "Send"
	expectedErrorLen := 1
	expectedError := errors.New("error-error")
	logger := server.ProvideLogger()
	mediator := &mocks.IDispatcher{}
	controller := controllers.NewController(logger, mediator)

	w := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
	ctx, _ := gin.CreateTestContext(w)

	mediator.On(methodMock, mock.Anything).Return(nil, expectedError)

	// Act
	controller.GetAllUsers(ctx)

	// Assert
	assert.Equal(t, expectedErrorLen, len(ctx.Errors))
	assert.Equal(t, expectedError.Error(), ctx.Errors[0].Error())
	mediator.AssertNumberOfCalls(t, methodMock, 1)
}
