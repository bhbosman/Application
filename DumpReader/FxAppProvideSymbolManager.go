package main

import (
	"go.uber.org/fx"
	"log"
)

type FxProvideSymbolManagerInput struct {
	fx.In
	Logger                       *log.Logger
	MitchMessageHandlerRegistrar IMitchMessageHandlerRegistrar
}

type FxProvideSymbolManagerOutput struct {
	fx.Out
	Service *fx.App `name:"SymbolManager"`
}

func FxAppProvideSymbolManager() fx.Option {
	return fx.Provide(func(inputData FxProvideSymbolManagerInput) (FxProvideSymbolManagerOutput, error) {

		return FxProvideSymbolManagerOutput{
			Service: NewSymbolDirectoryManagerFxApp(
				inputData.Logger,
				inputData.MitchMessageHandlerRegistrar),
		}, nil
	})
}
