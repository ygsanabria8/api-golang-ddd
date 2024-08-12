package fxmodule

import (
	command "api.ddd/src/application/command/user"
	"go.uber.org/fx"
)

var ApplicationModule = fx.Module(
	"Application Module",
	fx.Provide(
		command.NewCreateUserCommandHandler,
		command.NewUpdateUserCommandHandler,
	),
)
