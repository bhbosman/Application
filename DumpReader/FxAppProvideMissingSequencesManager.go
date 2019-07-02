package main

import (
	"github.com/bhbosman/Application/Managers/MissingSequence"
	"go.uber.org/fx"
	"log"
)

func FxAppProvideMissingSequencesManager() fx.Option {
	return fx.Provide(func(logger *log.Logger) (MissingSequence.IMissingSequencesManager, error) {
		return NewMissingSequencesManager(logger), nil
	})
}
