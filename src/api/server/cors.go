package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
)

// ConfigureCors Configure Server CORS
func ConfigureCors(server *gin.Engine) {
	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{http.MethodGet, http.MethodPatch, http.MethodPost, http.MethodDelete},
		AllowHeaders:     []string{"*"},
		AllowCredentials: false,
	}))
}
