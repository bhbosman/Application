package main

import "github.com/prometheus/client_golang/prometheus"

type PerformanceCounters struct {
	registry       *prometheus.Registry
	messageCounter *prometheus.CounterVec
}

func (self *PerformanceCounters) MessageCounterInc(source string, feedName string) error {
	counter := self.messageCounter.WithLabelValues(source, feedName)
	counter.Inc()

	return nil
}

func (self *PerformanceCounters) Gatherer() prometheus.Gatherer {
	return self.registry
}

func NewCounters() (*PerformanceCounters, error) {
	registry := prometheus.NewRegistry()
	messagesProcessed := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace:   "",
			Subsystem:   "",
			Name:        "messages_processed_counter",
			Help:        "",
			ConstLabels: nil,
		},
		[]string{"Source", "FeeName"})

	registerError := registry.Register(messagesProcessed)
	if registerError != nil {
		return nil, registerError
	}

	return &PerformanceCounters{
		registry:       registry,
		messageCounter: messagesProcessed,
	}, nil
}
