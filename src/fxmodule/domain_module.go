package fxmodule

import (
	services "api.ddd/src/domain/services/user"
	"go.uber.org/fx"
)

var DomainModule = fx.Options(
	fx.Provide(services.NewUserService),
)
