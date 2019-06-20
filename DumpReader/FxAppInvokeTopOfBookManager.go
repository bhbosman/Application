package main

import (
	"context"
	"go.uber.org/fx"
)

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
					return inputData.MessageService.(IService).StartService(startContext)
				},
				OnStop: func(stopContext context.Context) error {
					return inputData.MessageService.(IService).StopService(stopContext)
				},
			})
			return nil
		})
}



