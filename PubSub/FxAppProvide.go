package PubSub

import (
	"github.com/bhbosman/Application/UniqueNumber"
	"go.uber.org/fx"
	"log"
)

func FxAppProvide() fx.Option {
	return fx.Provide(
		func(logger *log.Logger, uniqueNumberGenerator UniqueNumber.IGenerator) (IPublisher, error) {
			publisher, err := NewPublisher(logger, uniqueNumberGenerator)
			if err != nil {
				return nil, err
			}
			return publisher, nil
		})
}
