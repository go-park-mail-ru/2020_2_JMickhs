package wishlistusecase

import (
	"errors"

	wishlistModel "github.com/go-park-mail-ru/2020_2_JMickhs/main/internal/app/wishlist/models"
	"github.com/go-park-mail-ru/2020_2_JMickhs/package/clientError"
	customerror "github.com/go-park-mail-ru/2020_2_JMickhs/package/error"

	"github.com/go-park-mail-ru/2020_2_JMickhs/main/internal/app/wishlist"
)

type WishlistUseCase struct {
	wishlistRepo wishlist.Repository
}

func NewWishlistUseCase(r wishlist.Repository) *WishlistUseCase {
	return &WishlistUseCase{
		wishlistRepo: r,
	}
}

func (w *WishlistUseCase) GetWishlistMeta(userID int, wishlistID int) ([]wishlistModel.WishlistHotel, error) {
	wishListMeta := []wishlistModel.WishlistHotel{}
	check, err := w.wishlistRepo.CheckWishListOwner(wishlistID, userID)
	if err != nil {
		return wishListMeta, err
	}
	if check == false {
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
	if check == false {
		return customerror.NewCustomError(errors.New("not the owner of wishlist"), clientError.Locked, 1)
	}
	return w.wishlistRepo.DeleteWishlist(wishlistID)
}

func (w *WishlistUseCase) AddHotel(userID int, hotelID int, wishlistID int) error {
	check, err := w.wishlistRepo.CheckWishListOwner(wishlistID, userID)
	if err != nil {
		return err
	}
	if check == false {
		return customerror.NewCustomError(errors.New("not the owner of wishlist"), clientError.Locked, 1)
	}
	return w.wishlistRepo.AddHotel(hotelID, wishlistID)
}

func (w *WishlistUseCase) DeleteHotel(userID int, hotelID int, wishlistID int) error {
	check, err := w.wishlistRepo.CheckWishListOwner(wishlistID, userID)
	if err != nil {
		return err
	}
	if check == false {
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
