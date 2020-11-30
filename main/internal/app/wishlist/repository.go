//go:generate mockgen -source repository.go -destination mocks/wishlists_repository_mock.go -package wishlists_mock
package wishlist

import (
	wishlistModel "github.com/go-park-mail-ru/2020_2_JMickhs/main/internal/app/wishlist/models"
)

type Repository interface {
	CheckWishListOwner(wishListID int, UserID int) (bool, error)
	GetWishlistMeta(wishlistID int) ([]wishlistModel.WishlistHotel, error)
	CreateWishlist(wishlist wishlistModel.Wishlist) (wishlistModel.Wishlist, error)
	DeleteWishlist(wishlistID int) error
	AddHotel(hotelID int, wishlistID int) error
	DeleteHotel(hotelID int, wishlistID int) error
	GetUserWishlists(userID int) (wishlistModel.UserWishLists, error)
	CheckHotelInWishlists(userID int, hotelID int) (string, error)
}
