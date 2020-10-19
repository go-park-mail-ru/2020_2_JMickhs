package middlewareApi

import (
	"fmt"
	"net/http"
	"reflect"

	customerror "github.com/go-park-mail-ru/2020_2_JMickhs/internal/pkg/error"

	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/pkg/responses"

	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/pkg/logger"

	"github.com/go-park-mail-ru/2020_2_JMickhs/configs"
	"github.com/gorilla/mux"
)

func ErrorMiddleware(log *logger.CustomLogger) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			next.ServeHTTP(w, req)

			err, ok := req.Context().Value(configs.DeliveryError).(error)
			fmt.Println(reflect.TypeOf(req.Context().Value(configs.DeliveryError)))
			if !ok {
				fmt.Println("cant cast to error")
				return
			}
			log.LogError(req.Context(), err)
			responses.SendErrorResponse(w, customerror.ParseCode(err))
		})
	}
}
