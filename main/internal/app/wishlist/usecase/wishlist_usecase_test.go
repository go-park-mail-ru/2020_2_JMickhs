package wishlistusecase

import (
	"errors"
	"testing"

	hotelmodel "github.com/go-park-mail-ru/2020_2_JMickhs/main/internal/app/hotels/models"

	hotels_mock "github.com/go-park-mail-ru/2020_2_JMickhs/main/internal/app/hotels/mocks"

	"github.com/go-park-mail-ru/2020_2_JMickhs/package/clientError"
	customerror "github.com/go-park-mail-ru/2020_2_JMickhs/package/error"

	wishlists_mock "github.com/go-park-mail-ru/2020_2_JMickhs/main/internal/app/wishlist/mocks"
	wishlistmodel "github.com/go-park-mail-ru/2020_2_JMickhs/main/internal/app/wishlist/models"

	"github.com/stretchr/testify/assert"

	"github.com/golang/mock/gomock"
)

func TestWishlistUseCase_GetWishlistMeta(t *testing.T) {
	wishlist := []wishlistmodel.WishlistHotel{
		{WishlistID: 1, HotelID: 2},
		{WishlistID: 1, HotelID: 3},
	}
	t.Run("GetWishlistMeta", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockWishlistRepository := wishlists_mock.NewMockRepository(ctrl)
		mockHotelsRepository := hotels_mock.NewMockRepository(ctrl)
		mockWishlistRepository.EXPECT().
			GetWishlistMeta(1).
			Return(wishlist, nil)

		mockWishlistRepository.EXPECT().
			CheckWishListOwner(1, 3).
			Return(true, nil)

		wishlistUsecase := NewWishlistUseCase(mockWishlistRepository, mockHotelsRepository)

		wishlistMeta, err := wishlistUsecase.GetWishlistMeta(3, 1)
		assert.NoError(t, err)
		assert.Equal(t, wishlist, wishlistMeta)
	})

	t.Run("GetWishlistMetaErr", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockWishlistRepository := wishlists_mock.NewMockRepository(ctrl)
		mockHotelsRepository := hotels_mock.NewMockRepository(ctrl)

		mockWishlistRepository.EXPECT().
			CheckWishListOwner(1, 3).
			Return(false, nil)

		wishlistUsecase := NewWishlistUseCase(mockWishlistRepository, mockHotelsRepository)

		_, err := wishlistUsecase.GetWishlistMeta(3, 1)
		assert.Error(t, err)
		assert.Equal(t, customerror.ParseCode(err), clientError.Locked)
	})

	t.Run("GetWishlistMetaErr", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockWishlistRepository := wishlists_mock.NewMockRepository(ctrl)
		mockHotelsRepository := hotels_mock.NewMockRepository(ctrl)

		mockWishlistRepository.EXPECT().
			CheckWishListOwner(1, 3).
			Return(true, nil)

		mockWishlistRepository.EXPECT().
			GetWishlistMeta(1).
			Return(wishlist, customerror.NewCustomError(errors.New(""), clientError.BadRequest, 1))

		wishlistUsecase := NewWishlistUseCase(mockWishlistRepository, mockHotelsRepository)

		_, err := wishlistUsecase.GetWishlistMeta(3, 1)
		assert.Error(t, err)
		assert.Equal(t, customerror.ParseCode(err), clientError.BadRequest)
	})
}

