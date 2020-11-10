package userDelivery

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	csrf_mock "github.com/go-park-mail-ru/2020_2_JMickhs/internal/app/csrf/mocks"
	hotelmodel "github.com/go-park-mail-ru/2020_2_JMickhs/internal/app/hotels/models"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/go-park-mail-ru/2020_2_JMickhs/configs"

	"github.com/mitchellh/mapstructure"

	"github.com/gorilla/mux"

	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/pkg/clientError"

	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/pkg/responses"

	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/pkg/logger"

	customerror "github.com/go-park-mail-ru/2020_2_JMickhs/internal/pkg/error"
	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/pkg/serverError"

	"github.com/golang/mock/gomock"

	SessionMocks "github.com/go-park-mail-ru/2020_2_JMickhs/internal/app/sessions/mocks"
	user_mock "github.com/go-park-mail-ru/2020_2_JMickhs/internal/app/user/mocks"
	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/app/user/models"
	"github.com/stretchr/testify/assert"
)

func TestUserHandler_Auth(t *testing.T) {
	testUser := models.User{ID: 3, Username: "kostik", Email: "sdfs@mail.ru", Password: "12345"}

	t.Run("Auth", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		getUser := testUser
		getUser.ID = 3
		mockUCase := user_mock.NewMockUsecase(ctrl)
		mockSCase := SessionMocks.NewMockUsecase(ctrl)
		mockSCase.EXPECT().
			AddToken(3).
			Return("123dscv432", nil)
		mockUCase.EXPECT().
			GetByUserName(testUser.Username).
			Return(getUser, nil)
		mockUCase.EXPECT().
			ComparePassword(testUser.Password, getUser.Password).
			Return(nil)

		body, _ := json.Marshal(testUser)
		req, err := http.NewRequest("POST", "/api/v1/sessions", bytes.NewBuffer(body))
		assert.NoError(t, err)

		rec := httptest.NewRecorder()
		handler := UserHandler{
			UserUseCase:     mockUCase,
			SessionsUseCase: mockSCase,
		}

		handler.Auth(rec, req)
		resp := rec.Result()
		user := models.User{}
		body, err = ioutil.ReadAll(resp.Body)
		err = json.Unmarshal(body, &user)

		assert.Equal(t, testUser, getUser)
		assert.Equal(t, http.StatusOK, rec.Code)
	})

	t.Run("Auth-error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockSCase := SessionMocks.NewMockUsecase(ctrl)
		mockUCaseErr := user_mock.NewMockUsecase(ctrl)
		kek := "fdsfsd"
		body, _ := json.Marshal(&kek)
		req, err := http.NewRequest("POST", "/api/v1/sessions", bytes.NewBuffer(body))
		assert.NoError(t, err)

		rec := httptest.NewRecorder()
		handler := UserHandler{
			UserUseCase:     mockUCaseErr,
			SessionsUseCase: mockSCase,
			log:             logger.NewLogger(os.Stdout),
		}

		handler.Auth(rec, req)
		resp := rec.Result()
		body, err = ioutil.ReadAll(resp.Body)
		response := responses.HttpResponse{}
		err = json.Unmarshal(body, &response)
		assert.Equal(t, clientError.BadRequest, response.Code)
	})

	t.Run("Auth-error1", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockSCase := SessionMocks.NewMockUsecase(ctrl)
		mockUCaseErr := user_mock.NewMockUsecase(ctrl)

		mockUCaseErr.EXPECT().
			GetByUserName(testUser.Username).
			Return(testUser, customerror.NewCustomError(errors.New(""), serverError.ServerInternalError, 1))

		body, _ := json.Marshal(testUser)
		req, err := http.NewRequest("POST", "/api/v1/sessions", bytes.NewBuffer(body))
		assert.NoError(t, err)

		rec := httptest.NewRecorder()
		handler := UserHandler{
			UserUseCase:     mockUCaseErr,
			SessionsUseCase: mockSCase,
			log:             logger.NewLogger(os.Stdout),
		}

		handler.Auth(rec, req)
		resp := rec.Result()
		body, err = ioutil.ReadAll(resp.Body)
		response := responses.HttpResponse{}
		err = json.Unmarshal(body, &response)
		assert.Equal(t, serverError.ServerInternalError, response.Code)
	})

	t.Run("Auth-error2", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockSCase := SessionMocks.NewMockUsecase(ctrl)
		mockUCaseErr := user_mock.NewMockUsecase(ctrl)
		errorUser := testUser
		errorUser.Password = "1234"
		mockUCaseErr.EXPECT().
			GetByUserName(testUser.Username).
			Return(errorUser, nil)

		mockUCaseErr.EXPECT().
			ComparePassword(testUser.Password, errorUser.Password).
			Return(customerror.NewCustomError(errors.New(""), clientError.Unauthorizied, 1))

		body, _ := json.Marshal(testUser)
		req, err := http.NewRequest("POST", "/api/v1/sessions", bytes.NewBuffer(body))
		assert.NoError(t, err)

		rec := httptest.NewRecorder()
		handler := UserHandler{
			UserUseCase:     mockUCaseErr,
			SessionsUseCase: mockSCase,
			log:             logger.NewLogger(os.Stdout),
		}

		handler.Auth(rec, req)
		resp := rec.Result()
		body, err = ioutil.ReadAll(resp.Body)
		response := responses.HttpResponse{}
		err = json.Unmarshal(body, &response)
		assert.Equal(t, clientError.Unauthorizied, response.Code)
	})

	t.Run("Auth-error3", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockSCase := SessionMocks.NewMockUsecase(ctrl)
		mockUCaseErr := user_mock.NewMockUsecase(ctrl)

		errorUser := testUser
		errorUser.Password = "1234"
		mockUCaseErr.EXPECT().
			GetByUserName(testUser.Username).
			Return(errorUser, nil)

		mockUCaseErr.EXPECT().
			ComparePassword(testUser.Password, errorUser.Password).
			Return(nil)

		mockSCase.EXPECT().
			AddToken(3).
			Return("123dscv432", customerror.NewCustomError(errors.New(""), serverError.ServerInternalError, 1))

		body, _ := json.Marshal(testUser)
		req, err := http.NewRequest("POST", "/api/v1/sessions", bytes.NewBuffer(body))
		assert.NoError(t, err)

		rec := httptest.NewRecorder()
		handler := UserHandler{
			UserUseCase:     mockUCaseErr,
			SessionsUseCase: mockSCase,
			log:             logger.NewLogger(os.Stdout),
		}

		handler.Auth(rec, req)
		resp := rec.Result()
		body, err = ioutil.ReadAll(resp.Body)
		response := responses.HttpResponse{}
		err = json.Unmarshal(body, &response)
		assert.Equal(t, serverError.ServerInternalError, response.Code)
	})

}

