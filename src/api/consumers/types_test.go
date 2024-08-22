package consumers_test

import (
	"api.ddd/src/api/consumers"
	"api.ddd/src/api/server"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWhenCallNewCreatedUserEventMessageShouldReturnInstance(t *testing.T) {
	// Arrange
	logger := server.ProvideLogger()

	// Act
	instance := consumers.NewCreatedUserEventMessage(logger)

	// Assert
	assert.NotNil(t, instance)
}

func TestWhenCallNewUpdatedUserEventMessageShouldReturnInstance(t *testing.T) {
	// Arrange
	logger := server.ProvideLogger()

	// Act
	instance := consumers.NewUpdatedUserEventMessage(logger)

	// Assert
	assert.NotNil(t, instance)
}

func TestWhenCallDeletedUserEventMessageShouldReturnInstance(t *testing.T) {
	// Arrange
	logger := server.ProvideLogger()

	// Act
	instance := consumers.NewDeletedUserEventMessage(logger)

	// Assert
	assert.NotNil(t, instance)
}
