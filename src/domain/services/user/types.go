package services

import (
	"api.ddd/src/api/server"
	"api.ddd/src/domain/interfaces"
	"go.uber.org/fx"
)

type UserService struct {
	fx.In
	Logger *server.Logger
}

func NewUserService(
	service UserService,
) interfaces.IUserService {
	return &service
}
