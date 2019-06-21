package main

import (
	"context"
	"go.uber.org/fx"
)

type FxInvokeSymbolManagerInput struct {
	fx.In
	Lifecycle fx.Lifecycle
	Service   *fx.App `name:"SymbolManager"`
}

func FxAppInvokeSymbolManager() fx.Option {
	return fx.Invoke(
		func(inputData FxInvokeSymbolManagerInput) error {
			inputData.Lifecycle.Append(fx.Hook{
				OnStart: func(startContext context.Context) error {
					return inputData.Service.Start(startContext)
				},
				OnStop: func(stopContext context.Context) error {
					return inputData.Service.Stop(stopContext)
				},
			})
			return nil
		})
}

