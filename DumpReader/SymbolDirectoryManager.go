package main

import (
	"context"
	"github.com/bhbosman/Application/MitchFiles/GeneratedFiles"
	"go.uber.org/fx"
	"log"
)

type SymbolDirectoryManager struct {
	ch     chan IMessageServiceItem
	logger *log.Logger
}

func NewSymbolDirectoryManager(logger *log.Logger) *SymbolDirectoryManager {
	return &SymbolDirectoryManager{
		logger: logger,
	}
}

func (self *SymbolDirectoryManager) DeclareInterestInMessages() ([]byte, error) {
	return []byte{GeneratedFiles.SymbolDirectoryMessageMessageType}, nil
}

func (self *SymbolDirectoryManager) Push(message IMessageServiceItem) {
	ch := self.ch
	if ch != nil {
		err := message.Add()
		if err != nil {
			return
		}
		ch <- message
	}
}

func NewSymbolDirectoryManagerFxApp(logger *log.Logger, mitchMessageHandlerRegistrar IMitchMessageHandlerRegistrar) *fx.App {
	return fx.New(
		FxAppProvideApplicationLogger(logger),
		FxAppProvideFxAppOverrideLogger(logger),
		fx.Provide(func(logger *log.Logger) (*SymbolDirectoryManager, error) {
			symbolDirectoryManager  := NewSymbolDirectoryManager(logger)
			err := mitchMessageHandlerRegistrar.RegisterFeed(symbolDirectoryManager)
			if err != nil {
				return nil, err
			}
			return symbolDirectoryManager, nil
		}),
		fx.Invoke(func(lc fx.Lifecycle, symbolDirectoryManager *SymbolDirectoryManager) error {
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
									symbolDirectoryManager.processMessage(message)

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
		}))

}


func (self *SymbolDirectoryManager) HandleSymbolDirectoryMessage(item *GeneratedFiles.SymbolDirectoryMessage) error {
	return nil
}

func (self *SymbolDirectoryManager) processMessage(item IMessageServiceItem) error {
	msg, err := item.Message()
	if err != nil {
		return err
	}
	switch v := msg.(type) {
	case *GeneratedFiles.SymbolDirectoryMessage:
		return self.HandleSymbolDirectoryMessage(v)
		break
	}
	return nil
}