func TestUserHandler_Registration(t *testing.T) {
	testUser := models.User{Username: "kostik", Email: "sdfs@mail.ru", Password: "12345"}

	t.Run("Registration", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockUCase := user_mock.NewMockUsecase(ctrl)
		mockSCase := SessionMocks.NewMockUsecase(ctrl)

		retUser := testUser
		retUser.ID = 4
		mockUCase.EXPECT().
			Add(testUser).
			Return(retUser, nil)
		mockSCase.EXPECT().
			AddToken(4).
			Return("123dscv432", nil)

		body, _ := json.Marshal(testUser)
		req, err := http.NewRequest("POST", "/api/v1/signup", bytes.NewBuffer(body))
		assert.NoError(t, err)

		rec := httptest.NewRecorder()
		handler := UserHandler{
			UserUseCase:     mockUCase,
			SessionsUseCase: mockSCase,
		}

		handler.Registration(rec, req)
		resp := rec.Result()
		body, err = ioutil.ReadAll(resp.Body)
		response := responses.HttpResponse{}
		err = json.Unmarshal(body, &response)

		assert.Equal(t, retUser.ID, 4)
		assert.Equal(t, http.StatusOK, response.Code)
	})
	t.Run("RegistrationErr", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockUCase := user_mock.NewMockUsecase(ctrl)
		mockSCase := SessionMocks.NewMockUsecase(ctrl)

		kek := "fdsfs"
		body, _ := json.Marshal(kek)
		req, err := http.NewRequest("POST", "/api/v1/signup", bytes.NewBuffer(body))
		assert.NoError(t, err)

		rec := httptest.NewRecorder()
		handler := UserHandler{
			UserUseCase:     mockUCase,
			SessionsUseCase: mockSCase,
			log:             logger.NewLogger(os.Stdout),
		}

		handler.Registration(rec, req)
		resp := rec.Result()
		body, err = ioutil.ReadAll(resp.Body)
		response := responses.HttpResponse{}
		err = json.Unmarshal(body, &response)

		assert.Equal(t, clientError.BadRequest, response.Code)
	})
	t.Run("RegistrationErr1", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockUCase := user_mock.NewMockUsecase(ctrl)
		mockSCase := SessionMocks.NewMockUsecase(ctrl)

		retUser := testUser
		retUser.ID = 4
		mockUCase.EXPECT().
			Add(testUser).
			Return(retUser, customerror.NewCustomError(errors.New(""), clientError.Conflict, 1))

		body, _ := json.Marshal(testUser)
		req, err := http.NewRequest("POST", "/api/v1/signup", bytes.NewBuffer(body))
		assert.NoError(t, err)

		rec := httptest.NewRecorder()
		handler := UserHandler{
			UserUseCase:     mockUCase,
			SessionsUseCase: mockSCase,
			log:             logger.NewLogger(os.Stdout),
		}

		handler.Registration(rec, req)
		resp := rec.Result()
		body, err = ioutil.ReadAll(resp.Body)
		response := responses.HttpResponse{}
		err = json.Unmarshal(body, &response)

		assert.Equal(t, clientError.Conflict, response.Code)
	})
	t.Run("RegistrationErr2", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockUCase := user_mock.NewMockUsecase(ctrl)
		mockSCase := SessionMocks.NewMockUsecase(ctrl)

		retUser := testUser
		retUser.ID = 4
		mockUCase.EXPECT().
			Add(testUser).
			Return(retUser, nil)

		mockSCase.EXPECT().
			AddToken(4).
			Return("123dscv432", customerror.NewCustomError(errors.New(""), clientError.BadRequest, 1))

		body, _ := json.Marshal(testUser)
		req, err := http.NewRequest("POST", "/api/v1/signup", bytes.NewBuffer(body))
		assert.NoError(t, err)

		rec := httptest.NewRecorder()
		handler := UserHandler{
			UserUseCase:     mockUCase,
			SessionsUseCase: mockSCase,
			log:             logger.NewLogger(os.Stdout),
		}

		handler.Registration(rec, req)
		resp := rec.Result()
		body, err = ioutil.ReadAll(resp.Body)
		response := responses.HttpResponse{}
		err = json.Unmarshal(body, &response)

		assert.Equal(t, clientError.BadRequest, response.Code)
	})
}

