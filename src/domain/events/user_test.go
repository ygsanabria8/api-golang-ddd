package events_test

import (
	"api.ddd/src/domain/aggregates"
	"api.ddd/src/domain/events"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGivenUserWhenCallNewCreatedUserEventShouldReturnEvent(t *testing.T) {
	// Arrange
	user := &aggregates.User{
		Id: "some error",
	}

	// Act
	event := events.NewCreatedUserEvent(user)

	// Assert
	assert.NotNil(t, event)
	assert.NotNil(t, event.User)
}

func TestGivenUserWhenCallNewUpdatedUserEventShouldReturnEvent(t *testing.T) {
	// Arrange
	user := &aggregates.User{
		Id: "some error",
	}

	// Act
	event := events.NewUpdatedUserEvent(user)

	// Assert
	assert.NotNil(t, event)
	assert.NotNil(t, event.User)
}

func TestGivenUserIdWhenCallNewDeletedUserEventShouldReturnEvent(t *testing.T) {
	// Arrange
	userId := "some error"

	// Act
	event := events.NewDeletedUserEvent(userId)

	// Assert
	assert.NotNil(t, event)
	assert.NotNil(t, event.UserId)
}
