package fxmodule

import (
	"api.ddd/src/api/server"
	"go.uber.org/fx"
)

var ApiModule = fx.Options(
	fx.Provide(server.ProviderGinServer),
	fx.Invoke(server.StartGinServer),
)
