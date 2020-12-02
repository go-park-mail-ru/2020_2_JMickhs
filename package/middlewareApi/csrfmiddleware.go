package middlewareApi

import (
	"context"
	"errors"
	"net/http"

	packageConfig "github.com/go-park-mail-ru/2020_2_JMickhs/package/configs"

	sessionService "github.com/go-park-mail-ru/2020_2_JMickhs/package/proto/sessions"

	"github.com/go-park-mail-ru/2020_2_JMickhs/package/clientError"
	customerror "github.com/go-park-mail-ru/2020_2_JMickhs/package/error"
	"github.com/go-park-mail-ru/2020_2_JMickhs/package/logger"
	"github.com/go-park-mail-ru/2020_2_JMickhs/package/responses"

	"github.com/gorilla/mux"
)

type CsrfMiddleware struct {
	CsrfUsecase sessionService.AuthorizationServiceClient
	Log         *logger.CustomLogger
}

func NewCsrfMiddleware(csrf sessionService.AuthorizationServiceClient, log *logger.CustomLogger) CsrfMiddleware {
	return CsrfMiddleware{
		CsrfUsecase: csrf,
		Log:         log,
	}
}

func (m *CsrfMiddleware) CSRFCheck() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			sid, ok := ctx.Value(packageConfig.SessionID).(string)
			if !ok {
				m.Log.Error(customerror.NewCustomError(errors.New("can't get sessId"), clientError.BadRequest, 1))
				ctx = context.WithValue(ctx, packageConfig.CorrectToken, false)
				next.ServeHTTP(w, r.WithContext(ctx))
				return
			}
			CSRFToken := r.Header.Get("X-Csrf-Token")
			res, err := m.CsrfUsecase.CheckCsrfToken(r.Context(), &sessionService.CsrfTokenCheck{SessionID: sid, Token: CSRFToken})

			if err != nil || !res.Result {
				m.Log.Error(customerror.NewCustomError(errors.New("can't get token from redis"), 0, 1))
				ctx = context.WithValue(ctx, packageConfig.CorrectToken, false)
				next.ServeHTTP(w, r.WithContext(ctx))
				return
			}
			ctx = context.WithValue(ctx, packageConfig.CorrectToken, true)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func CheckCSRFOnHandler(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			token, ok := r.Context().Value(packageConfig.CorrectToken).(bool)
			if !token || !ok {
				responses.SendErrorResponse(w, clientError.Forbidden)
				return
			}
			next.ServeHTTP(w, r)
		})
}
