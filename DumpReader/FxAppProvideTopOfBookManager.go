package main

import (
	"context"
	"go.uber.org/fx"
)

type FxAppProvideTopOfBookManagerOutput struct {
	fx.Out
	MessageService IMessageService `name:"TopOfBookManager"`
}

func FxAppProvideTopOfBookManager() fx.Option {
	return fx.Provide(
		func() (FxAppProvideTopOfBookManagerOutput, error) {
			topOfBookManager := NewTopOfBookManager()
			return FxAppProvideTopOfBookManagerOutput{
				MessageService: topOfBookManager,
			}, nil
	})
}

type FxAppInvokeTopOfBookManagerInput struct {
	fx.In
	LifeCycle fx.Lifecycle
	MessageService IMessageService `name:"TopOfBookManager"`
}

func FxAppInvokeTopOfBookManager() fx.Option {
	return fx.Invoke(
		func(inputData FxAppInvokeTopOfBookManagerInput) error {
			inputData.LifeCycle.Append(fx.Hook{
				OnStart: func(startContext context.Context) error {
					return inputData.MessageService.Start(startContext)
				},
				OnStop: func(stopContext context.Context) error {
					return inputData.MessageService.Stop(stopContext)
				},
			})
			return nil
		})
}



