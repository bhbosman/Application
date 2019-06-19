package main

import (
	"context"
	"fmt"
	"go.uber.org/fx"
	"log"
	"sync"
)

type IApplicationContext interface {
	Context() context.Context
	Cancel() context.CancelFunc
	WaitGroup() *sync.WaitGroup
}

type ApplicationContext struct {
	context   context.Context
	cancel    context.CancelFunc
	waitGroup *sync.WaitGroup
}

func (self *ApplicationContext) WaitGroup() *sync.WaitGroup {
	return self.waitGroup
}

func (self *ApplicationContext) Context() context.Context {
	return self.context
}

func (self *ApplicationContext) Cancel() context.CancelFunc {
	return self.cancel
}

func FxAppProvideApplicationContext() fx.Option {
	return fx.Provide(func(logger *log.Logger, lifecycle fx.Lifecycle) (IApplicationContext, error) {
		applicationContext := &ApplicationContext{
			context:   nil,
			cancel:    nil,
			waitGroup: nil,
		}
		lifecycle.Append(fx.Hook{
			OnStart: func(startContext context.Context) error {
				ctx, cancel := context.WithCancel(context.Background())
				wg := &sync.WaitGroup{}
				if ctx == nil || cancel == nil {
					return fmt.Errorf("context and cancel must be assigned")
				}
				applicationContext.context = ctx
				applicationContext.cancel = cancel
				applicationContext.waitGroup = wg
				return nil
			},
			OnStop: func(stopContext context.Context) error {
				applicationContext.cancel()
				return nil
			},
		})
		return applicationContext, nil
	})
}
