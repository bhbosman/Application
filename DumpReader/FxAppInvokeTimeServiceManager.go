package main

import (
	"context"
	"github.com/bhbosman/Application/Managers"
	"go.uber.org/fx"
	"log"
)

func FxAppInvokeTimeServiceManager() fx.Option {
	type InputType struct {
		fx.In
		TimeServiceManager Managers.IMitchDataProcessor `name:"TimeServiceManager"`
	}
	return fx.Invoke(
		func(lifecycle fx.Lifecycle, logger *log.Logger, input InputType) error {
			lifecycle.Append(fx.Hook{
				OnStart: nil,
				OnStop: func(stopContext context.Context) error {
					logger.Println("Closing TimeServiceManager")
					return input.TimeServiceManager.Close()
				},
			})
			return nil
		})
}
