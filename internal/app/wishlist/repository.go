package wishlist

import (
	hotelmodel "github.com/go-park-mail-ru/2020_2_JMickhs/internal/app/hotels/models"
	wishlistModel "github.com/go-park-mail-ru/2020_2_JMickhs/internal/app/wishlist/models"
)

type Repository interface {
	GetWishlist(wishlistID int) ([]hotelmodel.MiniHotel, error)
	CreateWishlist(wishlist wishlistModel.Wishlist) error
	DeleteWishlist(wishlistID int) error
	AddHotel(hotelID int, wishlistID int) error
	DeleteHotel(hotelID int, wishlistID int) error
	GetTable() ([]wishlistModel.WishlisstHotel, error)
	// GetWhatYouWant() ([]hotelmodel.MiniHotel, error)
}
