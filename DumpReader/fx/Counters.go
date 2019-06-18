package fx

import "github.com/prometheus/client_golang/prometheus"

type Counters struct {
	registry       *prometheus.Registry
	messageCounter *prometheus.CounterVec
}

func (self *Counters) MessageCounterInc(source string, feedName string) error {
	counter := self.messageCounter.WithLabelValues(source, feedName)
	counter.Inc()

	return nil
}

func (self *Counters) Gatherer() prometheus.Gatherer {
	return self.registry
}

func NewCounters() (*Counters, error) {
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

	return &Counters{
		registry:       registry,
		messageCounter: messagesProcessed,
	}, nil
}
