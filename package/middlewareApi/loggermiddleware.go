package middlewareApi

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/go-park-mail-ru/2020_2_JMickhs/package/metrics"

	"github.com/go-park-mail-ru/2020_2_JMickhs/package/logger"
	"github.com/gorilla/mux"
)

type statusResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func NewStatusResponseWriter(w http.ResponseWriter) *statusResponseWriter {
	return &statusResponseWriter{w, http.StatusOK}
}

func (srw *statusResponseWriter) WriteHeader(code int) {
	srw.statusCode = code
	srw.ResponseWriter.WriteHeader(code)
}

func LoggerMiddleware(log *logger.CustomLogger, metrics *metrics.PromMetrics) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			rand.Seed(time.Now().UnixNano())
			id := fmt.Sprintf("%016x", rand.Int())[:5]

			log.StartReq(*req, id)
			start := time.Now()
			srw := NewStatusResponseWriter(w)
			next.ServeHTTP(srw, req)

			respTime := time.Since(start)
			log.EndReq(respTime.Microseconds(), req.Context())
			fmt.Println(srw.statusCode)
			if req.RequestURI != "/api/v1/metrics" {
				metrics.Hits.WithLabelValues(strconv.Itoa(srw.statusCode), req.URL.String(), req.Method).Inc()
				metrics.Total.Add(1)
				metrics.Timings.WithLabelValues(strconv.Itoa(srw.statusCode), req.URL.String(), req.Method).
					Observe(respTime.Seconds())
			}
			w.WriteHeader(200)
		})
	}
}
