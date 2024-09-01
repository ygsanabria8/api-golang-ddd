package consumers_test

import (
	mocks "api.ddd/mocks/pkgs/message_bus/utils"
	"api.ddd/src/api/consumers"
	"api.ddd/src/api/server"
	"api.ddd/src/domain/aggregates"
	"api.ddd/src/domain/events"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGivenValidMessageWhenCallCreatedUserEventHandlerShouldNotReturnError(t *testing.T) {
	// Arrange
	mockedMethod := "GetMessage"
	event := &mocks.IEvent{}
	eventData := &events.CreatedUserEvent{
		User: &aggregates.User{
			Id:       "some-user-id",
			Name:     "Clark",
			Lastname: "Tower",
			Email:    "clark@email.com",
			Age:      11,
		},
	}

	eventDataBytes, _ := json.Marshal(eventData)
	event.On(mockedMethod).Return(eventDataBytes)
	handlerEvent := consumers.NewCreatedUserEventMessage(server.ProvideLogger())

	// Act
	err := handlerEvent.OnMessage(event)

	// Assert
	assert.Nil(t, err)
	event.AssertNumberOfCalls(t, mockedMethod, 1)
}

func TestGivenValidMessageWhenCallCreatedUserEventHandlerShouldReturnError(t *testing.T) {
	// Arrange
	mockedMethod := "GetMessage"
	event := &mocks.IEvent{}

	event.On(mockedMethod).Return(nil)
	handlerEvent := consumers.NewCreatedUserEventMessage(server.ProvideLogger())

	// Act
	err := handlerEvent.OnMessage(event)

	// Assert
	assert.NotNil(t, err)
	event.AssertNumberOfCalls(t, mockedMethod, 1)
}

func TestGivenValidMessageWhenCallUpdatedUserEventHandlerShouldNotReturnError(t *testing.T) {
	// Arrange
	mockedMethod := "GetMessage"
	event := &mocks.IEvent{}
	eventData := &events.UpdatedUserEvent{
		User: &aggregates.User{
			Id:       "some-user-id",
			Name:     "Clark",
			Lastname: "Tower",
			Email:    "clark@email.com",
			Age:      11,
		},
	}

	eventDataBytes, _ := json.Marshal(eventData)
	event.On(mockedMethod).Return(eventDataBytes)
	handlerEvent := consumers.NewUpdatedUserEventMessage(server.ProvideLogger())

	// Act
	err := handlerEvent.OnMessage(event)

	// Assert
	assert.Nil(t, err)
	event.AssertNumberOfCalls(t, mockedMethod, 1)
}

func TestGivenValidMessageWhenCallUpdatedUserEventHandlerShouldReturnError(t *testing.T) {
	// Arrange
	mockedMethod := "GetMessage"
	event := &mocks.IEvent{}

	event.On(mockedMethod).Return(nil)
	handlerEvent := consumers.NewUpdatedUserEventMessage(server.ProvideLogger())

	// Act
	err := handlerEvent.OnMessage(event)

	// Assert
	assert.NotNil(t, err)
	event.AssertNumberOfCalls(t, mockedMethod, 1)
}

func TestGivenValidMessageWhenCallDeletedUserEventHandlerShouldNotReturnError(t *testing.T) {
	// Arrange
	mockedMethod := "GetMessage"
	event := &mocks.IEvent{}
	eventData := &events.DeletedUserEvent{
		UserId: "some-user-id",
	}

	eventDataBytes, _ := json.Marshal(eventData)
	event.On(mockedMethod).Return(eventDataBytes)
	handlerEvent := consumers.NewDeletedUserEventMessage(server.ProvideLogger())

	// Act
	err := handlerEvent.OnMessage(event)

	// Assert
	assert.Nil(t, err)
	event.AssertNumberOfCalls(t, mockedMethod, 1)
}

func TestGivenValidMessageWhenCallDeletedUserEventHandlerShouldReturnError(t *testing.T) {
	// Arrange
	mockedMethod := "GetMessage"
	event := &mocks.IEvent{}

	event.On(mockedMethod).Return(nil)
	handlerEvent := consumers.NewDeletedUserEventMessage(server.ProvideLogger())

	// Act
	err := handlerEvent.OnMessage(event)

	// Assert
	assert.NotNil(t, err)
	event.AssertNumberOfCalls(t, mockedMethod, 1)
}
