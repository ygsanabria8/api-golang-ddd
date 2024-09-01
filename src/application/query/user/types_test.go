package query_test

import (
	mocks "api.ddd/mocks/src/domain/interfaces"
	"api.ddd/src/api/server"
	query "api.ddd/src/application/query/user"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWhenCallNewGetUserByIdQueryHandlerShouldReturnGetUserByIdQueryHandler(t *testing.T) {
	// Arrange
	logger := server.ProvideLogger()
	finder := &mocks.IUserFinder{}

	// Act
	queryHandler := query.NewGetUserByIdQueryHandler(logger, finder)

	// Assert
	assert.NotNil(t, queryHandler)
	assert.IsType(t, &query.GetUserByIdQueryHandler{}, queryHandler)
}

func TestWhenCallNewGetAllUsersQueryHandlerShouldReturnGetAllUsersQueryHandler(t *testing.T) {
	// Arrange
	logger := server.ProvideLogger()
	finder := &mocks.IUserFinder{}

	// Act
	queryHandler := query.NewGetAllUsersQueryHandler(logger, finder)

	// Assert
	assert.NotNil(t, queryHandler)
	assert.IsType(t, &query.GetAllUsersQueryHandler{}, queryHandler)
}