func TestUserHandler_GetAccInfo(t *testing.T) {
	testUser := models.User{ID: 1, Username: "kostik", Email: "sdfs@mail.ru", Password: "12345", Avatar: "kek/img.jpeg"}

	t.Run("GetAccInfo", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockUCase := user_mock.NewMockUsecase(ctrl)
		mockSCase := SessionMocks.NewMockUsecase(ctrl)
		mockUCase.EXPECT().
			GetUserByID(1).
			Return(testUser, nil)

		req, err := http.NewRequest("GET", "/api/v1/getAccInfo", nil)
		req = mux.SetURLVars(req, map[string]string{
			"id": "1",
		})
		assert.NoError(t, err)

		rec := httptest.NewRecorder()
		handler := UserHandler{
			UserUseCase:     mockUCase,
			SessionsUseCase: mockSCase,
		}

		handler.getAccInfo(rec, req)
		resp := rec.Result()
		response := responses.HttpResponse{}
		user := models.SafeUser{}
		json.NewDecoder(resp.Body).Decode(&response)
		err = mapstructure.Decode(response.Data.(map[string]interface{}), &user)
		assert.Equal(t, user.ID, 1)
		assert.Equal(t, http.StatusOK, response.Code)
	})

	t.Run("GetAccInfoErr", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockUCase := user_mock.NewMockUsecase(ctrl)
		mockSCase := SessionMocks.NewMockUsecase(ctrl)
		mockUCase.EXPECT().
			GetUserByID(1).
			Return(testUser, customerror.NewCustomError(errors.New(""), clientError.Gone, 1))

		req, err := http.NewRequest("GET", "/api/v1/getAccInfo", nil)
		req = mux.SetURLVars(req, map[string]string{
			"id": "1",
		})
		assert.NoError(t, err)

		rec := httptest.NewRecorder()
		handler := UserHandler{
			UserUseCase:     mockUCase,
			SessionsUseCase: mockSCase,
			log:             logger.NewLogger(os.Stdout),
		}

		handler.getAccInfo(rec, req)
		resp := rec.Result()
		response := responses.HttpResponse{}
		json.NewDecoder(resp.Body).Decode(&response)
		assert.Equal(t, clientError.Gone, response.Code)
	})
}

