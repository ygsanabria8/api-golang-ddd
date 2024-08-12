package services

import (
	"api.ddd/src/domain/interfaces"
	"go.uber.org/fx"
)

type UserService struct {
	fx.In
	NoSqlRepository interfaces.IUserRepository `name:"NoSqlRepository"`
	NoSqlFinder     interfaces.IUserFinder     `name:"NoSqlFinder"`
	SqlRepository   interfaces.IUserRepository `name:"SqlRepository"`
	SqlFinder       interfaces.IUserFinder     `name:"SqlFinder"`
}

func NewUserService(
	service UserService,
) interfaces.IUserService {
	return &service
}
