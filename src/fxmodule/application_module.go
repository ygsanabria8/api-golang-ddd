package fxmodule

import (
	"api.ddd/pkgs/mediator"
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
