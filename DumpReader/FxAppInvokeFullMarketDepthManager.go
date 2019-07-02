package main

import (
	"context"
	"github.com/bhbosman/Application/Managers"
	"go.uber.org/fx"
	"log"
)

func FxAppInvokeFullMarketDepthManager() fx.Option {
	type InputType struct {
		fx.In
		FullMarketDepthManager Managers.IMitchDataProcessor `name:"FullMarketDepthManager"`
	}
	return fx.Invoke(
		func(lifecycle fx.Lifecycle, logger *log.Logger, input InputType) error {
			lifecycle.Append(fx.Hook{
				OnStart: nil,
				OnStop: func(stopContext context.Context) error {
					logger.Println("FullMarketDepthManager")
					return input.FullMarketDepthManager.Close()
				},
			})
			return nil
		})
}
