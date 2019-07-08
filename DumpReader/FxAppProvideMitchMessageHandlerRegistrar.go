package main

import (
	"github.com/bhbosman/Application/Managers"
	"go.uber.org/fx"
	"log"
)

func FxAppProvideMitchMessageHandlerRegistrar() fx.Option {
	return fx.Provide(
		func(logger *log.Logger) (Managers.IMitchMessageHandlerRegistrar, error) {
			return Managers.NewMitchMessageHandlerRegistrar(logger), nil
		})
}
