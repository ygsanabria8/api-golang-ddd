package fxmodule

import (
	"api.ddd/src/infrastructure/register_services"
	"go.uber.org/fx"
)

var InfrastructureModule = fx.Module(
	"Infrastructure Module",
	fx.Provide(register_services.ProvideMongoClient),
	fx.Provide(register_services.ProvideSqlClient),
)
