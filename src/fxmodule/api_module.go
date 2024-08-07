package fxmodule

import (
	"api.ddd/src/api/controllers"
	"api.ddd/src/api/routes"
	"api.ddd/src/api/server"
	"go.uber.org/fx"
)

var ApiModule = fx.Module(
	"Api Module",
	fx.Provide(server.ProvideLogger),
	fx.Provide(server.ProvideGinServer),
	fx.Provide(controllers.NewController),
	fx.Invoke(server.ConfigureCors),
	fx.Invoke(routes.ConfigureRoutes),
)
