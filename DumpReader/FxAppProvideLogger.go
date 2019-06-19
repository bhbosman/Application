package main

import (
	"go.uber.org/fx"
	"log"
)

func FxAppProvideApplicationLogger(logger *log.Logger) fx.Option {
	return fx.Provide(
		func() (*log.Logger, error) {
			logger := logger
			logger.Printf("Create logger\n")
			return logger, nil
		})
}