func TestUserHandler_UpdatePassword(t *testing.T) {
	testUser := models.User{ID: 1, Username: "kostik", Email: "sdfs@mail.ru", Password: "12345", Avatar: "kek/img.jpeg"}

	t.Run("UpdatePassword", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockUCase := user_mock.NewMockUsecase(ctrl)
		mockSCase := SessionMocks.NewMockUsecase(ctrl)
		retUser := testUser
		retUser.Password = "12345567"

		mockUCase.EXPECT().
			ComparePassword(testUser.Password, "12345").
			Return(nil)
		mockUCase.EXPECT().
			UpdatePassword(retUser).
			Return(nil)

		testPassword := models.UpdatePassword{OldPassword: "12345", NewPassword: "12345567"}
		body, _ := json.Marshal(testPassword)
		req, err := http.NewRequest("PUT", "/api/v1/users/password", bytes.NewBuffer(body))

		assert.NoError(t, err)

		rec := httptest.NewRecorder()
		req = req.WithContext(context.WithValue(req.Context(), configs.RequestUser, testUser))
		handler := UserHandler{
			UserUseCase:     mockUCase,
			SessionsUseCase: mockSCase,
		}

		handler.updatePassword(rec, req)
		resp := rec.Result()
		response := responses.HttpResponse{}
		json.NewDecoder(resp.Body).Decode(&response)

		assert.Equal(t, http.StatusOK, response.Code)
	})

	t.Run("UpdatePasswordErr", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockUCase := user_mock.NewMockUsecase(ctrl)
		mockSCase := SessionMocks.NewMockUsecase(ctrl)

		kek := "fdsfs"
		body, _ := json.Marshal(kek)
		req, err := http.NewRequest("PUT", "/api/v1/users/password", bytes.NewBuffer(body))
		assert.NoError(t, err)

		rec := httptest.NewRecorder()
		handler := UserHandler{
			UserUseCase:     mockUCase,
			SessionsUseCase: mockSCase,
			log:             logger.NewLogger(os.Stdout),
		}

		handler.updatePassword(rec, req)
		resp := rec.Result()
		body, err = ioutil.ReadAll(resp.Body)
		response := responses.HttpResponse{}
		err = json.Unmarshal(body, &response)

		assert.Equal(t, clientError.BadRequest, response.Code)
	})

	t.Run("UpdatePassword2", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockUCase := user_mock.NewMockUsecase(ctrl)
		mockSCase := SessionMocks.NewMockUsecase(ctrl)
		retUser := testUser
		retUser.Password = "12345567"
		testPassword := models.UpdatePassword{OldPassword: "1234", NewPassword: "12345567"}
		mockUCase.EXPECT().
			ComparePassword(testPassword.OldPassword, testUser.Password).
			Return(customerror.NewCustomError(errors.New(""), clientError.PaymentReq, 1))

		body, _ := json.Marshal(testPassword)
		req, err := http.NewRequest("PUT", "/api/v1/users/password", bytes.NewBuffer(body))

		assert.NoError(t, err)

		rec := httptest.NewRecorder()
		req = req.WithContext(context.WithValue(req.Context(), configs.RequestUser, testUser))
		handler := UserHandler{
			UserUseCase:     mockUCase,
			SessionsUseCase: mockSCase,
			log:             logger.NewLogger(os.Stdout),
		}

		handler.updatePassword(rec, req)
		resp := rec.Result()
		response := responses.HttpResponse{}
		json.NewDecoder(resp.Body).Decode(&response)

		assert.Equal(t, clientError.PaymentReq, response.Code)
	})

	t.Run("UpdatePassword3", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockUCase := user_mock.NewMockUsecase(ctrl)
		mockSCase := SessionMocks.NewMockUsecase(ctrl)
		retUser := testUser
		retUser.Password = "12345567"
		testPassword := models.UpdatePassword{OldPassword: "12345", NewPassword: "12345567"}
		mockUCase.EXPECT().
			ComparePassword(testPassword.OldPassword, testUser.Password).
			Return(nil)
		mockUCase.EXPECT().
			UpdatePassword(retUser).
			Return(customerror.NewCustomError(errors.New(""), serverError.ServerInternalError, 1))

		body, _ := json.Marshal(testPassword)
		req, err := http.NewRequest("PUT", "/api/v1/users/password", bytes.NewBuffer(body))

		assert.NoError(t, err)

		rec := httptest.NewRecorder()
		req = req.WithContext(context.WithValue(req.Context(), configs.RequestUser, testUser))
		handler := UserHandler{
			UserUseCase:     mockUCase,
			SessionsUseCase: mockSCase,
			log:             logger.NewLogger(os.Stdout),
		}

		handler.updatePassword(rec, req)
		resp := rec.Result()
		response := responses.HttpResponse{}
		json.NewDecoder(resp.Body).Decode(&response)

		assert.Equal(t, serverError.ServerInternalError, response.Code)
	})

	t.Run("UpdatePassword4", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockUCase := user_mock.NewMockUsecase(ctrl)
		mockSCase := SessionMocks.NewMockUsecase(ctrl)
		retUser := testUser
		retUser.Password = "12345567"
		testPassword := models.UpdatePassword{OldPassword: "12345", NewPassword: "12345567"}

		body, _ := json.Marshal(testPassword)
		req, err := http.NewRequest("PUT", "/api/v1/users/password", bytes.NewBuffer(body))

		assert.NoError(t, err)

		rec := httptest.NewRecorder()
		handler := UserHandler{
			UserUseCase:     mockUCase,
			SessionsUseCase: mockSCase,
			log:             logger.NewLogger(os.Stdout),
		}

		handler.updatePassword(rec, req)
		resp := rec.Result()
		response := responses.HttpResponse{}
		json.NewDecoder(resp.Body).Decode(&response)

		assert.Equal(t, clientError.Unauthorizied, response.Code)
	})
}

