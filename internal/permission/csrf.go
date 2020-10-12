package permissions

import (
	"errors"
	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/responses"
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"
)

func generateCsrfLogic(w http.ResponseWriter) {
	csrf := uuid.NewV4()

	timeDelta := time.Now().Add(time.Hour * 2)
	cookie1 := &http.Cookie{Name: "csrf", Value: csrf.String(), Path: "/", HttpOnly: true, Expires: timeDelta}

	http.SetCookie(w, cookie1)
	w.Header().Set("csrf", csrf.String())

}

func SetCSRF(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			generateCsrfLogic(w)
			next.ServeHTTP(w, r)
		})
}

func CheckCSRF(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			csrf := r.Header.Get("X-Csrf-Token")
			csrfCookie, err := r.Cookie("csrf")

			if err != nil || csrf == "" || csrfCookie.Value == "" || csrfCookie.Value != csrf {
				responses.SendErrorResponse(w,419, errors.New("csrf unvalid"))
				return
			}
			generateCsrfLogic(w)
			next.ServeHTTP(w, r)
		})

}
