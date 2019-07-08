package PubSub

import (
	"context"
	"go.uber.org/fx"
	"log"
)

func FxAppInvoke() fx.Option {
	return fx.Invoke(
		func(lifecycle fx.Lifecycle, logger *log.Logger, input IPublisher) error {
			lifecycle.Append(fx.Hook{
				OnStart: func(startContext context.Context) error {
					return nil
				},
				OnStop: func(stopContext context.Context) error {
					return input.Close()
				},
			})
			return nil
		})
}
