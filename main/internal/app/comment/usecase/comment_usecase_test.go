package commentUsecase

import (
	"errors"
	"testing"

	customerror "github.com/go-park-mail-ru/2020_2_JMickhs/package/error"
	userService "github.com/go-park-mail-ru/2020_2_JMickhs/package/proto/user"

	comment_mock "github.com/go-park-mail-ru/2020_2_JMickhs/main/internal/app/comment/mocks"
	"github.com/go-park-mail-ru/2020_2_JMickhs/package/clientError"
	"github.com/go-park-mail-ru/2020_2_JMickhs/package/serverError"

	commModel "github.com/go-park-mail-ru/2020_2_JMickhs/main/internal/app/comment/models"

	paginationModel "github.com/go-park-mail-ru/2020_2_JMickhs/main/internal/app/paginator/model"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCommentUseCase_GetComments(t *testing.T) {

	testComments := []commModel.FullCommentInfo{
		{3, 1, 3, "kekw", 4, "avatar/kek.jpeg", "kostikan", "20.20.2010"},
	}
	paginfo := paginationModel.PaginationInfo{NextLink: "",
		PrevLink: "", ItemsCount: 20}

	searchTestData := commModel.Comments{Comments: testComments, Info: paginfo}
	t.Run("GetComments", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		defer ctrl.Finish()

		mockCommentRepo := comment_mock.NewMockRepository(ctrl)
		mockUserService := userService.NewMockUserServiceClient(ctrl)

		mockCommentRepo.EXPECT().
			GetCommentsCount(3).
			Return(20, nil)
		mockCommentRepo.EXPECT().
			GetComments("3", 1, "2", 3).
			Return(testComments, nil)

		mockCommentRepo.EXPECT().
			CheckRateExistForComments(3, 3).
			Return(false, nil)
		mockUserService.EXPECT().
			GetUserByID(gomock.Any(), &userService.UserID{UserID: 3}).
			Return(&userService.User{UserID: 3, Username: "kostikan", Email: "email@mail.ru", Avatar: "avatar/kek.jpeg"}, nil)

		u := NewCommentUsecase(mockCommentRepo, mockUserService)

		_, comments, err := u.GetComments("3", "1", "2", 3)

		assert.NoError(t, err)
		assert.Equal(t, comments, searchTestData)
	})
	t.Run("GetComments1", func(t *testing.T) {
		paginfo = paginationModel.PaginationInfo{NextLink: "",
			PrevLink: "", ItemsCount: 3}
		searchTestData1 := commModel.Comments{Comments: testComments, Info: paginfo}
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockCommentRepo := comment_mock.NewMockRepository(ctrl)
		mockUserService := userService.NewMockUserServiceClient(ctrl)

		mockCommentRepo.EXPECT().
			GetCommentsCount(3).
			Return(3, nil)
		mockCommentRepo.EXPECT().
			GetComments("3", 1, "3", 3).
			Return(testComments, nil)

		mockCommentRepo.EXPECT().
			CheckRateExistForComments(3, 3).
			Return(false, nil)

		mockUserService.EXPECT().
			GetUserByID(gomock.Any(), &userService.UserID{UserID: 3}).
			Return(&userService.User{UserID: 3, Username: "kostikan", Email: "email@mail.ru", Avatar: "avatar/kek.jpeg"}, nil)

		u := NewCommentUsecase(mockCommentRepo, mockUserService)

		_, comments, err := u.GetComments("3", "1", "3", 3)

		assert.NoError(t, err)
		assert.Equal(t, comments, searchTestData1)
	})
	t.Run("GetComments2", func(t *testing.T) {
		paginfo = paginationModel.PaginationInfo{NextLink: "",
			PrevLink: "", ItemsCount: 3}
		searchTestData1 := commModel.Comments{Comments: testComments, Info: paginfo}
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockCommentRepo := comment_mock.NewMockRepository(ctrl)
		mockUserService := userService.NewMockUserServiceClient(ctrl)

		mockCommentRepo.EXPECT().
			GetCommentsCount(3).
			Return(3, nil)
		mockCommentRepo.EXPECT().
			GetComments("3", 1, "0", 3).
			Return(testComments, nil)

		mockCommentRepo.EXPECT().
			CheckRateExistForComments(3, 3).
			Return(false, nil)

		mockUserService.EXPECT().
			GetUserByID(gomock.Any(), &userService.UserID{UserID: 3}).
			Return(&userService.User{UserID: 3, Username: "kostikan", Email: "email@mail.ru", Avatar: "avatar/kek.jpeg"}, nil)

		u := NewCommentUsecase(mockCommentRepo, mockUserService)

		_, comments, err := u.GetComments("3", "1", "0", 3)

		assert.NoError(t, err)
		assert.Equal(t, comments, searchTestData1)
	})
	t.Run("GetCommentsErr1", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockCommentRepo := comment_mock.NewMockRepository(ctrl)
		mockUserService := userService.NewMockUserServiceClient(ctrl)

		mockCommentRepo.EXPECT().
			GetCommentsCount(3).
			Return(20, customerror.NewCustomError(errors.New(""), serverError.ServerInternalError, 1))

		u := NewCommentUsecase(mockCommentRepo, mockUserService)

		_, _, err := u.GetComments("3", "1", "2", 3)

		assert.Error(t, err)
		assert.Equal(t, customerror.ParseCode(err), serverError.ServerInternalError)
	})
	t.Run("GetCommentsErr2", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockCommentRepo := comment_mock.NewMockRepository(ctrl)
		mockUserService := userService.NewMockUserServiceClient(ctrl)

		mockCommentRepo.EXPECT().
			GetCommentsCount(3).
			Return(20, nil)
		mockCommentRepo.EXPECT().
			GetComments("3", 1, "2", 3).
			Return(testComments, customerror.NewCustomError(errors.New(""), serverError.ServerInternalError, 1))

		mockCommentRepo.EXPECT().
			CheckRateExistForComments(3, 3).
			Return(false, nil)

		u := NewCommentUsecase(mockCommentRepo, mockUserService)

		_, _, err := u.GetComments("3", "1", "2", 3)

		assert.Error(t, err)
		assert.Equal(t, customerror.ParseCode(err), serverError.ServerInternalError)
	})
}

