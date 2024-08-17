package services

import (
	"api.ddd/pkgs/message_bus"
	"api.ddd/src/api/server"
	"api.ddd/src/domain/interfaces"
	"go.uber.org/fx"
)

type UserService struct {
	fx.In
	Logger          *server.Logger
	NoSqlFinder     interfaces.IUserFinder     `name:"NoSqlFinder"`
	NoSqlRepository interfaces.IUserRepository `name:"NoSqlRepository"`
	SqlRepository   interfaces.IUserRepository `name:"SqlRepository"`
	SqlFinder       interfaces.IUserFinder     `name:"SqlFinder"`
	Kafka           message_bus.IMessageBusClient
}

func NewUserService(
	service *UserService,
) interfaces.IUserService {
	return service
}
