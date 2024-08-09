package fxmodule

import (
	"api.ddd/src/infrastructure/finder/mongo"
	"api.ddd/src/infrastructure/register_services"
	repository "api.ddd/src/infrastructure/repository/mongo"
	"go.uber.org/fx"
)

var InfrastructureModule = fx.Module(
	"Infrastructure Module",
	fx.Provide(register_services.ProvideMongoClient),
	fx.Provide(register_services.ProvideSqlClient),
	fx.Provide(finder.NewMongoFinder),
	fx.Provide(repository.NewMongoRepository),
)
