package main

import (
	"go.uber.org/fx"
	"log"
)

func FxAppProvideMissingSequencesManager() fx.Option {
	return fx.Provide(func(logger *log.Logger) (IMissingSequencesManager, error) {
		return NewMissingSequencesManager(logger), nil
	})
}

