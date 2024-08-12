package fxmodule

import (
	command "api.ddd/src/application/command/user"
	query "api.ddd/src/application/query/user"
	"go.uber.org/fx"
)

var ApplicationModule = fx.Module(
	"Application Module",
	fx.Provide(
		command.NewCreateUserCommandHandler,
		command.NewUpdateUserCommandHandler,
		command.NewDeleteUserCommandHandler,
	),
	fx.Provide(
		query.NewGetUserByIdQueryHandler,
		query.NewGetAllUsersQueryHandler,
	),
)
