package controllers

import (
	"api.ddd/pkgs/mediator"
	"api.ddd/src/api/server"
	"github.com/gin-gonic/gin"
)

type IUserController interface {
	CreateUser(ctx *gin.Context)
	GetUserById(ctx *gin.Context)
	GetAllUsers(ctx *gin.Context)
	DeleteUser(ctx *gin.Context)
	UpdateUser(ctx *gin.Context)
}

type IController interface {
	IUserController
}

type Controller struct {
	logger   *server.Logger
	mediator mediator.IDispatcher
}

func NewController(
	logger *server.Logger,
	mediator mediator.IDispatcher,
) IController {
	return &Controller{
		logger:   logger,
		mediator: mediator,
	}
}
