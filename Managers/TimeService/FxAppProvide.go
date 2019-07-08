package TimeService

import (
	"github.com/bhbosman/Application/Managers"
	"go.uber.org/fx"
	"log"
)

func FxAppProvide() fx.Option {
	type ReturnType struct {
		fx.Out
		TimeServiceManager Managers.IMitchDataProcessor `name:"TimeServiceManager"`
	}
	return fx.Provide(
		func(logger *log.Logger, mitchMessageHandlerRegistrar Managers.IMitchMessageHandlerRegistrar) (ReturnType, error) {
			nextHandler, err := NewTimeServiceManager(logger)
			if err != nil {
				return ReturnType{}, err
			}

			handler, err := Managers.NewMitchDataProcessorChannelWrapper(nextHandler)
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
