package main

import (
	"github.com/bhbosman/Application/MissingSequences"
	"go.uber.org/fx"
	"log"
)

func FxAppProvideMissingSequencesManager() fx.Option {
	return fx.Provide(func(logger *log.Logger) (MissingSequences.IMissingSequencesManager, error) {
		return MissingSequences.NewMissingSequencesManager(logger), nil
	})
}
