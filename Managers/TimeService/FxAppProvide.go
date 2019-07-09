package TimeService

import (
	"github.com/bhbosman/Application/Managers"
	"github.com/bhbosman/Application/PubSub"
	"go.uber.org/fx"
	"log"
)

func FxAppProvide() fx.Option {
	type ReturnType struct {
		fx.Out
		TimeServiceManager Managers.IMitchDataProcessor `name:"TimeServiceManager"`
	}
	return fx.Provide(
		func(logger *log.Logger, publisher PubSub.IPublisher) (ReturnType, error) {
			nextHandler, err := NewTimeServiceManager(logger)
			if err != nil {
				return ReturnType{}, err
			}
			handler := Managers.NewMitchDataProcessorChannelWrapper(nextHandler, publisher)
			return ReturnType{
				TimeServiceManager: handler,
			}, nil

		})
}
