package middlewareApi

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/pkg/logger"
	"github.com/gorilla/mux"
)

func LoggerMiddleware(log *logger.CustomLogger) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			rand.Seed(time.Now().UnixNano())
			id := fmt.Sprintf("%016x", rand.Int())[:5]

			log.StartReq(*req, id)
			start := time.Now()

			next.ServeHTTP(w, req)

			log.EndReq(start, req.Context())
		})
	}
}
