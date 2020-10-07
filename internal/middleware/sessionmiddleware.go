package middlewareApi

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/sessions"
	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/user"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type SessionMidleware struct {
	SessUseCase sessions.Usecase
	UserUseCase user.Usecase
	log         *logrus.Logger
}

func NewSessionMiddleware(su sessions.Usecase, uuc user.Usecase, log *logrus.Logger) SessionMidleware {
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
				next.ServeHTTP(w, r)
				return
			}
			if c != nil {
				sessionToken := c.Value
				fmt.Println(sessionToken)
				id, err := u.SessUseCase.GetIDByToken(sessionToken)
				fmt.Println(id)
				user, err := u.UserUseCase.GetUserByID(id)
				if err != nil {
					u.log.Info(err.Error())
				}
				ctx := context.WithValue(r.Context(), "User", user)
				r = r.WithContext(ctx)
			}
			next.ServeHTTP(w, r)
		})
	}
}
