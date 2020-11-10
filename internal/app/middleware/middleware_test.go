package middlewareApi

import (
	"context"
	"encoding/json"
	"github.com/go-park-mail-ru/2020_2_JMickhs/configs"
	csrf_mock "github.com/go-park-mail-ru/2020_2_JMickhs/internal/app/csrf/mocks"
	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/pkg/clientError"
	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/pkg/logger"
	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/pkg/responses"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestCsrfCheck(t *testing.T) {
	t.Run("CsrfCheck", func(t *testing.T) {

		handler := func(w http.ResponseWriter, r *http.Request) {}
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockCsrfCase := csrf_mock.NewMockUsecase(ctrl)
		mockCsrfCase.EXPECT().
			CheckToken("fdsfsd", "token").
			Return(true, nil)
		middleware := NewCsrfMiddleware(mockCsrfCase, logger.NewLogger(os.Stdout))

		req, err := http.NewRequest("POST", "/api/v1/comments", nil)
		assert.NoError(t, err)
		req = req.WithContext(context.WithValue(req.Context(), configs.CorrectToken, false))
		req = req.WithContext(context.WithValue(req.Context(), configs.SessionID, "fdsfsd"))
		req.Header.Set("X-Csrf-Token", "token")

		res := httptest.NewRecorder()
		handler(res, req)
		middle := middleware.CSRFCheck().Middleware(http.HandlerFunc(handler))
		middle.ServeHTTP(res, req)

	})
	t.Run("CsrfCheckErr", func(t *testing.T) {

		handler := func(w http.ResponseWriter, r *http.Request) {}
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockCsrfCase := csrf_mock.NewMockUsecase(ctrl)
		middleware := NewCsrfMiddleware(mockCsrfCase, logger.NewLogger(os.Stdout))
		mockCsrfCase.EXPECT().
			CheckToken("fdsfsd", "token").
			Return(true, nil)

		req, err := http.NewRequest("POST", "/api/v1/comments", nil)
		assert.NoError(t, err)

		req = req.WithContext(context.WithValue(req.Context(), configs.SessionID, "fdsfsd"))
		req = req.WithContext(context.WithValue(req.Context(), configs.CorrectToken, false))
		req.Header.Set("X-Csrf-Token", "token")

		res := httptest.NewRecorder()
		handler(res, req)
		middle := middleware.CSRFCheck().Middleware(http.HandlerFunc(handler))
		middle.ServeHTTP(res, req)

		corrTok, ok := req.Context().Value(configs.CorrectToken).(bool)
		if !ok {
			t.Error("correctToken value doesn't exist")
		}
		assert.Equal(t, corrTok, false)

	})

	t.Run("CsrfCheckErr", func(t *testing.T) {

		handler := func(w http.ResponseWriter, r *http.Request) {}

		req, err := http.NewRequest("POST", "/api/v1/comments", nil)
		assert.NoError(t, err)

		req = req.WithContext(context.WithValue(req.Context(), configs.CorrectToken, false))
		req.Header.Set("X-Csrf-Token", "token")

		res := httptest.NewRecorder()
		handler(res, req)
		middle := CheckCSRFOnHandler(http.HandlerFunc(handler))
		middle.ServeHTTP(res, req)

		resp := res.Result()
		response := responses.HttpResponse{}
		json.NewDecoder(resp.Body).Decode(&response)

		assert.Equal(t, clientError.Forbidden, response.Code)

	})

}
