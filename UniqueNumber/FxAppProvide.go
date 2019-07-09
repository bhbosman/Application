package UniqueNumber

import (
	"go.uber.org/fx"
	"log"
)

func FxAppProvide() fx.Option {
	return fx.Provide(
		func(logger *log.Logger) (IGenerator, error) {
			generator, err := NewUniqueNumberGenerator(logger)
			if err != nil {
				return nil, err
			}
			return generator, nil
		})
}
