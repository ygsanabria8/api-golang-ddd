package mediator_test

import (
	"api.ddd/pkgs/mediator"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestStruct struct{}

func TestCreateMessage(t *testing.T) {

	// Arrange
	structInstance := &TestStruct{}

	// Act
	message := mediator.CreateMessage(structInstance)

	// Assert
	assert.NotNil(t, message)
}

func TestGetMessageString(t *testing.T) {

	// Arrange
	structInstance := &TestStruct{}
	structInstanceString := "TestStruct"
	message := mediator.CreateMessage(structInstance)

	// Act
	messageString := message.GetMessageString()

	// Assert
	assert.NotNil(t, messageString)
	assert.Equal(t, structInstanceString, messageString)
}

func TestGetMessage(t *testing.T) {

	// Arrange
	structInstance := &TestStruct{}
	message := mediator.CreateMessage(structInstance)

	// Act
	messageStruct := message.GetMessage()

	// Assert
	assert.NotNil(t, message)
	assert.Equal(t, structInstance, messageStruct.(*TestStruct))
}
