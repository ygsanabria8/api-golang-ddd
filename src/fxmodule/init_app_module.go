package fxmodule

import (
	"api.ddd/src/api/server"
	"go.uber.org/fx"
)

var InitModule = fx.Options(
	ApiModule,
	fx.Invoke(server.StartGinServer),
)
