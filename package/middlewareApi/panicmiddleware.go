package middlewareApi

import (
	"net/http"
	"strconv"
	"time"

	"github.com/go-park-mail-ru/2020_2_JMickhs/package/metrics"
	"github.com/go-park-mail-ru/2020_2_JMickhs/package/serverError"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func NewPanicMiddleware(metrics *metrics.PromMetrics) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			reqTime := time.Now()
			defer func() {
				if err := recover(); err != nil {
					respTime := time.Since(reqTime)

					metrics.Hits.WithLabelValues(strconv.Itoa(serverError.ServerInternalError), req.URL.String(), req.Method).Inc()
					metrics.Timings.WithLabelValues(strconv.Itoa(serverError.ServerInternalError), req.URL.String(), req.Method).
						Observe(respTime.Seconds())

					logrus.Error(err)
					http.Error(w, "Internal server error", http.StatusInternalServerError)

				}
			}()
			next.ServeHTTP(w, req)
		})
	}
}
