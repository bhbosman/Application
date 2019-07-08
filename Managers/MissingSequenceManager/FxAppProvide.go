package MissingSequenceManager

import (
	"github.com/bhbosman/Application/Managers"
	"github.com/bhbosman/Application/PubSub"
	"go.uber.org/fx"
	"log"
)

func FxAppProvide() fx.Option {
	type ReturnType struct {
		fx.Out
		FullMarketDepthManager Managers.IMitchDataProcessor `name:"MissingSequenceManager"`
	}
	return fx.Provide(
		func(publisher PubSub.IPublisher, logger *log.Logger, mitchMessageHandlerRegistrar Managers.IMitchMessageHandlerRegistrar) (ReturnType, error) {
			nextHandler, err := NewManager(logger, publisher)
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
				FullMarketDepthManager: handler,
			}, nil
		})
}
