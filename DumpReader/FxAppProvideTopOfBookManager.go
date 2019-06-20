package main

import (
	"go.uber.org/fx"
)

type FxAppProvideTopOfBookManagerOutput struct {
	fx.Out
	MessageService IMessageService `name:"TopOfBookManager"`
}

func FxAppProvideTopOfBookManager() fx.Option {
	return fx.Provide(
		func() (FxAppProvideTopOfBookManagerOutput, error) {
			topOfBookManager := NewTopOfBookManager()
			return FxAppProvideTopOfBookManagerOutput{
				MessageService: topOfBookManager,
			}, nil
	})
}




