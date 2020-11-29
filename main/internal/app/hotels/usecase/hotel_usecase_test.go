package hotelUsecase

import (
	"errors"
	"testing"

	wishlists_mock "github.com/go-park-mail-ru/2020_2_JMickhs/main/internal/app/wishlist/mocks"

	userService "github.com/go-park-mail-ru/2020_2_JMickhs/package/proto/user"

	customerror "github.com/go-park-mail-ru/2020_2_JMickhs/package/error"

	commModel "github.com/go-park-mail-ru/2020_2_JMickhs/main/internal/app/comment/models"
	"github.com/go-park-mail-ru/2020_2_JMickhs/package/serverError"

	paginationModel "github.com/go-park-mail-ru/2020_2_JMickhs/main/internal/app/paginator/model"

	"github.com/bxcodec/faker/v3"

	hotels_mock "github.com/go-park-mail-ru/2020_2_JMickhs/main/internal/app/hotels/mocks"
	"github.com/golang/mock/gomock"

	hotelmodel "github.com/go-park-mail-ru/2020_2_JMickhs/main/internal/app/hotels/models"

	"github.com/stretchr/testify/assert"
)

func TestHotelUseCase_GetHotelByID(t *testing.T) {

	mockHotel := hotelmodel.Hotel{HotelID: 3, Name: "bestHotel", Description: "in the world"}
	t.Run("HotelGetByID", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockHotelRepo := hotels_mock.NewMockRepository(ctrl)
		mockUserService := userService.NewMockUserServiceClient(ctrl)
		mockWishListRepo := wishlists_mock.NewMockRepository(ctrl)

		mockHotelRepo.EXPECT().
			GetHotelByID(3).
			Return(mockHotel, nil)
		mockWishListRepo.EXPECT().
			CheckHotelInWishlists(2, 3).
			Return("", nil)

		u := NewHotelUsecase(mockHotelRepo, mockUserService, mockWishListRepo)

		hotel, err := u.GetHotelByID(3, 2)

		assert.NoError(t, err)
		assert.Equal(t, hotel, mockHotel)
	})
	t.Run("HotelGetByID-error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockHotelRepoErr := hotels_mock.NewMockRepository(ctrl)
		mockUserService := userService.NewMockUserServiceClient(ctrl)
		mockWishListRepo := wishlists_mock.NewMockRepository(ctrl)

		mockHotelRepoErr.EXPECT().
			GetHotelByID(3).
			Return(mockHotel, customerror.NewCustomError(errors.New(""), serverError.ServerInternalError, 1))

		uEr := NewHotelUsecase(mockHotelRepoErr, mockUserService, mockWishListRepo)

		_, err := uEr.GetHotelByID(3, 2)
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
		mockUserService := userService.NewMockUserServiceClient(ctrl)
		mockWishListRepo := wishlists_mock.NewMockRepository(ctrl)

		mockHotelRepo.EXPECT().
			GetHotels(0).
			Return(testHotels, nil)

		u := NewHotelUsecase(mockHotelRepo, mockUserService, mockWishListRepo)

		hotels, err := u.GetHotels(0)

		assert.NoError(t, err)
		assert.Equal(t, hotels, testHotels)
	})
	t.Run("HotelGetHotelsErr", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockHotelRepo := hotels_mock.NewMockRepository(ctrl)
		mockUserService := userService.NewMockUserServiceClient(ctrl)
		mockWishListRepo := wishlists_mock.NewMockRepository(ctrl)

		mockHotelRepo.EXPECT().
			GetHotels(0).
			Return(testHotels, customerror.NewCustomError(errors.New(""), serverError.ServerInternalError, 1))

		u := NewHotelUsecase(mockHotelRepo, mockUserService, mockWishListRepo)

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
		mockUserService := userService.NewMockUserServiceClient(ctrl)
		mockWishListRepo := wishlists_mock.NewMockRepository(ctrl)

		mockHotelRepo.EXPECT().
			GetHotelsPreview("Villa").
			Return(testHotels, nil)

		u := NewHotelUsecase(mockHotelRepo, mockUserService, mockWishListRepo)

		hotels, err := u.GetHotelsPreview("Villa")

		assert.NoError(t, err)
		assert.Equal(t, hotels, testHotels)
	})
	t.Run("HotelGetHotelsErr", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockHotelRepo := hotels_mock.NewMockRepository(ctrl)
		mockUserService := userService.NewMockUserServiceClient(ctrl)
		mockWishListRepo := wishlists_mock.NewMockRepository(ctrl)

		mockHotelRepo.EXPECT().
			GetHotelsPreview("Villa").
			Return(testHotels, customerror.NewCustomError(errors.New(""), serverError.ServerInternalError, 1))

		u := NewHotelUsecase(mockHotelRepo, mockUserService, mockWishListRepo)

		_, err := u.GetHotelsPreview("Villa")

		assert.Error(t, err)
		assert.Equal(t, customerror.ParseCode(err), serverError.ServerInternalError)
	})
}

