package fxmodule

import (
	"api.ddd/src/api/server"
	"go.uber.org/fx"
)

var InitModule = fx.Options(
	fx.Provide(server.SetUpConfig),
	DomainModule,
	InfrastructureModule,
	ApplicationModule,
	ApiModule,
	fx.Invoke(server.StartGinServer),
)
