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
) error {
	err := dispatcher.RegisterHandler(createUserCommandHandler, &command.CreateUserCommand{})
	err = dispatcher.RegisterHandler(updateUserCommandHandler, &command.UpdateUserCommand{})
	err = dispatcher.RegisterHandler(deleteUserCommandHandler, &command.DeleteUserCommand{})
	err = dispatcher.RegisterHandler(getUserByIdQueryHandler, &query.GetUserByIdQuery{})
	err = dispatcher.RegisterHandler(getAllUsersQueryHandler, &query.GetAllUsersQuery{})

	return err
}
