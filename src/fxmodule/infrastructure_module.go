package fxmodule

import (
	finder "api.ddd/src/infrastructure/finder/user"
	"api.ddd/src/infrastructure/register_services"
	repository "api.ddd/src/infrastructure/repository/user"
	"go.uber.org/fx"
)

var InfrastructureModule = fx.Module(
	"Infrastructure Module",
	fx.Provide(register_services.ProvideMongoClient),
	fx.Provide(register_services.ProvideSqlClient),
	fx.Provide(
		fx.Annotated{
			Name:   "NoSqlFinder",
			Target: finder.NewMongoUserFinder,
		},
		fx.Annotated{
			Name:   "SqlFinder",
			Target: finder.NewMySqlUserFinder,
		},
	),
	fx.Provide(
		fx.Annotated{
			Name:   "NoSqlRepository",
			Target: repository.NewMongoUserRepository,
		},
		fx.Annotated{
			Name:   "SqlRepository",
			Target: repository.NewUserMySqlRepository,
		},
	),
)
