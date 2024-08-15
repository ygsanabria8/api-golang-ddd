package fxmodule

import (
	"api.ddd/pkgs/mediator"
	"api.ddd/pkgs/message_bus"
	"api.ddd/src/api/consumers"
	"api.ddd/src/api/controllers"
	"api.ddd/src/api/routes"
	"api.ddd/src/api/server"
	"api.ddd/src/domain/events"
	"go.uber.org/fx"
)

var ApiModule = fx.Module(
	"Api Module",
	fx.Provide(server.ProvideLogger),
	fx.Provide(server.ProvideGinServer),
	fx.Provide(mediator.NewDispatcherMemory),
	fx.Provide(controllers.NewController),
	fx.Provide(ProvideMessageBusClient),
	fx.Provide(
		consumers.NewCreatedUserEventMessage,
		consumers.NewUpdatedUserEventMessage,
		consumers.NewDeletedUserEventMessage,
	),
	fx.Invoke(server.ConfigureCors),
	fx.Invoke(routes.ConfigureRoutes),
)

func ProvideMessageBusClient(
	conn message_bus.IMessageBus,
	createdUserEventMessage *consumers.CreatedUserEventMessage,
	updatedUserEventMessage *consumers.UpdatedUserEventMessage,
	deletedUserEventMessage *consumers.DeletedUserEventMessage,
) message_bus.IMessageBusClient {
	return conn.
		WithConsumer("created-user-event", createdUserEventMessage).
		WithConsumer("updated-user-event", updatedUserEventMessage).
		WithConsumer("deleted-user-event", deletedUserEventMessage).
		WitProducer("created-user-event", &events.CreatedUserEvent{}).
		WitProducer("updated-user-event", &events.UpdatedUserEvent{}).
		WitProducer("deleted-user-event", &events.DeletedUserEvent{}).
		Build()
}