func TestUserHandler_UpdateUser(t *testing.T) {
	testUser := models.User{ID: 1, Username: "kostik", Email: "sdfs@mail.ru", Password: "12345", Avatar: "kek/img.jpeg"}

	t.Run("UpdateUser", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockUCase := user_mock.NewMockUsecase(ctrl)
		mockSCase := SessionMocks.NewMockUsecase(ctrl)

		mockUCase.EXPECT().
			UpdateUser(testUser).
			Return(nil)

		body, _ := json.Marshal(testUser)
		req, err := http.NewRequest("PUT", "/api/v1/users/credentials", bytes.NewBuffer(body))

		assert.NoError(t, err)

		rec := httptest.NewRecorder()
		req = req.WithContext(context.WithValue(req.Context(), configs.RequestUser, testUser))
		handler := UserHandler{
			UserUseCase:     mockUCase,
			SessionsUseCase: mockSCase,
			log:             logger.NewLogger(os.Stdout),
		}

		handler.UpdateUser(rec, req)
		resp := rec.Result()
		response := responses.HttpResponse{}
		json.NewDecoder(resp.Body).Decode(&response)

		assert.Equal(t, http.StatusOK, response.Code)
	})

	t.Run("UpdateUserErr1", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockUCase := user_mock.NewMockUsecase(ctrl)
		mockSCase := SessionMocks.NewMockUsecase(ctrl)

		hotel:=[]hotelmodel.Hotel{{Name: "fds"},{Name: "lel"}}
		body, _ := json.Marshal(hotel)
		req, err := http.NewRequest("PUT", "/api/v1/users/credentials", bytes.NewBuffer(body))

		assert.NoError(t, err)

		rec := httptest.NewRecorder()
		req = req.WithContext(context.WithValue(req.Context(), configs.RequestUser, testUser))
		handler := UserHandler{
			UserUseCase:     mockUCase,
			SessionsUseCase: mockSCase,
			log:             logger.NewLogger(os.Stdout),
		}

		handler.UpdateUser(rec, req)
		resp := rec.Result()
		response := responses.HttpResponse{}
		json.NewDecoder(resp.Body).Decode(&response)

		assert.Equal(t, clientError.BadRequest, response.Code)
	})

	t.Run("UpdateUserErr2", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockUCase := user_mock.NewMockUsecase(ctrl)
		mockSCase := SessionMocks.NewMockUsecase(ctrl)

		body, _ := json.Marshal(testUser)
		req, err := http.NewRequest("PUT", "/api/v1/users/credentials", bytes.NewBuffer(body))

		assert.NoError(t, err)

		rec := httptest.NewRecorder()
		handler := UserHandler{
			UserUseCase:     mockUCase,
			SessionsUseCase: mockSCase,
			log:             logger.NewLogger(os.Stdout),
		}

		handler.UpdateUser(rec, req)
		resp := rec.Result()
		response := responses.HttpResponse{}
		json.NewDecoder(resp.Body).Decode(&response)

		assert.Equal(t, clientError.Unauthorizied, response.Code)
	})

	t.Run("UpdateUserErr2", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockUCase := user_mock.NewMockUsecase(ctrl)
		mockSCase := SessionMocks.NewMockUsecase(ctrl)

		mockUCase.EXPECT().
			UpdateUser(testUser).
			Return(customerror.NewCustomError(errors.New(""),clientError.Conflict,1))

		body, _ := json.Marshal(testUser)
		req, err := http.NewRequest("PUT", "/api/v1/users/credentials", bytes.NewBuffer(body))

		assert.NoError(t, err)
		req = req.WithContext(context.WithValue(req.Context(), configs.RequestUser, testUser))
		rec := httptest.NewRecorder()
		handler := UserHandler{
			UserUseCase:     mockUCase,
			SessionsUseCase: mockSCase,
			log:             logger.NewLogger(os.Stdout),
		}

		handler.UpdateUser(rec, req)
		resp := rec.Result()
		response := responses.HttpResponse{}
		json.NewDecoder(resp.Body).Decode(&response)

		assert.Equal(t, clientError.Conflict, response.Code)
	})
}

