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

	api := server.Group("/api/ms-example/v1")

	example := api.Group("/example")
	{
		example.POST("", controller.CreateExample)
		example.GET(":exampleId", controller.GetExample)
		example.GET("", controller.GetAllExamples)
		example.DELETE(":exampleId", controller.DeleteExample)
		example.PATCH("", controller.UpdateExample)
	}
}
