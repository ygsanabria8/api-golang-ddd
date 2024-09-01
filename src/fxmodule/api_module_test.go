package fxmodule_test

import (
	mocks "api.ddd/mocks/pkgs/message_bus"
	"api.ddd/pkgs/message_bus"
	"api.ddd/src/api/consumers"
	"api.ddd/src/api/server"
	"api.ddd/src/fxmodule"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestGivenConsumersWhenCallProvideMessageBusClientShouldReturnIMessageBusClient(t *testing.T) {
	// Arrange
	logger := server.ProvideLogger()
	messageBusMock := &mocks.IMessageBus{}
	config := &server.Configuration{
		AppName: "my-dummy-app",
		Kafka: &server.Kafka{
			Topics: &server.Topics{
				CreatedUser: "CreatedUser",
				UpdatedUser: "UpdatedUser",
				DeletedUser: "DeletedUser",
			},
		},
	}
	createdUserEventMessage := consumers.NewCreatedUserEventMessage(logger)
	updatedUserEventMessage := consumers.NewUpdatedUserEventMessage(logger)
	deletedUserEventMessage := consumers.NewDeletedUserEventMessage(logger)

	messageBusMock.On("WithAppName", mock.Anything, mock.Anything).Return(messageBusMock)
	messageBusMock.On("WithConsumer", mock.Anything, mock.Anything).Return(messageBusMock)
	messageBusMock.On("WitProducer", mock.Anything, mock.Anything).Return(messageBusMock)
	messageBusMock.On("Build").Return(&message_bus.MessageBusClient{})

	// Act
	messageBus := fxmodule.ProvideMessageBusClient(
		messageBusMock, config, createdUserEventMessage,
		updatedUserEventMessage, deletedUserEventMessage,
	)

	// Assert
	assert.NotNil(t, messageBus)
	assert.IsType(t, &message_bus.MessageBusClient{}, messageBus)
}
