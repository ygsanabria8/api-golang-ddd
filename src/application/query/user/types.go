package query

import (
	"api.ddd/src/api/server"
	"api.ddd/src/domain/interfaces"
	"go.uber.org/fx"
)

type Params struct {
	fx.In
	Logger      *server.Logger
	NoSqlFinder interfaces.IUserFinder `name:"NoSqlFinder"`
}

type GetUserByIdQuery struct {
	Id string
}

type GetUserByIdQueryHandler struct {
	logger *server.Logger
	finder interfaces.IUserFinder
}

func NewGetUserByIdQueryHandler(
	params Params,
) *GetUserByIdQueryHandler {
	return &GetUserByIdQueryHandler{
		logger: params.Logger,
		finder: params.NoSqlFinder,
	}
}

type GetAllUsersQuery struct {
	Id string
}

type GetAllUsersQueryHandler struct {
	logger *server.Logger
	finder interfaces.IUserFinder
}

func NewGetAllUsersQueryHandler(
	params Params,
) *GetAllUsersQueryHandler {
	return &GetAllUsersQueryHandler{
		logger: params.Logger,
		finder: params.NoSqlFinder,
	}
}