func TestUserHandler_GetCsrf(t *testing.T) {
	t.Run("UpdateUser", func(t *testing.T) {
		sId := "fsdfsd"
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockUCase := user_mock.NewMockUsecase(ctrl)
		mockSCase := SessionMocks.NewMockUsecase(ctrl)
		mockCsrfCase := csrf_mock.NewMockUsecase(ctrl)
		mockCsrfCase.EXPECT().
			CreateToken(sId, time.Now().Unix()).
			Return("21fds",nil)

		req, err := http.NewRequest("GET", "/api/v1/csrf", nil)

		assert.NoError(t, err)
		req = req.WithContext(context.WithValue(req.Context(), configs.SessionID,sId ))
		rec := httptest.NewRecorder()
		handler := UserHandler{
			UserUseCase:     mockUCase,
			SessionsUseCase: mockSCase,
			csrfUseCase: mockCsrfCase,
			log:             logger.NewLogger(os.Stdout),
		}

		handler.GetCsrf(rec, req)
		resp := rec.Result()
		response := responses.HttpResponse{}
		json.NewDecoder(resp.Body).Decode(&response)

		assert.Equal(t, http.StatusOK, response.Code)
	})

	t.Run("UpdateUserErr1", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockUCase := user_mock.NewMockUsecase(ctrl)
		mockSCase := SessionMocks.NewMockUsecase(ctrl)

		req, err := http.NewRequest("GET", "/api/v1/csrf", nil)

		assert.NoError(t, err)

		rec := httptest.NewRecorder()
		handler := UserHandler{
			UserUseCase:     mockUCase,
			SessionsUseCase: mockSCase,
			log:             logger.NewLogger(os.Stdout),
		}

		handler.GetCsrf(rec, req)
		resp := rec.Result()
		response := responses.HttpResponse{}
		json.NewDecoder(resp.Body).Decode(&response)

		assert.Equal(t, clientError.Unauthorizied, response.Code)
	})

	t.Run("UpdateUserErr2", func(t *testing.T) {
		sId := "fsdfsd"
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockUCase := user_mock.NewMockUsecase(ctrl)
		mockSCase := SessionMocks.NewMockUsecase(ctrl)
		mockCsrfCase := csrf_mock.NewMockUsecase(ctrl)
		mockCsrfCase.EXPECT().
			CreateToken(sId, time.Now().Unix()).
			Return("21fds",customerror.NewCustomError(errors.New("fds"),serverError.ServerInternalError,1))

		req, err := http.NewRequest("GET", "/api/v1/csrf", nil)

		assert.NoError(t, err)
		req = req.WithContext(context.WithValue(req.Context(), configs.SessionID,sId ))
		rec := httptest.NewRecorder()
		handler := UserHandler{
			UserUseCase:     mockUCase,
			SessionsUseCase: mockSCase,
			csrfUseCase: mockCsrfCase,
			log:             logger.NewLogger(os.Stdout),
		}

		handler.GetCsrf(rec, req)
		resp := rec.Result()
		response := responses.HttpResponse{}
		json.NewDecoder(resp.Body).Decode(&response)

		assert.Equal(t, serverError.ServerInternalError, response.Code)
	})

}

