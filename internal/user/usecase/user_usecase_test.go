package userUsecase

import (
	"errors"
	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/user/mocks"
	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/user/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"mime/multipart"
	"testing"
)

func TestUserUseCase_GetUserByID(t *testing.T) {
	mockUser := models.User{Username: "kotik",Email: "kek@mail.ru"}
	t.Run("GetUserByID",func(t *testing.T){
		mockUserRepo := new(mocks.UserRepository)
		mockUser := models.User{Username: "kotik",Email: "kek@mail.ru"}
		mockUserRepo.On("GetUserByID",mock.AnythingOfType("int")).Return(mockUser, nil).Once()
		u := NewUserUsecase(mockUserRepo)

		user, err := u.GetUserByID(mockUser.ID)


		assert.NoError(t,err)
		assert.NotNil(t,user)
		mockUserRepo.AssertExpectations(t)
	})


	t.Run("GetUserByID-error", func(t *testing.T) {
		mockUserRepoErr := new(mocks.UserRepository)
		mockUserRepoErr.On("GetUserByID",mock.AnythingOfType("int")).Return(mockUser, errors.New("fsdsd")).Once()
		u := NewUserUsecase(mockUserRepoErr)
		_, err := u.GetUserByID(mockUser.ID)

		assert.Error(t,err)
		mockUserRepoErr.AssertExpectations(t)
	})
}

func TestUserUseCase_Add(t *testing.T) {

	mockUser := models.User{Username: "kotik",Email: "kek@mail.ru"}

	t.Run("Add",func(t *testing.T) {
		mockUserRepo := new(mocks.UserRepository)

		mockUserRepo.On("Add", mock.AnythingOfType("models.User")).Return(models.User{}, nil).Once()
		u := NewUserUsecase(mockUserRepo)

		user, err := u.Add(mockUser)

		assert.NoError(t, err)
		assert.NotNil(t, user)

		mockUserRepo.AssertExpectations(t)
	})
	t.Run("AddUser-error", func(t *testing.T) {
		mockUserRepoErr := new(mocks.UserRepository)
		mockUserRepoErr.On("Add",mock.AnythingOfType("models.User")).Return(models.User{}, errors.New("fsdsd")).Once()
		uEr:= NewUserUsecase(mockUserRepoErr)
		_, err := uEr.Add(mockUser)

		assert.Error(t,err)
		mockUserRepoErr.AssertExpectations(t)
	})

}

func TestUserUseCase_GetByUserName(t *testing.T) {
	mockUser := models.User{Username: "kotik",Email: "kek@mail.ru"}
	t.Run("Add",func(t *testing.T) {

		mockUserRepo := new(mocks.UserRepository)
		mockUserRepo.On("GetByUserName",mock.AnythingOfType("string")).Return(mockUser, nil).Once()
		u := NewUserUsecase(mockUserRepo)

		user, err := u.GetByUserName("kotik")

		assert.NoError(t,err)
		assert.NotNil(t,user)
		assert.Equal(t,user.Username,mockUser.Username)
		mockUserRepo.AssertExpectations(t)
	})

	t.Run("GetByUserName-error",func(t *testing.T) {
		mockUserRepoErr := new(mocks.UserRepository)
		mockUserRepoErr.On("GetByUserName",mock.AnythingOfType("string")).Return(mockUser, errors.New("cxew")).Once()
		uEr := NewUserUsecase(mockUserRepoErr)

		_, err := uEr.GetByUserName("kotik")
		assert.Error(t,err)
		mockUserRepoErr.AssertExpectations(t)
	})

}

func TestUserUseCase_UpdateUser(t *testing.T) {
	mockUser := models.User{Username: "kotik",Email: "kek@mail.ru"}
	t.Run("UpdateUser", func(t *testing.T){
		mockUserRepo := new(mocks.UserRepository)
		mockUserRepo.On("UpdateUser",mock.AnythingOfType("models.User")).Return(nil).Once()
		u := NewUserUsecase(mockUserRepo)

		err := u.UpdateUser(mockUser)
		assert.NoError(t,err)
		mockUserRepo.AssertExpectations(t)
	})

	t.Run("UpdateUser-error", func(t *testing.T) {
		mockUserRepoErr := new(mocks.UserRepository)
		mockUserRepoErr.On("UpdateUser",mock.AnythingOfType("models.User")).Return(errors.New("fsdsd")).Once()
		u := NewUserUsecase(mockUserRepoErr)
		err := u.UpdateUser(mockUser)

		assert.Error(t,err)
		mockUserRepoErr.AssertExpectations(t)
	})
}



func TestUserUseCase_UpdatePassword(t *testing.T) {

	mockUser := models.User{Username: "kotik",Email: "kek@mail.ru"}
	t.Run("UpdatePassword", func(t *testing.T) {
		mockUserRepo := new(mocks.UserRepository)

		mockUserRepo.On("UpdatePassword", mock.AnythingOfType("models.User")).Return(nil).Once()
		u := NewUserUsecase(mockUserRepo)

		err := u.UpdatePassword(mockUser)

		assert.NoError(t, err)

		mockUserRepo.AssertExpectations(t)
	})

	t.Run("UpdatePassword-error", func(t *testing.T) {
		mockUserRepoErr := new(mocks.UserRepository)
		mockUserRepoErr.On("UpdatePassword",mock.AnythingOfType("models.User")).Return(errors.New("fsdsd")).Once()
		u := NewUserUsecase(mockUserRepoErr)
		err := u.UpdatePassword(mockUser)

		assert.Error(t,err)
		mockUserRepoErr.AssertExpectations(t)
	})

}

func TestUserUseCase_UpdateAvatar(t *testing.T) {

	mockUser := models.User{Username: "kotik",Email: "kek@mail.ru"}
	t.Run("UpdateAvatar-error", func(t *testing.T) {
		mockUserRepo := new(mocks.UserRepository)

		mockUserRepo.On("UpdateAvatar",mock.AnythingOfType("models.User")).Return(nil).Once()
		u := NewUserUsecase(mockUserRepo)
		err := u.UpdateAvatar(mockUser)

		assert.NoError(t,err)
		mockUserRepo.AssertExpectations(t)
	})

	t.Run("UpdateAvatar-error", func(t *testing.T) {
		mockUserRepoErr := new(mocks.UserRepository)
		mockUserRepoErr.On("UpdateAvatar",mock.AnythingOfType("models.User")).Return(errors.New("fsdsd")).Once()
		u := NewUserUsecase(mockUserRepoErr)
		err := u.UpdateAvatar(mockUser)

		assert.Error(t,err)
		mockUserRepoErr.AssertExpectations(t)
	})
}

func TestUserUseCase_SetDefaultAvatar(t *testing.T) {
	mockUserRepo := new(mocks.UserRepository)
	mockUser := models.User{Username: "kotik",Email: "kek@mail.ru"}
	u := NewUserUsecase(mockUserRepo)

	err := u.SetDefaultAvatar(&mockUser)

	assert.NoError(t,err)

	mockUserRepo.AssertExpectations(t)

}

func TestUserUseCase_UploadAvatar(t *testing.T) {
	mockUserRepo := new(mocks.UserRepository)
	mockUser := models.User{Username: "kotik",Email: "kek@mail.ru"}
	u := NewUserUsecase(mockUserRepo)
	file := new(multipart.File)
	fileType := "png"

	err := u.UploadAvatar(*file,fileType,&mockUser)

	assert.Error(t,err)

	mockUserRepo.AssertExpectations(t)

}