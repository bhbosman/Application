package main

import (
	"github.com/bhbosman/Application/Managers"
	"github.com/bhbosman/Application/Managers/SymbolService"
	"go.uber.org/fx"
	"log"
)

func FxAppProvideSymbolDirectoryManager() fx.Option {
	type ReturnType struct {
		fx.Out
		SymbolDirectoryManager Managers.IMitchDataProcessor `name:"SymbolDirectoryManager"`
	}
	return fx.Provide(
		func(logger *log.Logger, mitchMessageHandlerRegistrar IMitchMessageHandlerRegistrar) (ReturnType, error) {
			nextHandler, err := SymbolService.NewSymbolDirectoryManager(logger)
			if err != nil {
				return ReturnType{}, err
			}

			handler, err := NewMitchDataProcessorChannelWrapper(nextHandler)
			if err != nil {
				return ReturnType{}, err
			}
			err = mitchMessageHandlerRegistrar.RegisterFeed(handler)
			if err != nil {
				return ReturnType{}, err
			}
			return ReturnType{
				SymbolDirectoryManager: handler,
			}, nil
		})
}
