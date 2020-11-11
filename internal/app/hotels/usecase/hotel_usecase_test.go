package hotelUsecase

import (
	"errors"
	commModel "github.com/go-park-mail-ru/2020_2_JMickhs/internal/app/comment/models"
	"testing"

	paginationModel "github.com/go-park-mail-ru/2020_2_JMickhs/internal/app/paginator/model"

	"github.com/bxcodec/faker/v3"

	customerror "github.com/go-park-mail-ru/2020_2_JMickhs/internal/pkg/error"
	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/pkg/serverError"

	hotels_mock "github.com/go-park-mail-ru/2020_2_JMickhs/internal/app/hotels/mocks"
	"github.com/golang/mock/gomock"

	hotelmodel "github.com/go-park-mail-ru/2020_2_JMickhs/internal/app/hotels/models"

	"github.com/stretchr/testify/assert"
)

func TestHotelUseCase_GetHotelByID(t *testing.T) {

	mockHotel := hotelmodel.Hotel{Name: "bestHotel", Description: "in the world"}
	t.Run("HotelGetByID", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockHotelRepo := hotels_mock.NewMockRepository(ctrl)

		mockHotelRepo.EXPECT().
			GetHotelByID(3).
			Return(mockHotel, nil)

		u := NewHotelUsecase(mockHotelRepo)

		hotel, err := u.GetHotelByID(3)

		assert.NoError(t, err)
		assert.Equal(t, hotel, mockHotel)
	})
	t.Run("HotelGetByID-error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockHotelRepoErr := hotels_mock.NewMockRepository(ctrl)

		mockHotelRepoErr.EXPECT().
			GetHotelByID(3).
			Return(mockHotel, customerror.NewCustomError(errors.New(""), serverError.ServerInternalError, 1))

		uEr := NewHotelUsecase(mockHotelRepoErr)

		_, err := uEr.GetHotelByID(3)
		assert.Error(t, err)
		assert.Equal(t, customerror.ParseCode(err), serverError.ServerInternalError)
	})
}

func TestHotelUseCase_GetHotels(t *testing.T) {

	testHotels := make([]hotelmodel.Hotel, 4)
	err := faker.FakeData(&testHotels)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when create fake data", err)
	}
	t.Run("HotelGetHotels", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockHotelRepo := hotels_mock.NewMockRepository(ctrl)

		mockHotelRepo.EXPECT().
			GetHotels(0).
			Return(testHotels, nil)

		u := NewHotelUsecase(mockHotelRepo)

		hotels, err := u.GetHotels(0)

		assert.NoError(t, err)
		assert.Equal(t, hotels, testHotels)
	})
	t.Run("HotelGetHotelsErr", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockHotelRepo := hotels_mock.NewMockRepository(ctrl)

		mockHotelRepo.EXPECT().
			GetHotels(0).
			Return(testHotels, customerror.NewCustomError(errors.New(""), serverError.ServerInternalError, 1))

		u := NewHotelUsecase(mockHotelRepo)

		_, err := u.GetHotels(0)

		assert.Error(t, err)
		assert.Equal(t, customerror.ParseCode(err), serverError.ServerInternalError)
	})
}

func TestHotelUseCase_GetHotelsPreview(t *testing.T) {

	testHotels := []hotelmodel.HotelPreview{}
	err := faker.FakeData(&testHotels)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when create fake data", err)
	}
	t.Run("HotelGetHotels", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockHotelRepo := hotels_mock.NewMockRepository(ctrl)

		mockHotelRepo.EXPECT().
			GetHotelsPreview("Villa").
			Return(testHotels, nil)

		u := NewHotelUsecase(mockHotelRepo)

		hotels, err := u.GetHotelsPreview("Villa")

		assert.NoError(t, err)
		assert.Equal(t, hotels, testHotels)
	})
	t.Run("HotelGetHotelsErr", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockHotelRepo := hotels_mock.NewMockRepository(ctrl)

		mockHotelRepo.EXPECT().
			GetHotelsPreview("Villa").
			Return(testHotels, customerror.NewCustomError(errors.New(""), serverError.ServerInternalError, 1))

		u := NewHotelUsecase(mockHotelRepo)

		_, err := u.GetHotelsPreview("Villa")

		assert.Error(t, err)
		assert.Equal(t, customerror.ParseCode(err), serverError.ServerInternalError)
	})
}

func TestHotelUseCase_FetchHotels(t *testing.T) {

	testHotels := []hotelmodel.Hotel{}
	err := faker.FakeData(&testHotels)
	paginfo := paginationModel.PaginationInfo{ItemsCount: 56, NextLink: "api/v1/hotels/?id=3&limit=1&offset=3",
		PrevLink: "api/v1/hotels/?id=3&limit=1&offset=1"}

	searchTestData := hotelmodel.SearchData{Hotels: testHotels, PagInfo: paginfo}
	if err != nil {
		t.Fatalf("an error '%s' was not expected when create fake data", err)
	}
	t.Run("HotelFetchHotels", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockHotelRepo := hotels_mock.NewMockRepository(ctrl)

		mockHotelRepo.EXPECT().
			FetchHotels("Villa", 56).
			Return(testHotels, nil)

		u := NewHotelUsecase(mockHotelRepo)

		hotels, err := u.FetchHotels("Villa", 2)
		hotels.PagInfo = paginfo

		assert.NoError(t, err)
		assert.Equal(t, hotels, searchTestData)
	})
	t.Run("HotelFetchHotels", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mockHotelRepo := hotels_mock.NewMockRepository(ctrl)

		mockHotelRepo.EXPECT().
			FetchHotels("Villa", 56).
			Return(testHotels, nil)

		u := NewHotelUsecase(mockHotelRepo)

		hotels, err := u.FetchHotels("Villa", 2)
		hotels.PagInfo = paginfo

		assert.NoError(t, err)
		assert.Equal(t, searchTestData, hotels)
	})
	t.Run("HotelFetchHotelsErr2", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockHotelRepo := hotels_mock.NewMockRepository(ctrl)

		mockHotelRepo.EXPECT().
			FetchHotels("Villa", 56).
			Return(testHotels, customerror.NewCustomError(errors.New(""), serverError.ServerInternalError, 1))

		u := NewHotelUsecase(mockHotelRepo)

		_, err := u.FetchHotels("Villa", 2)

		assert.Error(t, err)
		assert.Equal(t, customerror.ParseCode(err), serverError.ServerInternalError)
	})
}

func TestHotelUseCase_CheckRateExist(t *testing.T) {
	comment := commModel.FullCommentInfo{}
	err := faker.FakeData(&comment)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when create fake data", err)
	}
	t.Run("CheckRateExist", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockHotelRepo := hotels_mock.NewMockRepository(ctrl)

		mockHotelRepo.EXPECT().
			CheckRateExist(2, 4).
			Return(comment, nil)

		u := NewHotelUsecase(mockHotelRepo)

		rate, err := u.CheckRateExist(2, 4)

		assert.NoError(t, err)
		assert.Equal(t, rate.Rating, comment.Rating)
	})
	t.Run("CheckRateExistErr", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockHotelRepo := hotels_mock.NewMockRepository(ctrl)

		mockHotelRepo.EXPECT().
			CheckRateExist(2, 4).
			Return(comment, customerror.NewCustomError(errors.New(""), serverError.ServerInternalError, 1))

		u := NewHotelUsecase(mockHotelRepo)

		_, err := u.CheckRateExist(2, 4)

		assert.Error(t, err)
		assert.Equal(t, customerror.ParseCode(err), serverError.ServerInternalError)
	})
}
