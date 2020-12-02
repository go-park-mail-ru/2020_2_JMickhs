package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

type PromMetrics struct {
	Total     prometheus.Counter
	Hits      *prometheus.CounterVec
	HitsError *prometheus.CounterVec
	Timings   *prometheus.HistogramVec
}

func RegisterMetrics() *PromMetrics {
	var metrics PromMetrics

	metrics.Total = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "foo_total",
	})

	metrics.Hits = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "hits",
	}, []string{"status", "path", "method"})

	metrics.HitsError = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "hitsError",
	}, []string{"status", "path", "method"})

	metrics.Timings = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "timings",
		},
		[]string{"status", "path", "method"},
	)

	prometheus.MustRegister(metrics.Hits, metrics.Timings, metrics.Total, metrics.HitsError)

	return &metrics
}
