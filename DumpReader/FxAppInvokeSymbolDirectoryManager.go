package main

import (
	"context"
	"github.com/bhbosman/Application/Managers"
	"go.uber.org/fx"
	"log"
)

func FxAppInvokeSymbolDirectoryManager() fx.Option {
	type InputType struct {
		fx.In
		SymbolDirectoryManager Managers.IMitchDataProcessor `name:"SymbolDirectoryManager"`
	}
	return fx.Invoke(
		func(lifecycle fx.Lifecycle, logger *log.Logger, input InputType) error {
			lifecycle.Append(fx.Hook{
				OnStart: nil,
				OnStop: func(stopContext context.Context) error {
					logger.Println("SymbolDirectoryManager")
					return input.SymbolDirectoryManager.Close()
				},
			})
			return nil
		})
}
