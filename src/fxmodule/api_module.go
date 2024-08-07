package fxmodule

import (
	"api.ddd/src/api/server"
	"go.uber.org/fx"
)

var ApiModule = fx.Module(
	"Api Module",
	fx.Provide(server.ProvideLogger),
	fx.Provide(server.ProvideGinServer),
	fx.Invoke(server.ConfigureCors),
)
