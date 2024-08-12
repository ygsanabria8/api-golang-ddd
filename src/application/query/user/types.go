package finder

import (
	"api.ddd/src/api/server"
	"api.ddd/src/domain/interfaces"
)

type GetUserByIdQuery struct {
	Id string
}

type GetUserByIdQueryHandler struct {
	logger  *server.Logger
	service interfaces.IUserService
}

func NewGetUserByIdQueryHandler(
	logger *server.Logger,
	service interfaces.IUserService,
) *GetUserByIdQueryHandler {
	return &GetUserByIdQueryHandler{
		logger:  logger,
		service: service,
	}
}
