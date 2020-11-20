package middlewareApi

import (
	"net/http"

	packageConfig "github.com/go-park-mail-ru/2020_2_JMickhs/package/configs"

	"github.com/gorilla/mux"
)

func NewOptionsHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		allowedOrigin := ""
		if packageConfig.AllowedOrigins[origin] {
			allowedOrigin = origin
		}
		w.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers",
			"Accept, Content-Type, X-CSRF-Token, csrf-token, Content-Length, Accept-Encoding, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Vary", "Accept, Cookie")
		w.WriteHeader(http.StatusNoContent)
		return
	})
}

func MyCORSMethodMiddleware() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			origin := req.Header.Get("Origin")
			allowedOrigin := ""
			if packageConfig.AllowedOrigins[origin] {
				allowedOrigin = origin
			}
			w.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length,"+
				" Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers,"+
				" Access-Control-Request-Method, Connection, Host, Origin, Cache-Control, X-header")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Expose-Headers", "Csrf")
			w.Header().Set("Vary", "Accept, Cookie")

			next.ServeHTTP(w, req)
		})
	}
}
