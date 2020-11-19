package middlewareApi

import (
	"context"
	http "net/http"

	"github.com/go-park-mail-ru/2020_2_JMickhs/JMickhs_main/proto/sessions"

	"github.com/go-park-mail-ru/2020_2_JMickhs/JMickhs_main/configs"
	"github.com/go-park-mail-ru/2020_2_JMickhs/JMickhs_main/internal/app/user"
	"github.com/go-park-mail-ru/2020_2_JMickhs/package/clientError"
	customerror "github.com/go-park-mail-ru/2020_2_JMickhs/package/error"
	"github.com/go-park-mail-ru/2020_2_JMickhs/package/logger"
	"github.com/gorilla/mux"
)

type SessionMidleware struct {
	SessionDelivery sessions.AuthorizationServiceClient
	UserUseCase     user.Usecase
	log             *logger.CustomLogger
}

func NewSessionMiddleware(su sessions.AuthorizationServiceClient, uuc user.Usecase, log *logger.CustomLogger) SessionMidleware {
	return SessionMidleware{
		SessionDelivery: su,
		UserUseCase:     uuc,
		log:             log,
	}
}

func (u *SessionMidleware) SessionMiddleware() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			c, err := r.Cookie("session_token")

			if err != nil {
				err = customerror.NewCustomError(err, clientError.BadRequest, 1)
				u.log.Info(err.Error())
				next.ServeHTTP(w, r)
				return
			}
			if c != nil {
				sessionToken := c.Value
				id, err := u.SessionDelivery.GetIDBySession(r.Context(), &sessions.SessionID{SessionID: sessionToken})
				if err != nil {
					u.log.LogError(r.Context(), err)
					next.ServeHTTP(w, r)
					return
				}
				user, err := u.UserUseCase.GetUserByID(int(id.UserID))
				if err != nil {
					u.log.LogError(r.Context(), err)
					next.ServeHTTP(w, r)
					return
				}
				ctx := context.WithValue(r.Context(), configs.RequestUser, user)
				ctx = context.WithValue(ctx, configs.SessionID, sessionToken)
				r = r.WithContext(ctx)
			}
			next.ServeHTTP(w, r)
		})
	}
}
