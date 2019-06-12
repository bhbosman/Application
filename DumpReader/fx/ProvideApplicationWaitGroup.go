package fx

import (
	"go.uber.org/fx"
	"sync"
)

func ProvideApplicationWaitGroup(wg *sync.WaitGroup) fx.Option {
	return fx.Provide(
		func() *sync.WaitGroup {
			return wg
		})
}

