package fx

import (
	"context"
	"fmt"
	"go.uber.org/fx"
	"log"
)

func ProvideApplicationContext(ctx context.Context, cancel context.CancelFunc) fx.Option {
	if ctx == nil || cancel == nil {
		return fx.Error(fmt.Errorf("context and cancel must be assigned"))
	}
	return fx.Provide(func(logger *log.Logger, lifecycle fx.Lifecycle) (context.Context, context.CancelFunc, error) {
		lifecycle.Append(fx.Hook{
			OnStart: nil,
			OnStop: func(stopContext context.Context) error {
				cancel()
				return nil
			},
		})
		return ctx, cancel, nil
	})
}

