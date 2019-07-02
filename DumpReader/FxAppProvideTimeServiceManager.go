package main

import (
	"github.com/bhbosman/Application/Managers"
	"github.com/bhbosman/Application/Managers/TimeService"
	"go.uber.org/fx"
	"log"
)

func FxAppProvideTimeServiceManager() fx.Option {
	type ReturnType struct {
		fx.Out
		TimeServiceManager Managers.IMitchDataProcessor `name:"TimeServiceManager"`
	}
	return fx.Provide(
		func(logger *log.Logger, mitchMessageHandlerRegistrar IMitchMessageHandlerRegistrar) (ReturnType, error) {
			nextHandler, err := TimeService.NewTimeServiceManager(logger)
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
				TimeServiceManager: handler,
			}, nil

		})
}
