package middlewareApi

import (
	"context"
	"net/http"

	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/pkg/clientError"

	"github.com/go-park-mail-ru/2020_2_JMickhs/configs"

	customerror "github.com/go-park-mail-ru/2020_2_JMickhs/internal/pkg/error"

	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/pkg/logger"

	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/app/sessions"
	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/app/user"
	"github.com/gorilla/mux"
)

type SessionMidleware struct {
	SessUseCase sessions.Usecase
	UserUseCase user.Usecase
	log         *logger.CustomLogger
}

func NewSessionMiddleware(su sessions.Usecase, uuc user.Usecase, log *logger.CustomLogger) SessionMidleware {
	return SessionMidleware{
		SessUseCase: su,
		UserUseCase: uuc,
		log:         log,
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
				id, err := u.SessUseCase.GetIDByToken(sessionToken)
				if err != nil {
					u.log.LogError(r.Context(), err)
					next.ServeHTTP(w, r)
					return
				}
				user, err := u.UserUseCase.GetUserByID(id)
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
