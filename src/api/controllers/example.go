package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (c *Controller) CreateUser(ctx *gin.Context) {

	c.logger.Infof("Invoked CreateUser Controller")
	ctx.JSON(http.StatusOK, gin.H{})
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
	ctx.JSON(http.StatusOK, gin.H{})
}
