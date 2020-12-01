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

func LoggerMiddleware(log *logger.CustomLogger, metrics *metrics.PromMetrics) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			rand.Seed(time.Now().UnixNano())
			id := fmt.Sprintf("%016x", rand.Int())[:5]

			log.StartReq(*req, id)
			start := time.Now()

			next.ServeHTTP(w, req)

			respTime := time.Since(start)
			log.EndReq(respTime.Microseconds(), req.Context())
			if req.URL.Path != "metrics" {
				metrics.Hits.WithLabelValues(strconv.Itoa(http.StatusOK), req.URL.String(), req.Method).Inc()
				metrics.Timings.WithLabelValues(strconv.Itoa(http.StatusOK), req.URL.String(), req.Method).
					Observe(respTime.Seconds())
			}
		})
	}
}