func TestHotelUseCase_FetchHotels(t *testing.T) {

	testHotels := []hotelmodel.Hotel{
		{HotelID: 3, Name: "plaza", Location: "street 4 Moscow, Russia", Description: "very good hotel"},
		{HotelID: 4, Name: "plaza", Location: "street 4 Moscow, Russia", Description: "very good hotel"},
	}
	paginfo := paginationModel.PaginationInfo{ItemsCount: 2, NextLink: "api/v1/hotels/?id=3&limit=1&offset=3",
		PrevLink: "api/v1/hotels/?id=3&limit=1&offset=1"}

	searchTestData := hotelmodel.SearchData{Hotels: testHotels, PagInfo: paginfo}
	t.Run("HotelFetchHotels", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockHotelRepo := hotels_mock.NewMockRepository(ctrl)
		mockUserService := userService.NewMockUserServiceClient(ctrl)
		mockWishListRepo := wishlists_mock.NewMockRepository(ctrl)

		filter := hotelmodel.HotelFiltering{}
		mockHotelRepo.EXPECT().
			FetchHotels(filter, "Villa", 0).
			Return(testHotels, nil)

		mockWishListRepo.EXPECT().
			CheckHotelInWishlists(2, 3).
			Return("", nil)
		mockWishListRepo.EXPECT().
			CheckHotelInWishlists(2, 4).
			Return("", nil)

		u := NewHotelUsecase(mockHotelRepo, mockUserService, mockWishListRepo)

		hotels, err := u.FetchHotels(filter, "Villa", 0, 2)
		hotels.PagInfo = paginfo

		assert.NoError(t, err)
		assert.Equal(t, hotels, searchTestData)
	})
	t.Run("HotelFetchHotels", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mockHotelRepo := hotels_mock.NewMockRepository(ctrl)
		mockUserService := userService.NewMockUserServiceClient(ctrl)
		mockWishListRepo := wishlists_mock.NewMockRepository(ctrl)

		filter := hotelmodel.HotelFiltering{}
		mockHotelRepo.EXPECT().
			FetchHotels(filter, "Villa", 0).
			Return(testHotels, nil)

		mockWishListRepo.EXPECT().
			CheckHotelInWishlists(2, 3).
			Return("", nil)
		mockWishListRepo.EXPECT().
			CheckHotelInWishlists(2, 4).
			Return("", nil)

		u := NewHotelUsecase(mockHotelRepo, mockUserService, mockWishListRepo)

		hotels, err := u.FetchHotels(filter, "Villa", 0, 2)
		hotels.PagInfo = paginfo

		assert.NoError(t, err)
		assert.Equal(t, searchTestData, hotels)
	})
	t.Run("HotelFetchHotelsErr2", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockHotelRepo := hotels_mock.NewMockRepository(ctrl)
		mockUserService := userService.NewMockUserServiceClient(ctrl)
		mockWishListRepo := wishlists_mock.NewMockRepository(ctrl)

		filter := hotelmodel.HotelFiltering{}
		mockHotelRepo.EXPECT().
			FetchHotels(filter, "Villa", 0).
			Return(testHotels, customerror.NewCustomError(errors.New(""), serverError.ServerInternalError, 1))

		u := NewHotelUsecase(mockHotelRepo, mockUserService, mockWishListRepo)

		_, err := u.FetchHotels(filter, "Villa", 0, 2)

		assert.Error(t, err)
		assert.Equal(t, customerror.ParseCode(err), serverError.ServerInternalError)
	})
}

func TestHotelUseCase_CheckRateExist(t *testing.T) {
	comment := commModel.FullCommentInfo{
		UserID: 2, CommID: 3, HotelID: 4, Message: "kekw", Rating: 3, Username: "kostikan", Avatar: "kek.jpg",
	}
	t.Run("CheckRateExist", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockHotelRepo := hotels_mock.NewMockRepository(ctrl)
		mockUserService := userService.NewMockUserServiceClient(ctrl)
		mockWishListRepo := wishlists_mock.NewMockRepository(ctrl)

		mockHotelRepo.EXPECT().
			CheckRateExist(2, 4).
			Return(comment, nil)

		mockUserService.
			EXPECT().
			GetUserByID(gomock.Any(), &userService.UserID{UserID: 2}).
			Return(&userService.User{UserID: 2, Username: "kostikan", Avatar: "kek.jpg"}, nil)

		u := NewHotelUsecase(mockHotelRepo, mockUserService, mockWishListRepo)

		resComment, err := u.CheckRateExist(2, 4)

		assert.NoError(t, err)
		assert.Equal(t, comment, resComment)
	})
	t.Run("CheckRateExistErr", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockHotelRepo := hotels_mock.NewMockRepository(ctrl)
		mockUserService := userService.NewMockUserServiceClient(ctrl)
		mockWishListRepo := wishlists_mock.NewMockRepository(ctrl)

		mockHotelRepo.EXPECT().
			CheckRateExist(2, 4).
			Return(comment, customerror.NewCustomError(errors.New(""), serverError.ServerInternalError, 1))

		u := NewHotelUsecase(mockHotelRepo, mockUserService, mockWishListRepo)

		_, err := u.CheckRateExist(2, 4)

		assert.Error(t, err)
		assert.Equal(t, customerror.ParseCode(err), serverError.ServerInternalError)
	})
}
