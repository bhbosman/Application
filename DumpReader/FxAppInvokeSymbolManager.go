package main

import (
	"context"
	"go.uber.org/fx"
)

func FxAppInvokeSymbolManager() fx.Option {
	return fx.Invoke(func(lc fx.Lifecycle, symbolDirectoryManager *SymbolDirectoryManager) error {
		var ch chan IMessageServiceItem = nil
		lc.Append(fx.Hook{
			OnStart: func(startContext context.Context) error {
				ch = make(chan IMessageServiceItem, 1024)
				symbolDirectoryManager.ch = ch

				go func() {
					for message := range ch {
						func(message IMessageServiceItem) {
							defer func() {
								err := message.Done()
								if err != nil {

								}
							}()
							err := symbolDirectoryManager.processMessage(message)
							if err != nil {

							}
						}(message)
					}
				}()
				return nil
			},
			OnStop: func(stopContext context.Context) error {
				symbolDirectoryManager.ch = nil
				close(ch)
				return nil
			},
		})
		return nil
	})
}



func FxAppInvokeTimerServiceManager() fx.Option {
	return fx.Invoke(func(lc fx.Lifecycle, timeServiceManager *TimeServiceManager) error {
		var ch chan IMessageServiceItem = nil
		lc.Append(fx.Hook{
			OnStart: func(startContext context.Context) error {
				ch = make(chan IMessageServiceItem, 1024)
				timeServiceManager.ch = ch

				go func() {
					for message := range ch {
						func(message IMessageServiceItem) {
							defer func() {
								err := message.Done()
								if err != nil {

								}
							}()
							err := timeServiceManager.processMessage(message)
							if err != nil {

							}
						}(message)
					}
				}()
				return nil
			},
			OnStop: func(stopContext context.Context) error {
				timeServiceManager.ch = nil
				close(ch)
				return nil
			},
		})
		return nil
	})
}
