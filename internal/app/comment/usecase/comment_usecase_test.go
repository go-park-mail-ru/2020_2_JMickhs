package commentUsecase

import (
	"errors"
	"testing"

	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/pkg/clientError"

	customerror "github.com/go-park-mail-ru/2020_2_JMickhs/internal/pkg/error"
	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/pkg/serverError"

	comment_mock "github.com/go-park-mail-ru/2020_2_JMickhs/internal/app/comment/mocks"

	commModel "github.com/go-park-mail-ru/2020_2_JMickhs/internal/app/comment/models"

	"github.com/bxcodec/faker/v3"
	paginationModel "github.com/go-park-mail-ru/2020_2_JMickhs/internal/app/paginator/model"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCommentUseCase_GetComments(t *testing.T) {

	testHotels := []commModel.FullCommentInfo{}
	err := faker.FakeData(&testHotels)
	paginfo := paginationModel.PaginationInfo{PageNum: 3, HasNext: true, HasPrev: true, NumPages: 4}

	searchTestData := paginationModel.PaginationModel{List: testHotels, PagInfo: paginfo}
	if err != nil {
		t.Fatalf("an error '%s' was not expected when create fake data", err)
	}
	t.Run("GetComments", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockCommentRepo := comment_mock.NewMockRepository(ctrl)

		mockCommentRepo.EXPECT().
			GetCommentsCount(3).
			Return(20, nil)
		mockCommentRepo.EXPECT().
			GetComments(3, 10).
			Return(testHotels, nil)

		u := NewCommentUsecase(mockCommentRepo)

		hotels, err := u.GetComments(3, 2)

		assert.NoError(t, err)
		assert.Equal(t, hotels, searchTestData)
	})
	t.Run("GetCommentsErr1", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockCommentRepo := comment_mock.NewMockRepository(ctrl)

		mockCommentRepo.EXPECT().
			GetCommentsCount(3).
			Return(20, customerror.NewCustomError(errors.New(""), serverError.ServerInternalError, 1))

		u := NewCommentUsecase(mockCommentRepo)

		_, err := u.GetComments(3, 2)

		assert.Error(t, err)
		assert.Equal(t, customerror.ParseCode(err), serverError.ServerInternalError)
	})
	t.Run("GetCommentsErr2", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockCommentRepo := comment_mock.NewMockRepository(ctrl)

		mockCommentRepo.EXPECT().
			GetCommentsCount(3).
			Return(20, nil)
		mockCommentRepo.EXPECT().
			GetComments(3, 10).
			Return(testHotels, customerror.NewCustomError(errors.New(""), serverError.ServerInternalError, 1))

		u := NewCommentUsecase(mockCommentRepo)

		_, err := u.GetComments(3, 2)

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

		u := NewCommentUsecase(mockCommentRepo)

		comment, err := u.AddComment(testComment)

		assert.NoError(t, err)
		assert.Equal(t, comment.Comment, GetComment)
		assert.Equal(t, comment.Rate, 3.8)
	})

	t.Run("AddCommentErr1", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockCommentRepo := comment_mock.NewMockRepository(ctrl)
		GetComment := commModel.Comment{CommID: 3, HotelID: 2, UserID: 1, Message: "fsdfsdfsd", Rate: 3, Time: "22-02-2000"}
		mockCommentRepo.EXPECT().
			AddComment(testComment).
			Return(GetComment, customerror.NewCustomError(errors.New(""), clientError.Locked, 1))

		u := NewCommentUsecase(mockCommentRepo)

		_, err := u.AddComment(testComment)

		assert.Error(t, err)
		assert.Equal(t, customerror.ParseCode(err), clientError.Locked)
	})

	t.Run("AddCommentErr2", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockCommentRepo := comment_mock.NewMockRepository(ctrl)
		GetComment := commModel.Comment{CommID: 3, HotelID: 2, UserID: 1, Message: "fsdfsdfsd", Rate: 3, Time: "22-02-2000"}
		mockCommentRepo.EXPECT().
			AddComment(testComment).
			Return(GetComment, nil)

		rateInf := commModel.RateInfo{RatesCount: 4, CurrRating: 4}
		mockCommentRepo.EXPECT().
			GetCurrentRating(2).
			Return(rateInf, customerror.NewCustomError(errors.New(""), serverError.ServerInternalError, 1))

		u := NewCommentUsecase(mockCommentRepo)

		_, err := u.AddComment(testComment)

		assert.Error(t, err)
		assert.Equal(t, customerror.ParseCode(err), serverError.ServerInternalError)
	})

	t.Run("AddCommentErr2", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockCommentRepo := comment_mock.NewMockRepository(ctrl)
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

		u := NewCommentUsecase(mockCommentRepo)

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

		u := NewCommentUsecase(mockCommentRepo)

		comment, err := u.UpdateComment(testComment)

		assert.NoError(t, err)
		assert.Equal(t, comment.Rate, 4.0)
	})

	t.Run("UpdateCommentErr1", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockCommentRepo := comment_mock.NewMockRepository(ctrl)
		mockCommentRepo.EXPECT().
			CheckUser(&testComment).
			Return(6, customerror.NewCustomError(errors.New(""), serverError.ServerInternalError, 1))

		u := NewCommentUsecase(mockCommentRepo)

		_, err := u.UpdateComment(testComment)

		assert.Error(t, err)
		assert.Equal(t, customerror.ParseCode(err), serverError.ServerInternalError)
	})

	t.Run("UpdateCommentErr1", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockCommentRepo := comment_mock.NewMockRepository(ctrl)
		mockCommentRepo.EXPECT().
			CheckUser(&testComment).
			Return(6, nil)

		mockCommentRepo.EXPECT().
			UpdateComment(&testComment).
			Return(customerror.NewCustomError(errors.New(""), serverError.ServerInternalError, 1))

		u := NewCommentUsecase(mockCommentRepo)

		_, err := u.UpdateComment(testComment)

		assert.Error(t, err)
		assert.Equal(t, customerror.ParseCode(err), serverError.ServerInternalError)
	})

	t.Run("UpdateCommentErr1", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockCommentRepo := comment_mock.NewMockRepository(ctrl)
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

		u := NewCommentUsecase(mockCommentRepo)

		_, err := u.UpdateComment(testComment)

		assert.Error(t, err)
		assert.Equal(t, customerror.ParseCode(err), serverError.ServerInternalError)
	})

	t.Run("UpdateCommentErr2", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockCommentRepo := comment_mock.NewMockRepository(ctrl)
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

		u := NewCommentUsecase(mockCommentRepo)

		_, err := u.UpdateComment(testComment)

		assert.Error(t, err)
		assert.Equal(t, customerror.ParseCode(err), serverError.ServerInternalError)
	})
}
