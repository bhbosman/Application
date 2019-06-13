package fx

import (
	"go.uber.org/fx"
	"log"
)

type loggerWrapper struct {
	applicationLogger *log.Logger
}

func (self *loggerWrapper) Printf(format string, v ...interface{}) {
	self.applicationLogger.Printf(format, v...)
}

func ProvideOverrideLogger(logger *log.Logger) fx.Option {
	return fx.Logger(&loggerWrapper{
		applicationLogger: logger,
	})
}

