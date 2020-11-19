package middlewareApi

import (
	"context"
	"errors"
	"net/http"

	"github.com/go-park-mail-ru/2020_2_JMickhs/JMickhs_main/configs"
	"github.com/go-park-mail-ru/2020_2_JMickhs/JMickhs_main/pkg/clientError"
	customerror "github.com/go-park-mail-ru/2020_2_JMickhs/JMickhs_main/pkg/error"
	"github.com/go-park-mail-ru/2020_2_JMickhs/JMickhs_main/pkg/logger"
	"github.com/go-park-mail-ru/2020_2_JMickhs/JMickhs_main/pkg/responses"

	"github.com/go-park-mail-ru/2020_2_JMickhs/JMickhs_main/internal/app/csrf"

	"github.com/gorilla/mux"
)

type CsrfMiddleware struct {
	CsrfUc csrf.Usecase
	Log    *logger.CustomLogger
}

func NewCsrfMiddleware(csrf csrf.Usecase, log *logger.CustomLogger) CsrfMiddleware {
	return CsrfMiddleware{
		CsrfUc: csrf,
		Log:    log,
	}
}

func (m *CsrfMiddleware) CSRFCheck() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			sid, ok := ctx.Value(configs.SessionID).(string)
			if !ok {
				m.Log.Error(customerror.NewCustomError(errors.New("can't get sessId"), 0, 1))
				ctx = context.WithValue(ctx, configs.CorrectToken, false)
				next.ServeHTTP(w, r.WithContext(ctx))
				return
			}
			CSRFToken := r.Header.Get("X-Csrf-Token")
			ok, err := m.CsrfUc.CheckToken(sid, CSRFToken)

			if err != nil || !ok {
				m.Log.Error(customerror.NewCustomError(errors.New("can't get token from redis"), 0, 1))
				ctx = context.WithValue(ctx, configs.CorrectToken, false)
				next.ServeHTTP(w, r.WithContext(ctx))
				return
			}
			ctx = context.WithValue(ctx, configs.CorrectToken, true)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func CheckCSRFOnHandler(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			token, ok := r.Context().Value(configs.CorrectToken).(bool)
			if !token || !ok {

				responses.SendErrorResponse(w, clientError.Forbidden)
				return
			}
			next.ServeHTTP(w, r)
		})
}
