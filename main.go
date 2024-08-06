package main

import (
	"api.ddd/src/api/server"
	"api.ddd/src/fxmodule"
	"go.uber.org/fx"
)

func main() {
	logger := server.ProvideLogger().GetFxLogger()
	fx.New(
		fxmodule.ApiModule,
		fx.Logger(logger),
	).Run()
}
