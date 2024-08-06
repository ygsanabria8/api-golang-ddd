package main

import (
	"api.ddd/src/fxmodule"
	"go.uber.org/fx"
)

func main() {
	fx.New(fxmodule.ApiModule).Run()
}
