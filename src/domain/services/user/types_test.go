package services_test

import (
	kafkamocks "api.ddd/mocks/pkgs/message_bus"
	domainmocks "api.ddd/mocks/src/domain/interfaces"
	"api.ddd/src/api/server"
	services "api.ddd/src/domain/services/user"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGivenUserServiceDependenciesWhenCallNewUserServiceShouldReturnNewUserService(t *testing.T) {
	// Arrange
	KafkaMock := &kafkamocks.IMessageBusClient{}
	NoSqlFinderMock := &domainmocks.IUserFinder{}
	NoSqlRepositoryMock := &domainmocks.IUserRepository{}
	SqlFinderMock := &domainmocks.IUserFinder{}
	SqlRepositoryMock := &domainmocks.IUserRepository{}
	loggerMock := server.ProvideLogger()

	serviceDependencies := services.UserService{
		NoSqlRepository: NoSqlRepositoryMock,
		SqlFinder:       SqlFinderMock,
		SqlRepository:   SqlRepositoryMock,
		NoSqlFinder:     NoSqlFinderMock,
		Kafka:           KafkaMock,
		Logger:          loggerMock,
	}
	// Act
	service := services.NewUserService(serviceDependencies)

	// Assert
	assert.NotNil(t, service)
	assert.Equal(t, service, &serviceDependencies)
}
