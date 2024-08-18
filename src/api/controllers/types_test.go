package controllers_test

import (
	mocks "api.ddd/mocks/pkgs/mediator"
	"api.ddd/src/api/controllers"
	"api.ddd/src/api/server"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWhenCallNewControllerShouldReturnController(t *testing.T) {
	// Arrange
	logger := server.ProvideLogger()
	mediator := &mocks.IDispatcher{}

	// Act
	controller := controllers.NewController(logger, mediator)

	// Assert
	assert.NotNil(t, controller)
	assert.IsType(t, &controllers.Controller{}, controller)
}
