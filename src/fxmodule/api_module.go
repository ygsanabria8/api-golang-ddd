package fxmodule

import (
	"api.ddd/pkgs/mediator"
	"api.ddd/src/api/controllers"
	"api.ddd/src/api/routes"
	"api.ddd/src/api/server"
	command "api.ddd/src/application/command/user"
	query "api.ddd/src/application/query/user"
	"go.uber.org/fx"
)

var ApiModule = fx.Module(
	"Api Module",
	fx.Provide(server.ProvideLogger),
	fx.Provide(server.ProvideGinServer),
	fx.Provide(mediator.NewDispatcherMemory),
	fx.Provide(controllers.NewController),
	fx.Invoke(server.ConfigureCors),
	fx.Invoke(routes.ConfigureRoutes),
	fx.Invoke(RegisterHandlers),
)

func RegisterHandlers(
	dispatcher mediator.IDispatcher,
	createUserCommandHandler *command.CreateUserCommandHandler,
	updateUserCommandHandler *command.UpdateUserCommandHandler,
	deleteUserCommandHandler *command.DeleteUserCommandHandler,
	getUserByIdQueryHandler *query.GetUserByIdQueryHandler,
	getAllUsersQueryHandler *query.GetAllUsersQueryHandler,
) {
	_ = dispatcher.RegisterHandler(createUserCommandHandler, &command.CreateUserCommand{})
	_ = dispatcher.RegisterHandler(updateUserCommandHandler, &command.UpdateUserCommand{})
	_ = dispatcher.RegisterHandler(deleteUserCommandHandler, &command.DeleteUserCommand{})
	_ = dispatcher.RegisterHandler(getUserByIdQueryHandler, &query.GetUserByIdQuery{})
	_ = dispatcher.RegisterHandler(getAllUsersQueryHandler, &query.GetAllUsersQuery{})
}
