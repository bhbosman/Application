package fx

import (
	"context"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/fx"
	"log"
	"net/http"
)

func InvokeCreatePrometheusService() fx.Option {
	return fx.Invoke(
		func(
			lifecycle fx.Lifecycle,
			logger *log.Logger,
			counters ICounters,
			applicationContext context.Context,
			applicationCancel context.CancelFunc) error {

			logger.Printf("Creating http service for prometheus\n")
			server := &http.Server{
				Addr: ":12345",
				Handler: promhttp.HandlerFor(counters.Gatherer(), promhttp.HandlerOpts{
					ErrorLog:            logger,
					ErrorHandling:       promhttp.ContinueOnError,
					DisableCompression:  false,
					MaxRequestsInFlight: 0,
					Timeout:             0,
				})}
			lifecycle.Append(fx.Hook{
				OnStart: func(startContext context.Context) error {
					go func() {
						logger.Printf("Starting prometheus http\n")
						_ = server.ListenAndServe()
					}()

					return nil
				},
				OnStop: func(stopContext context.Context) error {
					logger.Printf("Stopping prometheus http\n")
					err := server.Close()
					return err
				},
			})
			return nil
		})

}


