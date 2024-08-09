package mediator_test

import (
	"api.ddd/pkgs/mediator"
	"api.ddd/src/api/server"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest"
)

func TestNewDispatcherMemory(t *testing.T) {

	// Arrange
	logger := &server.Logger{
		SugaredLogger: zaptest.NewLogger(t, zaptest.Level(zap.WarnLevel)).Sugar(),
	}

	// Act
	dispatcher := mediator.NewDispatcherMemory(logger)

	// Assert
	assert.NotNil(t, dispatcher)
}
