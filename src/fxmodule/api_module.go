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
	conn message_bus.IMessageBus, config *server.Configuration,
	createdUserEventMessage *consumers.CreatedUserEventMessage,
	updatedUserEventMessage *consumers.UpdatedUserEventMessage,
	deletedUserEventMessage *consumers.DeletedUserEventMessage,
) message_bus.IMessageBusClient {
	return conn.
		WithConsumer(config.Kafka.Topics.CreatedUser, createdUserEventMessage).
		WithConsumer(config.Kafka.Topics.UpdatedUser, updatedUserEventMessage).
		WithConsumer(config.Kafka.Topics.DeletedUser, deletedUserEventMessage).
		WitProducer(config.Kafka.Topics.CreatedUser, &events.CreatedUserEvent{}).
		WitProducer(config.Kafka.Topics.UpdatedUser, &events.UpdatedUserEvent{}).
		WitProducer(config.Kafka.Topics.DeletedUser, &events.DeletedUserEvent{}).
		Build()
}