func TestUserHandler_SignOut(t *testing.T) {
	t.Run("SignOut", func(t *testing.T) {
		token := "dfsdxzc"
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockUCase := user_mock.NewMockUsecase(ctrl)
		mockSCase := SessionMocks.NewMockUsecase(ctrl)
		mockCsrfCase := csrf_mock.NewMockUsecase(ctrl)
		mockSCase.EXPECT().
			DeleteSession(token).
			Return(nil)

		req, err := http.NewRequest("DELETE", "/api/v1/sessions", nil)

		assert.NoError(t, err)
		cookie := &http.Cookie{
			Name:     "session_token",
			Value:    token,
			Path:     "/",
			HttpOnly: true,
			Expires:  time.Now().Add(configs.CookieLifeTime),
		}
		req.AddCookie(cookie)
		rec := httptest.NewRecorder()
		handler := UserHandler{
			UserUseCase:     mockUCase,
			SessionsUseCase: mockSCase,
			csrfUseCase: mockCsrfCase,
			log:             logger.NewLogger(os.Stdout),
		}

		handler.SignOut(rec, req)
		resp := rec.Result()
		response := responses.HttpResponse{}
		json.NewDecoder(resp.Body).Decode(&response)

		assert.Equal(t, http.StatusOK, response.Code)
	})

	t.Run("SignOutErr1", func(t *testing.T) {

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockUCase := user_mock.NewMockUsecase(ctrl)
		mockSCase := SessionMocks.NewMockUsecase(ctrl)
		mockCsrfCase := csrf_mock.NewMockUsecase(ctrl)


		req, err := http.NewRequest("DELETE", "/api/v1/sessions", nil)

		assert.NoError(t, err)
		rec := httptest.NewRecorder()
		handler := UserHandler{
			UserUseCase:     mockUCase,
			SessionsUseCase: mockSCase,
			csrfUseCase: mockCsrfCase,
			log:             logger.NewLogger(os.Stdout),
		}

		handler.SignOut(rec, req)
		resp := rec.Result()
		response := responses.HttpResponse{}
		json.NewDecoder(resp.Body).Decode(&response)

		assert.Equal(t, clientError.BadRequest, response.Code)
	})

	t.Run("SignOutErr2", func(t *testing.T) {
		token := "dfsdxzc"
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockUCase := user_mock.NewMockUsecase(ctrl)
		mockSCase := SessionMocks.NewMockUsecase(ctrl)
		mockCsrfCase := csrf_mock.NewMockUsecase(ctrl)
		mockSCase.EXPECT().
			DeleteSession(token).
			Return(customerror.NewCustomError(errors.New(""),serverError.ServerInternalError,1))

		req, err := http.NewRequest("DELETE", "/api/v1/sessions", nil)

		assert.NoError(t, err)
		cookie := &http.Cookie{
			Name:     "session_token",
			Value:    token,
			Path:     "/",
			HttpOnly: true,
			Expires:  time.Now().Add(configs.CookieLifeTime),
		}
		req.AddCookie(cookie)
		rec := httptest.NewRecorder()
		handler := UserHandler{
			UserUseCase:     mockUCase,
			SessionsUseCase: mockSCase,
			csrfUseCase: mockCsrfCase,
			log:             logger.NewLogger(os.Stdout),
		}

		handler.SignOut(rec, req)
		resp := rec.Result()
		response := responses.HttpResponse{}
		json.NewDecoder(resp.Body).Decode(&response)

		assert.Equal(t, serverError.ServerInternalError, response.Code)
	})
}

