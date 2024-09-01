package services_test

import (
	"api.ddd/src/api/server"
	services "api.ddd/src/domain/services/user"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGivenUserServiceDependenciesWhenCallNewUserServiceShouldReturnNewUserService(t *testing.T) {
	// Arrange
	loggerMock := server.ProvideLogger()

	serviceDependencies := services.UserService{
		Logger: loggerMock,
	}
	// Act
	service := services.NewUserService(serviceDependencies)

	// Assert
	assert.NotNil(t, service)
	assert.Equal(t, service, &serviceDependencies)
}