func TestCommentUseCase_AddComment(t *testing.T) {
	testComment := commModel.Comment{CommID: 3, HotelID: 2, UserID: 1, Message: "fsdfsdfsd", Rate: 3}

	t.Run("AddComment", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockCommentRepo := comment_mock.NewMockRepository(ctrl)
		mockUserService := userService.NewMockUserServiceClient(ctrl)
		GetComment := commModel.Comment{CommID: 3, HotelID: 2, UserID: 1, Message: "fsdfsdfsd", Rate: 3, Time: "22-02-2000"}
		mockCommentRepo.EXPECT().
			AddComment(testComment).
			Return(GetComment, nil)

		rateInf := commModel.RateInfo{RatesCount: 4, CurrRating: 4}
		mockCommentRepo.EXPECT().
			GetCurrentRating(2).
			Return(rateInf, nil)

		mockCommentRepo.EXPECT().
			UpdateHotelRating(2, 3.75).
			Return(nil)

		u := NewCommentUsecase(mockCommentRepo, mockUserService)

		comment, err := u.AddComment(testComment)

		assert.NoError(t, err)
		assert.Equal(t, comment.Comment, GetComment)
		assert.Equal(t, comment.Rate, 3.8)
	})

	t.Run("AddCommentErr1", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockCommentRepo := comment_mock.NewMockRepository(ctrl)
		mockUserService := userService.NewMockUserServiceClient(ctrl)
		GetComment := commModel.Comment{CommID: 3, HotelID: 2, UserID: 1, Message: "fsdfsdfsd", Rate: 3, Time: "22-02-2000"}
		mockCommentRepo.EXPECT().
			AddComment(testComment).
			Return(GetComment, customerror.NewCustomError(errors.New(""), clientError.Locked, 1))

		u := NewCommentUsecase(mockCommentRepo, mockUserService)

		_, err := u.AddComment(testComment)

		assert.Error(t, err)
		assert.Equal(t, customerror.ParseCode(err), clientError.Locked)
	})

	t.Run("AddCommentErr2", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockCommentRepo := comment_mock.NewMockRepository(ctrl)
		mockUserService := userService.NewMockUserServiceClient(ctrl)

		GetComment := commModel.Comment{CommID: 3, HotelID: 2, UserID: 1, Message: "fsdfsdfsd", Rate: 3, Time: "22-02-2000"}
		mockCommentRepo.EXPECT().
			AddComment(testComment).
			Return(GetComment, nil)

		rateInf := commModel.RateInfo{RatesCount: 4, CurrRating: 4}
		mockCommentRepo.EXPECT().
			GetCurrentRating(2).
			Return(rateInf, customerror.NewCustomError(errors.New(""), serverError.ServerInternalError, 1))

		u := NewCommentUsecase(mockCommentRepo, mockUserService)

		_, err := u.AddComment(testComment)

		assert.Error(t, err)
		assert.Equal(t, customerror.ParseCode(err), serverError.ServerInternalError)
	})

	t.Run("AddCommentErr2", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockCommentRepo := comment_mock.NewMockRepository(ctrl)
		mockUserService := userService.NewMockUserServiceClient(ctrl)
		GetComment := commModel.Comment{CommID: 3, HotelID: 2, UserID: 1, Message: "fsdfsdfsd", Rate: 3, Time: "22-02-2000"}
		mockCommentRepo.EXPECT().
			AddComment(testComment).
			Return(GetComment, nil)

		rateInf := commModel.RateInfo{RatesCount: 4, CurrRating: 4}
		mockCommentRepo.EXPECT().
			GetCurrentRating(2).
			Return(rateInf, nil)

		mockCommentRepo.EXPECT().
			UpdateHotelRating(2, 3.75).
			Return(customerror.NewCustomError(errors.New(""), serverError.ServerInternalError, 1))

		u := NewCommentUsecase(mockCommentRepo, mockUserService)

		_, err := u.AddComment(testComment)

		assert.Error(t, err)
		assert.Equal(t, customerror.ParseCode(err), serverError.ServerInternalError)
	})
}

