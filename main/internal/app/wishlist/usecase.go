//go:generate mockgen -source usecase.go -destination mocks/wishlists_usecase_mock.go -package wishlists_mock
package wishlist

import (
	wishlistModel "github.com/go-park-mail-ru/2020_2_JMickhs/main/internal/app/wishlist/models"
)

type Usecase interface {
	GetWishlistMeta(userID int, wishlistID int) ([]wishlistModel.WishlistHotel, error)
	CreateWishlist(wishlist wishlistModel.Wishlist) (wishlistModel.Wishlist, error)
	DeleteWishlist(userID int, wishlistID int) error
	AddHotel(userID int, hotelID int, wishlistID int) error
	DeleteHotel(userID int, hotelID int, wishlistID int) error
	GetUserWishlists(userID int) (wishlistModel.UserWishLists, error)
	CheckHotelInWishlists(userID int, hotelID int) (string, error)
}
