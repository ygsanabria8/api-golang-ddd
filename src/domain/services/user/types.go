package services

import (
	"api.ddd/src/domain/interfaces"
	"go.uber.org/fx"
)

type UserService struct {
	fx.In
	noSqlRepository interfaces.IUserRepository `name:"NoSqlRepository"`
	noSqlFinder     interfaces.IUserFinder     `name:"NoSqlFinder"`
	sqlRepository   interfaces.IUserRepository `name:"SqlRepository"`
	sqlFinder       interfaces.IUserFinder     `name:"SqlFinder"`
}

func NewUserService(
	service UserService,
) interfaces.IUserService {
	return &service
}
