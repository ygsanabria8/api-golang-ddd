package server

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func StartGinServer(
	lifecycle fx.Lifecycle, server *gin.Engine,
	logger *Logger, config *Configuration,
) {
	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go func() {
				server.Use(gin.Recovery())
				err := server.Run("127.0.0.1:" + config.Server.Port)
				if err != nil {
					logger.Fatalf("Error starting server: %s", err.Error())
					//panic("Error starting server")
				}
				logger.Infof("Server started on port " + config.Server.Port)
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Infof("Shutting down server")
			return nil
		},
	})
}

// ProvideGinServer Provide Gin Engine Server
func ProvideGinServer(logger *Logger) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = logger.GetGinLogger()
	return gin.New()
}
