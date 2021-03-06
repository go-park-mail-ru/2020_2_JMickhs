package middlewareApi

import (
	"context"
	http "net/http"

	userService "github.com/go-park-mail-ru/2020_2_JMickhs/package/proto/user"

	packageConfig "github.com/go-park-mail-ru/2020_2_JMickhs/package/configs"

	sessionService "github.com/go-park-mail-ru/2020_2_JMickhs/package/proto/sessions"

	"github.com/go-park-mail-ru/2020_2_JMickhs/package/clientError"
	customerror "github.com/go-park-mail-ru/2020_2_JMickhs/package/error"
	"github.com/go-park-mail-ru/2020_2_JMickhs/package/logger"
	"github.com/gorilla/mux"
)

type SessionMidleware struct {
	SessionDelivery sessionService.AuthorizationServiceClient
	UserDelivery    userService.UserServiceClient
	log             *logger.CustomLogger
}

func NewSessionMiddleware(sessionDelivery sessionService.AuthorizationServiceClient, userDelivery userService.UserServiceClient, log *logger.CustomLogger) SessionMidleware {
	return SessionMidleware{
		SessionDelivery: sessionDelivery,
		UserDelivery:    userDelivery,
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
				id, err := u.SessionDelivery.GetIDBySession(r.Context(), &sessionService.SessionID{SessionID: sessionToken})
				if err != nil {
					u.log.LogError(r.Context(), err)
					next.ServeHTTP(w, r)
					return
				}
				user, err := u.UserDelivery.GetUserByID(r.Context(), &userService.UserID{UserID: id.UserID})
				if err != nil {
					u.log.LogError(r.Context(), err)
					next.ServeHTTP(w, r)
					return
				}
				ctx := context.WithValue(r.Context(), packageConfig.RequestUser, user)
				ctx = context.WithValue(ctx, packageConfig.RequestUserID, int(user.UserID))
				ctx = context.WithValue(ctx, packageConfig.RequestUserRule, user.UserRule)
				ctx = context.WithValue(ctx, packageConfig.SessionID, sessionToken)
				r = r.WithContext(ctx)
			}
			next.ServeHTTP(w, r)
		})
	}
}
