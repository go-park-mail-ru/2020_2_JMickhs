package userGrpcDelivery

import (
	"context"
	"errors"
	"testing"

	userService "github.com/go-park-mail-ru/2020_2_JMickhs/package/proto/user"
	user_mock "github.com/go-park-mail-ru/2020_2_JMickhs/user/internal/user/mocks"
	"github.com/go-park-mail-ru/2020_2_JMickhs/user/internal/user/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUserDelivery_GetUserByID(t *testing.T) {
	user := models.User{ID: 3, Username: "fdsf", Email: "fdsfcx"}
	t.Run("GetUserByID", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		userMock := user_mock.NewMockUsecase(ctrl)
		in := &userService.UserID{UserID: 3}

		delivery := NewUserDelivery(userMock)

		userMock.EXPECT().
			GetUserByID(int(in.UserID)).
			Return(user, nil)

		userResponse, err := delivery.GetUserByID(context.Background(), in)
		assert.NoError(t, err)
		assert.Equal(t, user.Username, userResponse.Username)
	})

	t.Run("GetUserByIDErr", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		userMock := user_mock.NewMockUsecase(ctrl)
		in := &userService.UserID{UserID: 2}

		delivery := NewUserDelivery(userMock)

		testErr := errors.New("fds")
		userMock.EXPECT().
			GetUserByID(int(in.UserID)).
			Return(user, testErr)

		_, err := delivery.GetUserByID(context.Background(), in)
		assert.Error(t, err, testErr)
	})
}
