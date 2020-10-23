package wishlist

import (
	wishlistModel "github.com/go-park-mail-ru/2020_2_JMickhs/internal/app/wishlist/models"
)

type Repository interface {
	GetWhishlist(wishlistID int) (wishlistModel.Wishlist, error)
	CreateWishlist(wishlist wishlistModel.Wishlist) error
	DeleteWishlist(wishlistID int) error
	UpdateWishlist(wishlist wishlistModel.Wishlist) error
	AddHotel(hotelID int, wishlistID int) error
	DeleteHotel(hotelID string, wishlistID int) error
}
