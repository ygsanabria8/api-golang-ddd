package fxmodule

import (
	"api.ddd/pkgs/mediator"
	"api.ddd/src/api/controllers"
	"api.ddd/src/api/routes"
	"api.ddd/src/api/server"
	command "api.ddd/src/application/command/user"
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
) {
	_ = dispatcher.RegisterHandler(createUserCommandHandler, &command.CreateUserCommand{})
	_ = dispatcher.RegisterHandler(updateUserCommandHandler, &command.UpdateUserCommand{})
}
