package wishlist

import (
	wishlistModel "github.com/go-park-mail-ru/2020_2_JMickhs/main/internal/app/wishlist/models"
)

type Usecase interface {
	GetWishlistMeta(wishlistID int) ([]wishlistModel.WishlisstHotel, error)
	CreateWishlist(wishlist wishlistModel.Wishlist) error
	DeleteWishlist(wishlistID int) error
	AddHotel(hotelID int, wishlistID int) error
	DeleteHotel(hotelID int, wishlistID int) error
}
