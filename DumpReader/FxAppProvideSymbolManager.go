package main

import (
	"go.uber.org/fx"
	"log"
)

func FxAppProvideSymbolManager() fx.Option {
	return fx.Provide(
		func(logger *log.Logger, mitchMessageHandlerRegistrar IMitchMessageHandlerRegistrar) (*SymbolDirectoryManager, error) {
			symbolDirectoryManager := NewSymbolDirectoryManager(logger)
			err := mitchMessageHandlerRegistrar.RegisterFeed(symbolDirectoryManager)
			if err != nil {
				return nil, err
			}
			return symbolDirectoryManager, nil
		})
}

func FxAppProvideTimerServiceManager() fx.Option {
	return fx.Provide(
		func(logger *log.Logger, mitchMessageHandlerRegistrar IMitchMessageHandlerRegistrar) (*TimeServiceManager, error) {
			symbolDirectoryManager := NewTimeServiceManager(logger)
			err := mitchMessageHandlerRegistrar.RegisterFeed(symbolDirectoryManager)
			if err != nil {
				return nil, err
			}
			return symbolDirectoryManager, nil
		})
}

