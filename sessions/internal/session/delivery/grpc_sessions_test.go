package delivery

import (
	"context"
	"errors"
	"testing"
	"time"

	sessionService "github.com/go-park-mail-ru/2020_2_JMickhs/package/proto/sessions"

	csrf_mock "github.com/go-park-mail-ru/2020_2_JMickhs/sessions/internal/csrf/mocks"
	sessions_mock "github.com/go-park-mail-ru/2020_2_JMickhs/sessions/internal/session/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestSessionDelivery_CreateSession(t *testing.T) {
	userID := 3
	sessionID := "fsdfsd"
	testError := errors.New("zxc")
	t.Run("CreateSession", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sessionMock := sessions_mock.NewMockUsecase(ctrl)
		csrfMock := csrf_mock.NewMockUsecase(ctrl)
		in := &sessionService.UserID{UserID: 3}

		delivery := NewSessionDelivery(sessionMock, csrfMock)

		sessionMock.EXPECT().
			AddToken(int64(userID)).
			Return("fsdfsd", nil)

		userResponse, err := delivery.CreateSession(context.Background(), in)
		assert.NoError(t, err)
		assert.Equal(t, userResponse.SessionID, sessionID)
	})
	t.Run("CreateSessionErr", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sessionMock := sessions_mock.NewMockUsecase(ctrl)
		csrfMock := csrf_mock.NewMockUsecase(ctrl)
		in := &sessionService.UserID{UserID: 3}

		delivery := NewSessionDelivery(sessionMock, csrfMock)

		sessionMock.EXPECT().
			AddToken(int64(userID)).
			Return("fsdfsd", testError)

		_, err := delivery.CreateSession(context.Background(), in)
		assert.Error(t, err)
	})
}

func TestSessionDelivery_GetIDBySession(t *testing.T) {
	userID := 3
	sessionID := "fsdfsd"
	testError := errors.New("zxc")
	t.Run("GetIDBySession", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sessionMock := sessions_mock.NewMockUsecase(ctrl)
		csrfMock := csrf_mock.NewMockUsecase(ctrl)
		in := &sessionService.SessionID{SessionID: sessionID}

		delivery := NewSessionDelivery(sessionMock, csrfMock)

		sessionMock.EXPECT().
			GetIDByToken(sessionID).
			Return(int64(userID), nil)

		userResponse, err := delivery.GetIDBySession(context.Background(), in)
		assert.NoError(t, err)
		assert.Equal(t, userResponse.UserID, int64(userID))
	})
	t.Run("GetIDBySessionErr", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sessionMock := sessions_mock.NewMockUsecase(ctrl)
		csrfMock := csrf_mock.NewMockUsecase(ctrl)
		in := &sessionService.SessionID{SessionID: sessionID}
		delivery := NewSessionDelivery(sessionMock, csrfMock)

		sessionMock.EXPECT().
			GetIDByToken(sessionID).
			Return(int64(userID), testError)

		_, err := delivery.GetIDBySession(context.Background(), in)
		assert.Error(t, err)
	})

}

func TestSessionDelivery_DeleteSession(t *testing.T) {
	sessionID := "fsdfsd"
	testError := errors.New("zxc")
	t.Run("DeleteSessions", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sessionMock := sessions_mock.NewMockUsecase(ctrl)
		csrfMock := csrf_mock.NewMockUsecase(ctrl)
		in := &sessionService.SessionID{SessionID: sessionID}

		delivery := NewSessionDelivery(sessionMock, csrfMock)

		sessionMock.EXPECT().
			DeleteSession(sessionID).
			Return(nil)

		_, err := delivery.DeleteSession(context.Background(), in)
		assert.NoError(t, err)
	})
	t.Run("DeleteSessionsErr", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sessionMock := sessions_mock.NewMockUsecase(ctrl)
		csrfMock := csrf_mock.NewMockUsecase(ctrl)
		in := &sessionService.SessionID{SessionID: sessionID}
		delivery := NewSessionDelivery(sessionMock, csrfMock)

		sessionMock.EXPECT().
			DeleteSession(sessionID).
			Return(testError)

		_, err := delivery.DeleteSession(context.Background(), in)
		assert.Error(t, err)
	})

}

func TestSessionDelivery_CheckCsrfToken(t *testing.T) {
	sessionID := "fsdfsd"
	testError := errors.New("zxc")
	t.Run("CheckCsrfToken", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sessionMock := sessions_mock.NewMockUsecase(ctrl)
		csrfMock := csrf_mock.NewMockUsecase(ctrl)
		in := &sessionService.CsrfTokenInput{SessionID: sessionID, TimeStamp: time.Now().Unix()}

		delivery := NewSessionDelivery(sessionMock, csrfMock)

		csrfMock.EXPECT().
			CreateToken(sessionID, time.Now().Unix()).
			Return("fsdfsd", nil)

		_, err := delivery.CreateCsrfToken(context.Background(), in)
		assert.NoError(t, err)
	})
	t.Run("CheckCsrfTokenErr", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sessionMock := sessions_mock.NewMockUsecase(ctrl)
		csrfMock := csrf_mock.NewMockUsecase(ctrl)
		in := &sessionService.CsrfTokenInput{SessionID: sessionID, TimeStamp: time.Now().Unix()}

		delivery := NewSessionDelivery(sessionMock, csrfMock)

		csrfMock.EXPECT().
			CreateToken(sessionID, time.Now().Unix()).
			Return("fsdfsd", testError)

		_, err := delivery.CreateCsrfToken(context.Background(), in)
		assert.Error(t, err)
	})

}

func TestSessionDelivery_CreateCsrfToken(t *testing.T) {
	token := "token"
	sessionID := "fsdfsd"
	testError := errors.New("zxc")
	t.Run("CreateCsrfToken", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sessionMock := sessions_mock.NewMockUsecase(ctrl)
		csrfMock := csrf_mock.NewMockUsecase(ctrl)
		in := &sessionService.CsrfTokenCheck{SessionID: sessionID, Token: token}

		delivery := NewSessionDelivery(sessionMock, csrfMock)

		csrfMock.EXPECT().
			CheckToken(sessionID, token).
			Return(true, nil)

		_, err := delivery.CheckCsrfToken(context.Background(), in)
		assert.NoError(t, err)
	})
	t.Run("CreateCsrfTokenErr", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		sessionID = "fsdfsd"
		token = "token"
		sessionMock := sessions_mock.NewMockUsecase(ctrl)
		csrfMock := csrf_mock.NewMockUsecase(ctrl)
		in := &sessionService.CsrfTokenCheck{Token: "token", SessionID: sessionID}

		delivery := NewSessionDelivery(sessionMock, csrfMock)

		csrfMock.EXPECT().
			CheckToken(sessionID, token).
			Return(true, testError)

		_, err := delivery.CheckCsrfToken(context.Background(), in)
		assert.Error(t, err)
	})

}
