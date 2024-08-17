package server_test

import (
	"api.ddd/src/api/server"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func TestWhenCallConfigureCorsShouldConfigureCors(t *testing.T) {

	// Arrange
	w := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
	_, engine := gin.CreateTestContext(w)
	config := &server.Configuration{
		Server: &server.Server{
			Origin: []string{"https://example.com"},
			Port:   "1000",
			Header: []string{"Authorization", "Content-Type"},
		},
	}

	// Act
	server.ConfigureCors(engine, config)

	// Assert
	assert.Equal(t, len(engine.RouterGroup.Handlers), 1)
}
