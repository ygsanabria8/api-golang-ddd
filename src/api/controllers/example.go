package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (c *Controller) CreateExample(ctx *gin.Context) {

	c.logger.Infof("Invoked CreateExample Controller")
	ctx.JSON(http.StatusOK, gin.H{})
}

func (c *Controller) GetExample(ctx *gin.Context) {

	c.logger.Infof("Invoked GetExample Controller")
	ctx.JSON(http.StatusOK, gin.H{})
}

func (c *Controller) GetAllExamples(ctx *gin.Context) {

	c.logger.Infof("Invoked GetAllExamples Controller")
	ctx.JSON(http.StatusOK, gin.H{})
}

func (c *Controller) DeleteExample(ctx *gin.Context) {

	c.logger.Infof("Invoked DeleteExample Controller")
	ctx.JSON(http.StatusOK, gin.H{})
}

func (c *Controller) UpdateExample(ctx *gin.Context) {

	c.logger.Infof("Invoked UpdateExample Controller")
	ctx.JSON(http.StatusOK, gin.H{})
}
