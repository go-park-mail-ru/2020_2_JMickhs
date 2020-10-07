package userDelivery

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/bxcodec/faker/v3"
	"github.com/go-park-mail-ru/2020_2_JMickhs/configs"
	SessionMocks "github.com/go-park-mail-ru/2020_2_JMickhs/internal/sessions/mocks"
	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/user/mocks"
	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/user/models"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"image"
	"image/png"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"net/textproto"
	"strings"
	"testing"
	"time"
)

func TestUserHandler_Auth(t *testing.T) {
	var mockUser models.User
	err := faker.FakeData(&mockUser)

	assert.NoError(t,err)
	mockSCase := new(SessionMocks.SessionsUsecase)
	mockSCase.On("AddToken", mock.AnythingOfType("int")).Return("1", nil)
	t.Run("Auth",func(t *testing.T) {
		mockUCase := new(mocks.UserUsecase)


		mockUCase.On("GetByUserName", mock.AnythingOfType("string")).Return(mockUser, nil)
		mockUCase.On("ComparePassword", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(nil)


		body, _ := json.Marshal(mockUser)
		req, err := http.NewRequest("POST", "/api/v1/signin", bytes.NewBuffer(body))
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

		assert.Equal(t, user.ID, mockUser.ID)

		assert.Equal(t, http.StatusOK, rec.Code)
		mockUCase.AssertExpectations(t)
	})

	t.Run("Auth-error",func(t *testing.T) {
		mockUCaseErr := new(mocks.UserUsecase)

		mockUCaseErr.On("GetByUserName",mock.AnythingOfType("string")).Return(mockUser,nil)
		mockUCaseErr.On("ComparePassword",mock.AnythingOfType("string"),mock.AnythingOfType("string")).Return(errors.New("dsf"))
		body , _ := json.Marshal(mockUser)
		req,err := http.NewRequest("POST","/api/v1/signin",bytes.NewBuffer(body))
		assert.NoError(t,err)

		rec := httptest.NewRecorder()
		handler := UserHandler{
			UserUseCase: mockUCaseErr,
			SessionsUseCase: mockSCase,
			log: logrus.New(),
		}

		handler.Auth(rec,req)
		assert.Equal(t,http.StatusUnauthorized,rec.Code)
		mockUCaseErr.AssertExpectations(t)
	})

}

func TestUserHandler_Registration(t *testing.T) {
	var mockUser models.User
	err := faker.FakeData(&mockUser)

	assert.NoError(t,err)

	mockUCase := new(mocks.UserUsecase)
	mockSCase := new(SessionMocks.SessionsUsecase)

	mockUCase.On("Add",mock.AnythingOfType("models.User")).Return(mockUser,nil)
	mockUCase.On("SetDefaultAvatar", mock.AnythingOfType("*models.User")).Return(nil)
	mockSCase.On("AddToken",mock.AnythingOfType("int")).Return("1",nil)

	body , _ := json.Marshal(mockUser)
	req,err := http.NewRequest("POST","/api/v1/signup",bytes.NewBuffer(body))
	assert.NoError(t,err)

	rec := httptest.NewRecorder()
	handler := UserHandler{
		UserUseCase: mockUCase,
		SessionsUseCase: mockSCase,
	}

	handler.Registration(rec,req)
	resp := rec.Result()
	user := models.User{}
	body, err = ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body,&user)

	assert.Equal(t,user.ID,mockUser.ID)

	assert.Equal(t,http.StatusOK,rec.Code)
	mockUCase.AssertExpectations(t)
}

func TestUserHandler_GetAccInfo(t *testing.T) {
	var mockUser models.User
	err := faker.FakeData(&mockUser)

	assert.NoError(t,err)

	mockUCase := new(mocks.UserUsecase)
	mockSCase := new(SessionMocks.SessionsUsecase)

	mockUCase.On("GetByUserName",mock.AnythingOfType("string")).Return(mockUser,nil)

	body , _ := json.Marshal(mockUser)
	req,err := http.NewRequest("GET","/api/v1/getAccInfo",bytes.NewBuffer(body))
	assert.NoError(t,err)

	rec := httptest.NewRecorder()
	handler := UserHandler{
		UserUseCase: mockUCase,
		SessionsUseCase: mockSCase,
	}

	handler.getAccInfo(rec,req)
	resp := rec.Result()
	user := models.User{}
	body, err = ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body,&user)

	assert.Equal(t,user.ID,mockUser.ID)

	assert.Equal(t,http.StatusOK,rec.Code)
	mockUCase.AssertExpectations(t)
}

