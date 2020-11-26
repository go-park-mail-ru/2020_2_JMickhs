package userUsecase

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	customerror "github.com/go-park-mail-ru/2020_2_JMickhs/package/error"

	user_mock "github.com/go-park-mail-ru/2020_2_JMickhs/user/internal/user/mocks"

	"github.com/go-park-mail-ru/2020_2_JMickhs/package/clientError"
	"github.com/go-park-mail-ru/2020_2_JMickhs/user/internal/user/models"
	"github.com/go-playground/validator/v10"

	"github.com/golang/mock/gomock"
)

func TestUserUseCase_GetUserByID(t *testing.T) {
	mockUser := models.User{Username: "kotik", Email: "kek@mail.ru"}
	t.Run("GetUserByID", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockUserRepo := user_mock.NewMockRepository(ctrl)
		mockUser := models.User{Username: "kotik", Email: "kek@mail.ru"}

		mockUserRepo.EXPECT().
			GetUserByID(mockUser.ID).
			Return(mockUser, nil).Times(1)

		u := NewUserUsecase(mockUserRepo, validator.New())

		user, err := u.GetUserByID(mockUser.ID)

		assert.NoError(t, err)
		assert.NotNil(t, user)
	})

	t.Run("GetUserByID-error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockUserRepo := user_mock.NewMockRepository(ctrl)
		mockUserRepo.EXPECT().
			GetUserByID(mockUser.ID).
			Return(mockUser, errors.New("fdsfsd")).Times(1)

		u := NewUserUsecase(mockUserRepo, validator.New())
		_, err := u.GetUserByID(mockUser.ID)

		assert.Error(t, err)
	})
}

func TestUserUseCase_Add(t *testing.T) {

	mockUser := models.User{Username: "kotik", Email: "kek@mail.ru", Password: "12345"}

	t.Run("Add", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockUserRepo := user_mock.NewMockRepository(ctrl)

		mockUserRepo.EXPECT().
			Add(gomock.Any()).
			Return(mockUser, nil).Times(1)

		mockUserRepo.EXPECT().
			GenerateHashFromPassword("12345").
			Return([]byte("newPassword"), nil).Times(1)

		u := NewUserUsecase(mockUserRepo, validator.New())

		user, err := u.Add(mockUser)

		assert.NoError(t, err)
		assert.NotNil(t, user)

	})
	t.Run("AddUser-error2", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		testUser := models.User{Username: "авы", Email: "kek", Password: "12345"}
		mockUserRepoErr := user_mock.NewMockRepository(ctrl)
		uEr := NewUserUsecase(mockUserRepoErr, validator.New())
		_, err := uEr.Add(testUser)

		assert.Error(t, err)
		assert.Equal(t, customerror.ParseCode(err), clientError.BadRequest)
	})
	t.Run("AddUser-error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockUserRepoErr := user_mock.NewMockRepository(ctrl)
		mockUserRepoErr.EXPECT().
			Add(gomock.Any()).
			Return(models.User{}, errors.New("fdsfs")).Times(1)

		mockUserRepoErr.EXPECT().
			GenerateHashFromPassword(gomock.Any()).
			Return([]byte("newPassword"), nil).Times(1)

		uEr := NewUserUsecase(mockUserRepoErr, validator.New())
		_, err := uEr.Add(mockUser)

		assert.Error(t, err)
	})

}

func TestUserUseCase_GetByUserName(t *testing.T) {
	mockUser := models.User{Username: "kotik", Email: "kek@mail.ru"}
	t.Run("Add", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockUserRepo := user_mock.NewMockRepository(ctrl)

		mockUserRepo.EXPECT().
			GetByUserName("kotik").
			Return(mockUser, nil).Times(1)

		u := NewUserUsecase(mockUserRepo, validator.New())

		user, err := u.GetByUserName("kotik")

		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, user.Username, mockUser.Username)
	})

	t.Run("GetByUserName-error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mockUserRepoErr := user_mock.NewMockRepository(ctrl)

		mockUserRepoErr.EXPECT().
			GetByUserName("kotik").
			Return(mockUser, errors.New("fsdfs")).Times(1)

		uEr := NewUserUsecase(mockUserRepoErr, validator.New())

		_, err := uEr.GetByUserName("kotik")
		assert.Error(t, err)
	})

}

