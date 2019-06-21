package main

import (
	"github.com/bhbosman/Application/Streams"
	"go.uber.org/fx"
	"log"
)

type FxAppProvideMitchFeedReaderInput struct {
	fx.In
	Logger                       *log.Logger
	CurrentOpenFile              ICurrentOpenFile
	FeedCounter                  IFeedCounter
	MitchMessageHandlerRegistrar IMitchMessageHandlerRegistrar
}

func FxAppProvideMitchFeedProcessor() fx.Option {
	return fx.Provide(func(inputData FxAppProvideMitchFeedReaderInput) (*MitchFeedProcessor, error) {
		mitchDataHandler, err := NewReadMitchDataHandler(Streams.NewMitchReaderFactory())
		if err != nil {
			inputData.Logger.Printf("Error creating mitch handler. Error: %v\n", err)
			return nil, err
		}

		reader, err := NewMitchFeedProcessor(
			inputData.Logger,
			inputData.CurrentOpenFile,
			mitchDataHandler,
			inputData.FeedCounter,
			inputData.MitchMessageHandlerRegistrar)
		if err != nil {
			inputData.Logger.Printf("Could not create Mitch Reader. Error: %v\n", err)
			return nil, err
		}
		return reader, nil
	})
}
