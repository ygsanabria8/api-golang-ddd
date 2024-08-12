package routes

import (
	"api.ddd/src/api/controllers"
	"api.ddd/src/api/server"
	"github.com/gin-gonic/gin"
)

func ConfigureRoutes(
	server *gin.Engine, controller *controllers.Controller,
	logger *server.Logger,
) {
	logger.Info("Setting up routes")

	api := server.Group("/api/ms-user/v1")

	user := api.Group("/user")
	{
		user.POST("", controller.CreateUser)
		user.GET(":userId", controller.GetUserById)
		user.GET("", controller.GetAllUsers)
		user.DELETE(":userId", controller.DeleteUser)
		user.PATCH("", controller.UpdateUser)
	}
}