func TestUserUseCase_UpdateUser(t *testing.T) {
	mockUser := models.User{Username: "kotik", Email: "kek@mail.ru"}
	t.Run("UpdateUser", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mockUserRepo := user_mock.NewMockRepository(ctrl)
		mockUserRepo.EXPECT().UpdateUser(mockUser).Return(nil).Times(1)
		u := NewUserUsecase(mockUserRepo, validator.New())

		err := u.UpdateUser(mockUser)
		assert.NoError(t, err)
	})

	t.Run("UpdateUser-error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mockUserRepoErr := user_mock.NewMockRepository(ctrl)

		mockUserRepoErr.EXPECT().
			UpdateUser(mockUser).
			Return(errors.New("fdsfs")).Times(1)

		u := NewUserUsecase(mockUserRepoErr, validator.New())
		err := u.UpdateUser(mockUser)

		assert.Error(t, err)
	})
}

func TestUserUseCase_UpdatePassword(t *testing.T) {

	mockUser := models.User{Username: "kotik", Email: "kek@mail.ru"}
	t.Run("UpdatePassword", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mockUserRepo := user_mock.NewMockRepository(ctrl)

		mockUserRepo.EXPECT().
			UpdatePassword(gomock.Any()).
			Return(nil).Times(1)

		mockUserRepo.EXPECT().
			GenerateHashFromPassword(gomock.Any()).
			Return([]byte("newPassword"), nil).Times(1)

		u := NewUserUsecase(mockUserRepo, validator.New())

		err := u.UpdatePassword(mockUser)

		assert.NoError(t, err)

	})

	t.Run("UpdatePassword-error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mockUserRepoErr := user_mock.NewMockRepository(ctrl)

		mockUserRepoErr.EXPECT().
			UpdatePassword(gomock.Any()).
			Return(customerror.NewCustomError(errors.New(""), http.StatusInternalServerError, 1)).Times(1)

		mockUserRepoErr.EXPECT().
			GenerateHashFromPassword(gomock.Any()).
			Return([]byte("newPassword"), nil).Times(1)

		u := NewUserUsecase(mockUserRepoErr, validator.New())
		err := u.UpdatePassword(mockUser)

		assert.Error(t, err)
	})

}

func TestUserUseCase_UpdateAvatar(t *testing.T) {

	mockUser := models.User{Username: "kotik", Email: "kek@mail.ru"}
	t.Run("UpdateAvatar-error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mockUserRepo := user_mock.NewMockRepository(ctrl)

		mockUserRepo.EXPECT().
			UpdateAvatar(gomock.Eq(mockUser)).
			Return(nil).Times(1)

		u := NewUserUsecase(mockUserRepo, validator.New())
		err := u.UpdateAvatar(mockUser)

		assert.NoError(t, err)
	})

	t.Run("UpdateAvatar-error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mockUserRepoErr := user_mock.NewMockRepository(ctrl)

		mockUserRepoErr.EXPECT().
			UpdateAvatar(gomock.Eq(mockUser)).
			Return(errors.New("fdsfsd")).Times(1)

		u := NewUserUsecase(mockUserRepoErr, validator.New())
		err := u.UpdateAvatar(mockUser)

		assert.Error(t, err)
	})
}

func TestUserUseCase_SetDefaultAvatar(t *testing.T) {
	mockUserRepo := new(user_mock.MockRepository)
	mockUser := models.User{Username: "kotik", Email: "kek@mail.ru"}
	u := NewUserUsecase(mockUserRepo, validator.New())

	err := u.SetDefaultAvatar(&mockUser)

	assert.NoError(t, err)

}
