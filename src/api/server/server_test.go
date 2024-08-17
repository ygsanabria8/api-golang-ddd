package server_test

import (
	"api.ddd/src/api/server"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWhenCallProvideGinServerShouldReturnEngine(t *testing.T) {
	// Arrange
	logger := server.ProvideLogger()

	// Act
	engine := server.ProvideGinServer(logger)

	// Assert
	assert.NotNil(t, engine)
}