func TestWishlistUseCase_DeleteWishlist(t *testing.T) {
	t.Run("DeleteWishlist", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockWishlistRepository := wishlists_mock.NewMockRepository(ctrl)
		mockHotelsRepository := hotels_mock.NewMockRepository(ctrl)

		mockWishlistRepository.EXPECT().
			DeleteWishlist(1).
			Return(nil)

		mockWishlistRepository.EXPECT().
			CheckWishListOwner(1, 3).
			Return(true, nil)

		wishlistUsecase := NewWishlistUseCase(mockWishlistRepository, mockHotelsRepository)

		err := wishlistUsecase.DeleteWishlist(3, 1)
		assert.NoError(t, err)
	})

	t.Run("DeleteWishlistErr", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockWishlistRepository := wishlists_mock.NewMockRepository(ctrl)
		mockHotelsRepository := hotels_mock.NewMockRepository(ctrl)

		mockWishlistRepository.EXPECT().
			CheckWishListOwner(1, 3).
			Return(false, nil)

		wishlistUsecase := NewWishlistUseCase(mockWishlistRepository, mockHotelsRepository)

		err := wishlistUsecase.DeleteWishlist(3, 1)
		assert.Error(t, err)
		assert.Equal(t, customerror.ParseCode(err), clientError.Locked)
	})

	t.Run("DeleteWishlistErr1", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockWishlistRepository := wishlists_mock.NewMockRepository(ctrl)
		mockHotelsRepository := hotels_mock.NewMockRepository(ctrl)

		mockWishlistRepository.EXPECT().
			CheckWishListOwner(1, 3).
			Return(true, nil)

		mockWishlistRepository.EXPECT().
			DeleteWishlist(1).
			Return(customerror.NewCustomError(errors.New(""), clientError.BadRequest, 1))

		wishlistUsecase := NewWishlistUseCase(mockWishlistRepository, mockHotelsRepository)

		err := wishlistUsecase.DeleteWishlist(3, 1)
		assert.Error(t, err)
		assert.Equal(t, customerror.ParseCode(err), clientError.BadRequest)
	})
}
func TestWishlistUseCase_DeleteHotel(t *testing.T) {
	t.Run("DeleteHotel", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockWishlistRepository := wishlists_mock.NewMockRepository(ctrl)
		mockHotelsRepository := hotels_mock.NewMockRepository(ctrl)

		mockWishlistRepository.EXPECT().
			DeleteHotel(3, 1).
			Return(nil)

		mockWishlistRepository.EXPECT().
			CheckWishListOwner(1, 3).
			Return(true, nil)

		wishlistUsecase := NewWishlistUseCase(mockWishlistRepository, mockHotelsRepository)

		err := wishlistUsecase.DeleteHotel(3, 3, 1)
		assert.NoError(t, err)
	})

	t.Run("DeleteHotelErr", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockWishlistRepository := wishlists_mock.NewMockRepository(ctrl)
		mockHotelsRepository := hotels_mock.NewMockRepository(ctrl)

		mockWishlistRepository.EXPECT().
			CheckWishListOwner(1, 3).
			Return(false, nil)

		wishlistUsecase := NewWishlistUseCase(mockWishlistRepository, mockHotelsRepository)

		err := wishlistUsecase.DeleteHotel(3, 3, 1)
		assert.Error(t, err)
		assert.Equal(t, customerror.ParseCode(err), clientError.Locked)
	})

	t.Run("DeleteHotelErr2", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockWishlistRepository := wishlists_mock.NewMockRepository(ctrl)
		mockHotelsRepository := hotels_mock.NewMockRepository(ctrl)

		mockWishlistRepository.EXPECT().
			CheckWishListOwner(1, 3).
			Return(true, nil)

		mockWishlistRepository.EXPECT().
			DeleteHotel(3, 1).
			Return(customerror.NewCustomError(errors.New(""), clientError.BadRequest, 1))

		wishlistUsecase := NewWishlistUseCase(mockWishlistRepository, mockHotelsRepository)

		err := wishlistUsecase.DeleteHotel(3, 3, 1)
		assert.Error(t, err)
		assert.Equal(t, customerror.ParseCode(err), clientError.BadRequest)
	})
}
func TestWishlistUseCase_AddHotel(t *testing.T) {
	t.Run("AddHotel", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockWishlistRepository := wishlists_mock.NewMockRepository(ctrl)
		mockHotelsRepository := hotels_mock.NewMockRepository(ctrl)

		mockWishlistRepository.EXPECT().
			AddHotel(1, 1).
			Return(nil)

		mockHotelsRepository.EXPECT().
			GetMiniHotelByID(1).
			Return(hotelmodel.MiniHotel{}, nil)
		mockWishlistRepository.EXPECT().
			CheckWishListOwner(1, 3).
			Return(true, nil)

		wishlistUsecase := NewWishlistUseCase(mockWishlistRepository, mockHotelsRepository)

		err := wishlistUsecase.AddHotel(3, 1, 1)
		assert.NoError(t, err)
	})

	t.Run("AddHotelErr1", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockWishlistRepository := wishlists_mock.NewMockRepository(ctrl)
		mockHotelsRepository := hotels_mock.NewMockRepository(ctrl)

		mockWishlistRepository.EXPECT().
			CheckWishListOwner(1, 3).
			Return(false, nil)

		wishlistUsecase := NewWishlistUseCase(mockWishlistRepository, mockHotelsRepository)

		err := wishlistUsecase.AddHotel(3, 1, 1)
		assert.Error(t, err)
		assert.Equal(t, customerror.ParseCode(err), clientError.Locked)
	})

	t.Run("AddHotelErr2", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockWishlistRepository := wishlists_mock.NewMockRepository(ctrl)
		mockHotelsRepository := hotels_mock.NewMockRepository(ctrl)

		mockWishlistRepository.EXPECT().
			CheckWishListOwner(1, 3).
			Return(true, nil)
		mockHotelsRepository.EXPECT().
			GetMiniHotelByID(1).
			Return(hotelmodel.MiniHotel{}, nil)

		mockWishlistRepository.EXPECT().
			AddHotel(1, 1).
			Return(customerror.NewCustomError(errors.New(""), clientError.BadRequest, 1))

		wishlistUsecase := NewWishlistUseCase(mockWishlistRepository, mockHotelsRepository)

		err := wishlistUsecase.AddHotel(3, 1, 1)
		assert.Error(t, err)
		assert.Equal(t, customerror.ParseCode(err), clientError.BadRequest)
	})
}

