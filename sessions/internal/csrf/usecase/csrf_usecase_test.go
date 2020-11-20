package csrfUsecase

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"testing"
	"time"

	"github.com/go-park-mail-ru/2020_2_JMickhs/configs"
	csrf_mock "github.com/go-park-mail-ru/2020_2_JMickhs/main/internal/app/csrf/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCSRF(t *testing.T) {
	size := 24

	rb := make([]byte, size)
	_, err := rand.Read(rb)
	assert.NoError(t, err)
	sID := "ddgfdxc12sr"
	testError := errors.New("err")
	configs.SecretTokenKey = base64.URLEncoding.EncodeToString(rb)
	t.Run("TestToken1", func(t *testing.T) {

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		csrfRepo := csrf_mock.NewMockRepository(ctrl)

		cryptHashToken := NewCsrfUsecase(csrfRepo)
		assert.NoError(t, err)

		token, err := cryptHashToken.CreateToken(sID, time.Now().Unix())
		assert.NoError(t, err)

		csrfRepo.EXPECT().
			Check(token).
			Return(nil)

		csrfRepo.EXPECT().
			Add(token).
			Return(nil)

		ok, err := cryptHashToken.CheckToken(sID, token)
		assert.NoError(t, err)
		assert.Equal(t, ok, true)
	})

	t.Run("TestToken2", func(t *testing.T) {

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		csrfRepo := csrf_mock.NewMockRepository(ctrl)

		cryptHashToken := NewCsrfUsecase(csrfRepo)
		assert.NoError(t, err)

		token, err := cryptHashToken.CreateToken(sID, time.Now().Unix())
		assert.NoError(t, err)

		csrfRepo.EXPECT().
			Check(token).
			Return(testError)

		ok, err := cryptHashToken.CheckToken(sID, token)
		assert.NoError(t, err)
		assert.Equal(t, ok, false)
	})

	t.Run("TestToken3", func(t *testing.T) {

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		csrfRepo := csrf_mock.NewMockRepository(ctrl)

		cryptHashToken := NewCsrfUsecase(csrfRepo)
		assert.NoError(t, err)

		token, err := cryptHashToken.CreateToken(sID, time.Now().Unix())
		assert.NoError(t, err)

		csrfRepo.EXPECT().
			Check(token).
			Return(nil)

		csrfRepo.EXPECT().
			Add(token).
			Return(testError)

		ok, err := cryptHashToken.CheckToken(sID, token)
		assert.Error(t, err)
		assert.Equal(t, ok, false)
	})
}
