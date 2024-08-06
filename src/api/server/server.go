package server

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func StartGinServer(
	lifecycle fx.Lifecycle, server *gin.Engine,
) {
	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go func() {
				server.Use(gin.Recovery())
				err := server.Run("127.0.0.1:11000")
				if err != nil {
					panic("Error starting server")
				}
				fmt.Println("Server started")
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			fmt.Println("Shutting down server")
			return nil
		},
	})
}

func ProviderGinServer() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	return gin.New()
}
