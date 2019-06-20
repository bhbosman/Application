package main

import (
	"context"
	"go.uber.org/fx"
)



type IMessageServiceItem interface {
	Message() interface{}
	Add()
	Done()


}
type IMessageService interface {
	Push(message IMessageServiceItem)
	Stop(ctx context.Context) error
	Start(ctx context.Context) error
}

type TopOfBookManager struct {
	fxApp *fx.App
	ch    chan IMessageServiceItem
}

func (self *TopOfBookManager) Push(message IMessageServiceItem) {
	ch := self.ch
	if ch != nil{
		message.Add()
		ch <- message
	}
}

func NewTopOfBookManager() *TopOfBookManager {

	topOfBookManager := &TopOfBookManager{
		fxApp: nil,
		ch:    nil,
	}

	fxApp := fx.New(
		fx.Invoke(func(lc fx.Lifecycle) error {
			var ch chan IMessageServiceItem = nil
			lc.Append(fx.Hook{
				OnStart: func(startContext context.Context) error {
					ch = make(chan IMessageServiceItem, 1024)
					topOfBookManager.ch = ch

					go func() {
						for message := range ch{
							func(message IMessageServiceItem){
								defer func() {
									message.Done()
								}()
								topOfBookManager.processMessage(message)

							}(message)
						}
					}()
					return nil
				},
				OnStop: func(stopContext context.Context) error {
					topOfBookManager.ch = nil
					close(ch)
					return nil
				},
			})
			return nil
		}))

	topOfBookManager.fxApp = fxApp
	return topOfBookManager
}

func (self *TopOfBookManager) Start(ctx context.Context) error {
	return self.fxApp.Start(ctx)

}

func (self *TopOfBookManager) Stop(ctx context.Context) error {
	return self.fxApp.Stop(ctx)
}

func (self *TopOfBookManager) processMessage(item IMessageServiceItem) {

}
