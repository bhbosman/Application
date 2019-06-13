package fx

import (
	"github.com/bhbosman/Application/DumpReader/MitchFeedReader"
	"go.uber.org/fx"
	"log"
)

func ProvideCounters() fx.Option {
	return fx.Provide(func(logger *log.Logger) (ICounters, MitchFeedReader.IFeedCounter, error) {
		logger.Println("Create Counters")
		counters, err := NewCounters()
		if err != nil {
			return nil, nil, err
		}
		return counters, counters, nil
	})
}




