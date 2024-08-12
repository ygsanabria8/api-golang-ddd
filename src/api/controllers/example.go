package controllers

import (
	"api.ddd/pkgs/mediator"
	command "api.ddd/src/application/command/user"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (c *Controller) CreateUser(ctx *gin.Context) {

	c.logger.Infof("Invoked CreateUser Controller")

	commandRequest := &command.CreateUserCommand{}
	if err := ctx.ShouldBindJSON(&commandRequest); err != nil {
		c.logger.Errorw("Error Body Deserialization", err)
		_ = ctx.Error(errors.New("please check your body request"))
		return
	}

	message := mediator.CreateMessage(commandRequest)
	response, err := c.mediator.Send(message)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *Controller) GetUser(ctx *gin.Context) {

	c.logger.Infof("Invoked GetUser Controller")
	ctx.JSON(http.StatusOK, gin.H{})
}

func (c *Controller) GetAllUsers(ctx *gin.Context) {

	c.logger.Infof("Invoked GetAllUsers Controller")
	ctx.JSON(http.StatusOK, gin.H{})
}

func (c *Controller) DeleteUser(ctx *gin.Context) {

	c.logger.Infof("Invoked DeleteUser Controller")
	ctx.JSON(http.StatusOK, gin.H{})
}

func (c *Controller) UpdateUser(ctx *gin.Context) {

	c.logger.Infof("Invoked UpdateUser Controller")

	commandRequest := &command.UpdateUserCommand{}
	if err := ctx.ShouldBindJSON(&commandRequest); err != nil {
		c.logger.Errorw("Error Body Deserialization", err)
		_ = ctx.Error(errors.New("please check your body request"))
		return
	}

	message := mediator.CreateMessage(commandRequest)
	response, err := c.mediator.Send(message)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, response)
}
