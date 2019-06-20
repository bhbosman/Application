package main

import (
	"github.com/bhbosman/Application/Streams"
	"go.uber.org/fx"
	"log"
)

func FxAppProvideMitchFeedReader() fx.Option {
	return fx.Provide(func(logger *log.Logger, currentOpenFile ICurrentOpenFile, feedCounter IFeedCounter) (*MitchFeedReader, error) {
		mitchDataHandler, err := NewReadMitchDataHandler(Streams.NewMitchReaderFactory())
		if err != nil {
			logger.Printf("Error creating mitch handler. Error: %v\n", err)
			return nil, err
		}

		reader, err := NewMitchFeedReader(logger, currentOpenFile, mitchDataHandler, feedCounter)
		if err != nil {
			logger.Printf("Could not create Mitch Reader. Error: %v\n", err)
			return nil, err
		}
		return reader, nil
	})
}

