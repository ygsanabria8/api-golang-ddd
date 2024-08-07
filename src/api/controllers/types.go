package controllers

import (
	"api.ddd/src/api/server"
)

type Controller struct {
	logger *server.Logger
}

func NewController(
	logger *server.Logger,
) *Controller {
	return &Controller{
		logger: logger,
	}
}
