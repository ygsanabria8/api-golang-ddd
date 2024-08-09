package fxmodule

import (
	"api.ddd/src/api/server"
	"go.uber.org/fx"
)

var InitModule = fx.Options(
	InfrastructureModule,
	ApiModule,
	fx.Invoke(server.StartGinServer),
)