func TestUserHandler_UpdateUser(t *testing.T) {
	var mockUser models.User
	err := faker.FakeData(&mockUser)

	assert.NoError(t,err)

	mockUCase := new(mocks.UserUsecase)
	mockSCase := new(SessionMocks.SessionsUsecase)

	mockUCase.On("UpdateUser",mock.AnythingOfType("models.User")).Return(nil)
	mockSCase.On("GetIDByToken",mock.AnythingOfType("string")).Return(mockUser.ID,err)

	body , _ := json.Marshal(mockUser)
	req,err := http.NewRequest("PUT","/api/v1/updateUser",bytes.NewBuffer(body))

	assert.NoError(t,err)

	rec := httptest.NewRecorder()
	req = req.WithContext(context.WithValue(req.Context(),"User",mockUser))
	handler := UserHandler{
		UserUseCase: mockUCase,
		SessionsUseCase: mockSCase,
	}

	handler.UpdateUser(rec,req)

	assert.Equal(t,http.StatusOK,rec.Code)
	mockUCase.AssertExpectations(t)
}

func TestUserHandler_UpdatePassword(t *testing.T) {
	var mockUser models.User
	err := faker.FakeData(&mockUser)

	assert.NoError(t,err)
	mockSCase := new(SessionMocks.SessionsUsecase)
	mockSCase.On("GetIDByToken", mock.AnythingOfType("string")).Return(mockUser.ID, err)
	body, _ := json.Marshal(mockUser)
	t.Run("UpdatePassword",func(t *testing.T) {

		mockUCase := new(mocks.UserUsecase)

		mockUCase.On("UpdatePassword", mock.AnythingOfType("models.User")).Return(nil)
		mockUCase.On("ComparePassword", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(nil)

		req, err := http.NewRequest("PUT", "/api/v1/updatePassword", bytes.NewBuffer(body))

		assert.NoError(t, err)

		rec := httptest.NewRecorder()
		req = req.WithContext(context.WithValue(req.Context(), "User", mockUser))
		handler := UserHandler{
			UserUseCase:     mockUCase,
			SessionsUseCase: mockSCase,
		}

		handler.updatePassword(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)
		mockUCase.AssertExpectations(t)
	})

	t.Run("UpdatePassword-error",func(t *testing.T) {
		mockUCaseErr := new(mocks.UserUsecase)
		req,err := http.NewRequest("PUT","/api/v1/updatePassword",bytes.NewBuffer(body))
		assert.NoError(t,err)
		rec := httptest.NewRecorder()
		mockUCaseErr.On("UpdatePassword",mock.AnythingOfType("models.User")).Return(nil)
		mockUCaseErr.On("ComparePassword",mock.AnythingOfType("string"),mock.AnythingOfType("string")).Return(errors.New("resd"))

		req = req.WithContext(context.WithValue(req.Context(),"User",mockUser))
		handler := UserHandler{
			UserUseCase: mockUCaseErr,
			SessionsUseCase: mockSCase,
			log: logrus.New(),
		}

		handler.updatePassword(rec,req)

		assert.Equal(t,409,rec.Code)

	})
}

func TestUserHandler_UpdateAvatar(t *testing.T) {
	var mockUser models.User
	err := faker.FakeData(&mockUser)

	assert.NoError(t,err)
	buf, wr := createMultipartFormFile(t)
	mockUCase := new(mocks.UserUsecase)
	mockSCase := new(SessionMocks.SessionsUsecase)

	mockUCase.On("UpdateAvatar", mock.AnythingOfType("models.User")).Return(nil)
	mockUCase.On("UploadAvatar", mock.AnythingOfType("multipart.sectionReadCloser"),
		mock.AnythingOfType("string"), mock.AnythingOfType("*models.User")).Return(nil)
	mockSCase.On("GetIDByToken", mock.AnythingOfType("string")).Return(mockUser.ID, err)

	t.Run("UpdateAvatar",func(t *testing.T) {

		req, err := http.NewRequest("PUT", "/api/v1/updateAvatar", &buf)

		assert.NoError(t, err)
		req.Header.Set("Content-Type", wr.FormDataContentType())
		rec := httptest.NewRecorder()
		req = req.WithContext(context.WithValue(req.Context(), "User", mockUser))
		handler := UserHandler{
			UserUseCase:     mockUCase,
			SessionsUseCase: mockSCase,
		}

		handler.UpdateAvatar(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)
		mockUCase.AssertExpectations(t)
	})

	t.Run("UpdateAvatar-error",func(t *testing.T){
		req,err := http.NewRequest("PUT","/api/v1/updateAvatar",&buf)

		assert.NoError(t,err)
		rec := httptest.NewRecorder()
		req = req.WithContext(context.WithValue(req.Context(),"User",mockUser))
		handler := UserHandler{
			UserUseCase: mockUCase,
			SessionsUseCase: mockSCase,
			log: logrus.New(),
		}

		handler.UpdateAvatar(rec,req)

		assert.Equal(t,400,rec.Code)
	})
}

func createMultipartFormFile(t *testing.T) (bytes.Buffer, *multipart.Writer) {
	var b bytes.Buffer
	var err error
	w := multipart.NewWriter(&b)
	header := textproto.MIMEHeader{}
	header.Set("Content-Disposition",fmt.Sprintf(`form-data; name="%s"; filename="%s"`,
		"avatar","kek.png"))
	header.Set("Content-Type","application/png")
	fw ,err := w.CreatePart(header)
	if err != nil {
		t.Errorf("Error creating writer: %v", err)
	}

	img := image.NewRGBA(image.Rectangle{image.Point{0, 0},image.Point{0, 0}})
	err = png.Encode(fw,img)


	err = w.Close()
	if err != nil {
		t.Error(err)
	}

	return b, w
}

func TestUserHandler_SignOut(t *testing.T) {
	var mockUser models.User
	err := faker.FakeData(&mockUser)

	assert.NoError(t,err)

	mockUCase := new(mocks.UserUsecase)
	mockSCase := new(SessionMocks.SessionsUsecase)

	mockSCase.On("DeleteSession",mock.AnythingOfType("string")).Return(err)

	body , _ := json.Marshal(mockUser)
	req,err := http.NewRequest("POST","/api/v1/signOut",bytes.NewBuffer(body))

	assert.NoError(t,err)

	rec := httptest.NewRecorder()
	req = req.WithContext(context.WithValue(req.Context(),"User",mockUser))
	req.AddCookie(&http.Cookie{
		Name:     "session_token",
		Value:    "dsf",
		Path:     "/",
		HttpOnly: true,
		Expires:  time.Now().Add( configs.CookieLifeTime),
	})
	handler := UserHandler{
		UserUseCase: mockUCase,
		SessionsUseCase: mockSCase,
	}

	handler.SignOut(rec,req)

	assert.Equal(t,http.StatusOK,rec.Code)
	mockUCase.AssertExpectations(t)
}

func TestUserHandler_GetCurrentUser(t *testing.T) {
	var mockUser models.User
	err := faker.FakeData(&mockUser)

	assert.NoError(t,err)

	mockUCase := new(mocks.UserUsecase)
	mockSCase := new(SessionMocks.SessionsUsecase)

	mockUCase.On("CheckEmpty",mock.AnythingOfType("models.User")).Return(false)

	req,err := http.NewRequest("GET","/api/v1/get_current_user",strings.NewReader(""))

	assert.NoError(t,err)

	rec := httptest.NewRecorder()
	req = req.WithContext(context.WithValue(req.Context(),"User",mockUser))
	handler := UserHandler{
		UserUseCase: mockUCase,
		SessionsUseCase: mockSCase,
	}


	handler.GetCurrentUser(rec,req)
	resp := rec.Result()
	body, err := ioutil.ReadAll(resp.Body)
	var user models.User
	err = json.Unmarshal(body,&user)

	assert.Equal(t,user.ID,mockUser.ID)
	assert.Equal(t,http.StatusOK,rec.Code)
	mockUCase.AssertExpectations(t)
}

