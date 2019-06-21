package main

import (
	"context"
	"go.uber.org/fx"
	"log"
)

type FxAppInvokeMitchFeedReaderInput struct {
	fx.In
	Lifecycle          fx.Lifecycle
	Logger             *log.Logger
	ApplicationContext IApplicationContext
	Reader             *MitchFeedProcessor
}

func FxAppInvokeMitchFeedProcessor() fx.Option {
	return fx.Invoke(
		func(inputData FxAppInvokeMitchFeedReaderInput) error {
			var lifeCycleContext context.Context
			var lifeCycleCancel context.CancelFunc
			inputData.Lifecycle.Append(fx.Hook{
				OnStart: func(startContext context.Context) error {
					lifeCycleContext, lifeCycleCancel = context.WithCancel(inputData.ApplicationContext.Context())
					err := inputData.Reader.Process(inputData.ApplicationContext.WaitGroup(), lifeCycleContext, "file", "static")
					if err != nil {
						inputData.Logger.Printf("error: %v", err)
					}
					return err
				},
				OnStop: func(stopContext context.Context) error {
					lifeCycleCancel()
					err := inputData.Reader.Close()
					if err != nil {
						inputData.Logger.Printf("error: %v", err)
					}
					return err
				},
			})
			return nil
		})
}
