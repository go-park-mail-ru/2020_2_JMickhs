package middlewareApi

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	RequestDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "service_request_duration",
		Buckets: prometheus.LinearBuckets(0.01, 0.01, 10),
	}, []string{"path", "method", "backend"})

	Hits = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "service_request_path_hits",
	}, []string{"path", "status", "backend"})
)

func CreatePrometheusMetricsMiddleware() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			template, err := mux.CurrentRoute(r).GetPathTemplate()
			if err != nil {
				template = r.URL.Path
			}

			defer func() {
				statusCode := 200
				Hits.With(
					prometheus.Labels{"path": template, "status": strconv.Itoa(statusCode), "backend": "1"},
				).Inc()
			}()

			timer := prometheus.NewTimer(RequestDuration.With(
				prometheus.Labels{"path": template, "method": r.Method, "backend": "1"},
			))
			defer timer.ObserveDuration()

			next.ServeHTTP(w, r)
		})
	}
}
