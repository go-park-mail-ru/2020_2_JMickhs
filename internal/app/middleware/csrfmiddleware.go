package middlewareApi

import (
	"context"
	"github.com/go-park-mail-ru/2020_2_JMickhs/configs"
	csrfUsecase "github.com/go-park-mail-ru/2020_2_JMickhs/internal/app/csrf/usecase"
	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/pkg/clientError"
	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/pkg/logger"
	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/pkg/responses"
	"github.com/gorilla/mux"
	"net/http"
)

type CsrfMiddleware struct {
	CsrfUc csrfUsecase.CsrfUsecase
	Log         *logger.CustomLogger
}


func NewCsrfMiddleware(csrf csrfUsecase.CsrfUsecase,log *logger.CustomLogger) CsrfMiddleware {
	return CsrfMiddleware{
		CsrfUc: csrf,
		Log: log,
	}
}

func (m *CsrfMiddleware) CSRFCheck() mux.MiddlewareFunc{
	return func(next http.Handler) http.Handler{
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
				ctx := r.Context()
				sid, ok := ctx.Value(configs.SessionID).(string)
				if !ok {
					ctx = context.WithValue(ctx, configs.CorrectToken, false)
					next.ServeHTTP(w, r.WithContext(ctx))
					return
				}
				CSRFToken := r.Header.Get("X-Csrf-Token")
				ok, err := m.CsrfUc.CheckToken(sid, CSRFToken)

				if err != nil || !ok {
					ctx = context.WithValue(ctx, configs.CorrectToken, false)
					next.ServeHTTP(w, r.WithContext(ctx))
					return
				}
				ctx = context.WithValue(ctx, configs.CorrectToken, true)
				next.ServeHTTP(w, r.WithContext(ctx))
			})
		}
}


func CheckCSRFOnHandle(next http.HandlerFunc) http.HandlerFunc {
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