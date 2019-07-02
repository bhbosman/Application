package main

import (
	"github.com/bhbosman/Application/Managers"
	"github.com/bhbosman/Application/Managers/FullMarketDepth"
	"go.uber.org/fx"
	"log"
)

func FxAppProvideFullMarketDepthManager() fx.Option {
	type ReturnType struct {
		fx.Out
		FullMarketDepthManager Managers.IMitchDataProcessor `name:"FullMarketDepthManager"`
	}
	return fx.Provide(
		func(logger *log.Logger, mitchMessageHandlerRegistrar IMitchMessageHandlerRegistrar) (ReturnType, error) {
			nextHandler, err := FullMarketDepth.NewManager(logger)
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
				FullMarketDepthManager: handler,
			}, nil
		})
}
