package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

type PromMetrics struct {
	Hits    *prometheus.CounterVec
	Timings *prometheus.HistogramVec
}

func RegisterMetrics() *PromMetrics {
	var metrics PromMetrics

	metrics.Hits = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "hits",
	}, []string{"status", "path", "method"})

	metrics.Timings = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "timings",
		},
		[]string{"status", "path", "method"},
	)

	prometheus.MustRegister(metrics.Hits, metrics.Timings)

	return &metrics
}
