package permissions

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSetCSRF(t *testing.T) {
	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		csrf := w.Header().Get("csrf")

		assert.NotEqual(t, "", csrf)
	})

	handlerToTest := SetCSRF(nextHandler)

	req := httptest.NewRequest("GET", "http://testing", nil)

	recorder := httptest.NewRecorder()
	handlerToTest.ServeHTTP(recorder, req)
}

func TestCheckCsrf(t *testing.T) {
	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		csrf := w.Header().Get("csrf")
		assert.Equal(t, "", csrf)
	})
	handlerToTest := CheckCSRF(nextHandler)

	req := httptest.NewRequest("GET", "http://testing", nil)

	recorder := httptest.NewRecorder()
	handlerToTest.ServeHTTP(recorder, req)
}
