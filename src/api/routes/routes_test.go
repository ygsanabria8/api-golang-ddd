package routes_test

import (
	mocks "api.ddd/mocks/src/api/controllers"
	"api.ddd/src/api/routes"
	"api.ddd/src/api/server"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWhenCallCreateUserShouldExecuteHandler(t *testing.T) {
	// Arrange
	w := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
	_, engine := gin.CreateTestContext(w)
	logger := server.ProvideLogger()
	mockedMethod := "CreateUser"

	controller := &mocks.IController{}
	controller.On(mockedMethod, mock.IsType(&gin.Context{})).Return()

	routes.ConfigureRoutes(engine, controller, logger)

	// Act
	req, _ := http.NewRequest("POST", "/api/ms-user/v1/user", nil)
	engine.ServeHTTP(w, req)

	// Assert
	controller.AssertCalled(t, mockedMethod, mock.IsType(&gin.Context{}))
	controller.AssertNumberOfCalls(t, mockedMethod, 1)
}

func TestWhenCallGetUserByIdShouldExecuteHandler(t *testing.T) {
	// Arrange
	w := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
	_, engine := gin.CreateTestContext(w)
	logger := server.ProvideLogger()
	mockedMethod := "GetUserById"

	controller := &mocks.IController{}
	controller.On(mockedMethod, mock.IsType(&gin.Context{})).Return()

	routes.ConfigureRoutes(engine, controller, logger)

	// Act
	req, _ := http.NewRequest("GET", "/api/ms-user/v1/user/234", nil)
	engine.ServeHTTP(w, req)

	// Assert
	controller.AssertCalled(t, mockedMethod, mock.IsType(&gin.Context{}))
	controller.AssertNumberOfCalls(t, mockedMethod, 1)
}

func TestWhenCallGetAllUsersShouldExecuteHandler(t *testing.T) {
	// Arrange
	w := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
	_, engine := gin.CreateTestContext(w)
	logger := server.ProvideLogger()
	mockedMethod := "GetAllUsers"

	controller := &mocks.IController{}
	controller.On(mockedMethod, mock.IsType(&gin.Context{})).Return()

	routes.ConfigureRoutes(engine, controller, logger)

	// Act
	req, _ := http.NewRequest("GET", "/api/ms-user/v1/user", nil)
	engine.ServeHTTP(w, req)

	// Assert
	controller.AssertCalled(t, mockedMethod, mock.IsType(&gin.Context{}))
	controller.AssertNumberOfCalls(t, mockedMethod, 1)
}

func TestWhenCallDeleteUserShouldExecuteHandler(t *testing.T) {
	// Arrange
	w := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
	_, engine := gin.CreateTestContext(w)
	logger := server.ProvideLogger()
	mockedMethod := "DeleteUser"

	controller := &mocks.IController{}
	controller.On(mockedMethod, mock.IsType(&gin.Context{})).Return()

	routes.ConfigureRoutes(engine, controller, logger)

	// Act
	req, _ := http.NewRequest("DELETE", "/api/ms-user/v1/user/123", nil)
	engine.ServeHTTP(w, req)

	// Assert
	controller.AssertCalled(t, mockedMethod, mock.IsType(&gin.Context{}))
	controller.AssertNumberOfCalls(t, mockedMethod, 1)
}

func TestWhenCallUpdateUserShouldExecuteHandler(t *testing.T) {
	// Arrange
	w := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
	_, engine := gin.CreateTestContext(w)
	logger := server.ProvideLogger()
	mockedMethod := "UpdateUser"

	controller := &mocks.IController{}
	controller.On(mockedMethod, mock.IsType(&gin.Context{})).Return()

	routes.ConfigureRoutes(engine, controller, logger)

	// Act
	req, _ := http.NewRequest("PATCH", "/api/ms-user/v1/user", nil)
	engine.ServeHTTP(w, req)

	// Assert
	controller.AssertCalled(t, mockedMethod, mock.IsType(&gin.Context{}))
	controller.AssertNumberOfCalls(t, mockedMethod, 1)
}
