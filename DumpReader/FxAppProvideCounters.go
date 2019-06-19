package main

import (
	"go.uber.org/fx"
	"log"
)

func FxAppProvideCounters() fx.Option {
	return fx.Provide(func(logger *log.Logger) (IPerformanceCounters, IFeedCounter, error) {
		logger.Println("Create PerformanceCounters")
		counters, err := NewCounters()
		if err != nil {
			return nil, nil, err
		}
		return counters, counters, nil
	})
}