func TestWishlistUseCase_CheckHotelInWishlists(t *testing.T) {
	t.Run("CheckHotelInWishlists", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockWishlistRepository := wishlists_mock.NewMockRepository(ctrl)
		mockHotelsRepository := hotels_mock.NewMockRepository(ctrl)

		mockWishlistRepository.EXPECT().
			CheckHotelInWishlists(2, 1).
			Return("In", nil)

		wishlistUsecase := NewWishlistUseCase(mockWishlistRepository, mockHotelsRepository)

		res, err := wishlistUsecase.CheckHotelInWishlists(2, 1)
		assert.NoError(t, err)
		assert.Equal(t, res, "In")
	})
}

func TestWishlistUseCase_GetUserWishlists(t *testing.T) {
	wishList := wishlistmodel.UserWishLists{Wishlists: []wishlistmodel.Wishlist{
		{WishlistID: 2, Name: "kekw", UserID: 3},
		{WishlistID: 3, Name: "kekw2", UserID: 3},
	}}
	t.Run("GetUserWishlists", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockWishlistRepository := wishlists_mock.NewMockRepository(ctrl)
		mockHotelsRepository := hotels_mock.NewMockRepository(ctrl)

		mockWishlistRepository.EXPECT().
			GetUserWishlists(2).
			Return(wishList, nil)

		wishlistUsecase := NewWishlistUseCase(mockWishlistRepository, mockHotelsRepository)

		res, err := wishlistUsecase.GetUserWishlists(2)
		assert.NoError(t, err)
		assert.Equal(t, wishList, res)
	})
}

func TestWishlistUseCase_CreateWishlist(t *testing.T) {
	wishList := wishlistmodel.Wishlist{WishlistID: 2, Name: "kekw", UserID: 3}
	t.Run("GetUserWishlists", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockWishlistRepository := wishlists_mock.NewMockRepository(ctrl)
		mockHotelsRepository := hotels_mock.NewMockRepository(ctrl)

		mockWishlistRepository.EXPECT().
			CreateWishlist(wishList).
			Return(wishList, nil)

		wishlistUsecase := NewWishlistUseCase(mockWishlistRepository, mockHotelsRepository)

		res, err := wishlistUsecase.CreateWishlist(wishList)
		assert.NoError(t, err)
		assert.Equal(t, wishList, res)
	})
}
