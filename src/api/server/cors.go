package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
)

// ConfigureCors Configure Server CORS
func ConfigureCors(server *gin.Engine, config *Configuration) {
	server.Use(
		cors.New(
			cors.Config{
				AllowOrigins:     config.Server.Origin,
				AllowMethods:     []string{http.MethodGet, http.MethodPatch, http.MethodPost, http.MethodDelete},
				AllowHeaders:     config.Server.Header,
				AllowCredentials: false,
			},
		),
	)
}
