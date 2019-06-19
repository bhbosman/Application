package main

import "github.com/prometheus/client_golang/prometheus"

type IPerformanceCounters interface {
	Gatherer() prometheus.Gatherer
}
