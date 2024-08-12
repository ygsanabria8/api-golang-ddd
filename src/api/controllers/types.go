package controllers

import (
	"api.ddd/pkgs/mediator"
	"api.ddd/src/api/server"
)

type Controller struct {
	logger   *server.Logger
	mediator mediator.IDispatcher
}

func NewController(
	logger *server.Logger,
	mediator mediator.IDispatcher,
) *Controller {
	return &Controller{
		logger:   logger,
		mediator: mediator,
	}
}