func TestCommentUseCase_UpdateComment(t *testing.T) {
	testComment := commModel.Comment{CommID: 3, HotelID: 2, UserID: 1, Message: "fsdfsdfsd", Rate: 3}

	t.Run("UpdateComment", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockCommentRepo := comment_mock.NewMockRepository(ctrl)
		mockUserService := userService.NewMockUserServiceClient(ctrl)

		mockCommentRepo.EXPECT().
			CheckUser(&testComment).
			Return(6, nil)

		mockCommentRepo.EXPECT().
			UpdateComment(&testComment).
			Return(nil)

		rateInf := commModel.RateInfo{RatesCount: 3, CurrRating: 5}
		mockCommentRepo.EXPECT().
			GetCurrentRating(2).
			Return(rateInf, nil)

		mockCommentRepo.EXPECT().
			UpdateHotelRating(2, 4.0).
			Return(nil)

		u := NewCommentUsecase(mockCommentRepo, mockUserService)

		comment, err := u.UpdateComment(testComment)

		assert.NoError(t, err)
		assert.Equal(t, comment.Rate, 4.0)
	})

	t.Run("UpdateCommentErr1", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockCommentRepo := comment_mock.NewMockRepository(ctrl)
		mockUserService := userService.NewMockUserServiceClient(ctrl)

		mockCommentRepo.EXPECT().
			CheckUser(&testComment).
			Return(6, customerror.NewCustomError(errors.New(""), serverError.ServerInternalError, 1))

		u := NewCommentUsecase(mockCommentRepo, mockUserService)

		_, err := u.UpdateComment(testComment)

		assert.Error(t, err)
		assert.Equal(t, customerror.ParseCode(err), serverError.ServerInternalError)
	})

	t.Run("UpdateCommentErr1", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockCommentRepo := comment_mock.NewMockRepository(ctrl)
		mockUserService := userService.NewMockUserServiceClient(ctrl)

		mockCommentRepo.EXPECT().
			CheckUser(&testComment).
			Return(6, nil)

		mockCommentRepo.EXPECT().
			UpdateComment(&testComment).
			Return(customerror.NewCustomError(errors.New(""), serverError.ServerInternalError, 1))

		u := NewCommentUsecase(mockCommentRepo, mockUserService)

		_, err := u.UpdateComment(testComment)

		assert.Error(t, err)
		assert.Equal(t, customerror.ParseCode(err), serverError.ServerInternalError)
	})

	t.Run("UpdateCommentErr1", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockCommentRepo := comment_mock.NewMockRepository(ctrl)
		mockUserService := userService.NewMockUserServiceClient(ctrl)

		mockCommentRepo.EXPECT().
			CheckUser(&testComment).
			Return(6, nil)

		mockCommentRepo.EXPECT().
			UpdateComment(&testComment).
			Return(nil)

		rateInf := commModel.RateInfo{RatesCount: 3, CurrRating: 5}
		mockCommentRepo.EXPECT().
			GetCurrentRating(2).
			Return(rateInf, customerror.NewCustomError(errors.New(""), serverError.ServerInternalError, 1))

		u := NewCommentUsecase(mockCommentRepo, mockUserService)

		_, err := u.UpdateComment(testComment)

		assert.Error(t, err)
		assert.Equal(t, customerror.ParseCode(err), serverError.ServerInternalError)
	})

	t.Run("UpdateCommentErr2", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockCommentRepo := comment_mock.NewMockRepository(ctrl)
		mockUserService := userService.NewMockUserServiceClient(ctrl)

		mockCommentRepo.EXPECT().
			CheckUser(&testComment).
			Return(6, nil)

		mockCommentRepo.EXPECT().
			UpdateComment(&testComment).
			Return(nil)

		rateInf := commModel.RateInfo{RatesCount: 3, CurrRating: 5}
		mockCommentRepo.EXPECT().
			GetCurrentRating(2).
			Return(rateInf, nil)

		mockCommentRepo.EXPECT().
			UpdateHotelRating(2, 4.0).
			Return(customerror.NewCustomError(errors.New(""), serverError.ServerInternalError, 1))

		u := NewCommentUsecase(mockCommentRepo, mockUserService)

		_, err := u.UpdateComment(testComment)

		assert.Error(t, err)
		assert.Equal(t, customerror.ParseCode(err), serverError.ServerInternalError)
	})
}
