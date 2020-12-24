package middlewareApi

import (
	"bufio"
	"errors"
	"fmt"
	"math/rand"
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"
	"unicode"

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
}

func (srw *statusResponseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	h, ok := srw.ResponseWriter.(http.Hijacker)
	if !ok {
		return nil, nil, errors.New("hijack not supported")
	}
	return h.Hijack()
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
			if req.RequestURI != "/api/v1/metrics" {
				vars := mux.Vars(req)
				var url string
				if len(vars) > 0 {
					url = strings.TrimRightFunc(req.URL.Path, func(r rune) bool {
						return unicode.IsNumber(r)
					})
					url = url[:len(url)-1]
				} else {
					url = req.URL.Path
				}
				if srw.statusCode != 500 {
					metrics.Hits.WithLabelValues(strconv.Itoa(srw.statusCode), url, req.Method).Inc()
					metrics.Total.Add(1)
					metrics.Timings.WithLabelValues(strconv.Itoa(srw.statusCode), url, req.Method).
						Observe(respTime.Seconds())
				} else {
					metrics.HitsError.WithLabelValues(strconv.Itoa(srw.statusCode), url, req.Method).Inc()
				}
			}
		})
	}
}
