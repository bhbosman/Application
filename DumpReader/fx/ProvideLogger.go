package fx

import (
	"go.uber.org/fx"
	"log"
)

func ProvideLogger(logger *log.Logger) fx.Option {
	return fx.Provide(
		func() (*log.Logger, error) {
			logger := logger
			logger.Printf("Create logger\n")
			return logger, nil
		})
}

