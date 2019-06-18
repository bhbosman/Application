package fx

import "github.com/prometheus/client_golang/prometheus"

type ICounters interface {
	Gatherer() prometheus.Gatherer
}
