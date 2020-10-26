package wishlistusecase

import (
	hotelmodel "github.com/go-park-mail-ru/2020_2_JMickhs/internal/app/hotels/models"
	wishlistModel "github.com/go-park-mail-ru/2020_2_JMickhs/internal/app/wishlist/models"

	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/app/wishlist"
)

type WishlistUseCase struct {
	wishlistRepo wishlist.Repository
}

func NewWishlistUseCase(r wishlist.Repository) *WishlistUseCase {
	return &WishlistUseCase{
		wishlistRepo: r,
	}
}
func (w *WishlistUseCase) GetWishlist(wishlistID int) ([]hotelmodel.Hotel, error) {
	return w.wishlistRepo.GetWishlist(wishlistID)
}

func (w *WishlistUseCase) CreateWishlist(wishlist wishlistModel.Wishlist) error {
	return w.wishlistRepo.CreateWishlist(wishlist)
}

func (w *WishlistUseCase) DeleteWishlist(wishlistID int) error {
	return w.wishlistRepo.DeleteWishlist(wishlistID)
}

func (w *WishlistUseCase) AddHotel(hotelID int, wishlistID int) error {
	return w.wishlistRepo.AddHotel(hotelID, wishlistID)
}

func (w *WishlistUseCase) DeleteHotel(hotelID string, wishlistID int) error {
	return w.wishlistRepo.DeleteHotel(hotelID, wishlistID)
}
