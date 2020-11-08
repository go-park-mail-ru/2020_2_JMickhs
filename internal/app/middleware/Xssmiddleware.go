package middlewareApi

import (
	"bytes"
	"fmt"
	"github.com/gorilla/mux"
	"html"
	"io/ioutil"
	"net/http"
)

func NewXssMiddleware() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if (r.Header.Get("Content-type") == "application/json"){
				fmt.Println("FSDFDSFDS")
				bod,_ := ioutil.ReadAll(r.Body)

				newBody := html.EscapeString(string(bod))
				fmt.Println(newBody)
				r.Body =ioutil.NopCloser(bytes.NewBuffer([]byte(newBody)))
			}
			next.ServeHTTP(w, r)
		})
	}
}
