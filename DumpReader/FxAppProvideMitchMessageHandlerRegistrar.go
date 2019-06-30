package main

import (
	"go.uber.org/fx"
	"log"
)

func FxAppProvideMitchMessageHandlerRegistrar() fx.Option {
	return fx.Provide(
		func(logger *log.Logger) (IMitchMessageHandlerRegistrar, error) {
			return NewMitchMessageHandlerRegistrar(logger), nil
		})
}

