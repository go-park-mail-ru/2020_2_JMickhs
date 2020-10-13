package middlewareApi

import (
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"

	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/logger"
	"github.com/gorilla/mux"
)

func LoggerMiddleware(log *logger.CustomLogger) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			id := uuid.NewV4().String()

			log.StartReq(*req, id)
			start := time.Now()

			next.ServeHTTP(w, req)

			log.EndReq(start, req.Context())
		})
	}
}
