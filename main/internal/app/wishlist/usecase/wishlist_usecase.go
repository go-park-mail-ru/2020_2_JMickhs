package wishlistusecase

import (
	"errors"

	"github.com/go-park-mail-ru/2020_2_JMickhs/main/internal/app/hotels"

	wishlistModel "github.com/go-park-mail-ru/2020_2_JMickhs/main/internal/app/wishlist/models"
	"github.com/go-park-mail-ru/2020_2_JMickhs/package/clientError"
	customerror "github.com/go-park-mail-ru/2020_2_JMickhs/package/error"

	"github.com/go-park-mail-ru/2020_2_JMickhs/main/internal/app/wishlist"
)

type WishlistUseCase struct {
	wishlistRepo wishlist.Repository
	hotelsRepo   hotels.Repository
}

func NewWishlistUseCase(r wishlist.Repository, hotelRepo hotels.Repository) *WishlistUseCase {
	return &WishlistUseCase{
		wishlistRepo: r,
		hotelsRepo:   hotelRepo,
	}
}

func (w *WishlistUseCase) WishListsByHotel(userID int, hotelID int) (wishlistModel.UserWishLists, error) {
	return w.wishlistRepo.WishListsByHotel(userID, hotelID)
}

func (w *WishlistUseCase) GetWishlistMeta(userID int, wishlistID int) ([]wishlistModel.WishlistHotel, error) {
	var wishListMeta []wishlistModel.WishlistHotel
	check, err := w.wishlistRepo.CheckWishListOwner(wishlistID, userID)
	if err != nil {
		return wishListMeta, err
	}
	if !check {
		return wishListMeta, customerror.NewCustomError(errors.New("not the owner of wishlist"), clientError.Locked, 1)
	}
	return w.wishlistRepo.GetWishlistMeta(wishlistID)
}

func (w *WishlistUseCase) CreateWishlist(wishlist wishlistModel.Wishlist) (wishlistModel.Wishlist, error) {
	return w.wishlistRepo.CreateWishlist(wishlist)
}

func (w *WishlistUseCase) DeleteWishlist(userID int, wishlistID int) error {
	check, err := w.wishlistRepo.CheckWishListOwner(wishlistID, userID)
	if err != nil {
		return err
	}
	if !check {
		return customerror.NewCustomError(errors.New("not the owner of wishlist"), clientError.Locked, 1)
	}
	return w.wishlistRepo.DeleteWishlist(wishlistID)
}

func (w *WishlistUseCase) AddHotel(userID int, hotelID int, wishlistID int) error {
	check, err := w.wishlistRepo.CheckWishListOwner(wishlistID, userID)
	if err != nil {
		return err
	}
	if !check {
		return customerror.NewCustomError(errors.New("not the owner of wishlist"), clientError.Locked, 1)
	}
	_, err = w.hotelsRepo.GetMiniHotelByID(hotelID)
	if err != nil {
		return err
	}
	return w.wishlistRepo.AddHotel(hotelID, wishlistID)
}

func (w *WishlistUseCase) DeleteHotel(userID int, hotelID int, wishlistID int) error {
	check, err := w.wishlistRepo.CheckWishListOwner(wishlistID, userID)
	if err != nil {
		return err
	}
	if !check {
		return customerror.NewCustomError(errors.New("not the owner of wishlist"), clientError.Locked, 1)
	}
	return w.wishlistRepo.DeleteHotel(hotelID, wishlistID)
}

func (w *WishlistUseCase) GetUserWishlists(userID int) (wishlistModel.UserWishLists, error) {
	return w.wishlistRepo.GetUserWishlists(userID)
}

func (w *WishlistUseCase) CheckHotelInWishlists(userID int, hotelID int) (string, error) {
	return w.wishlistRepo.CheckHotelInWishlists(userID, hotelID)
}
