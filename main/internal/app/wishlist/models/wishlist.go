//go:generate easyjson -all wishlist.go
package wishlistmodel

// easyjson:json
type Wishlist struct {
	WishlistID int    `json:"wishlist_id" db:"wishlist_id" mapstructure:"wishlist_id"`
	Name       string `json:"name" db:"name"`
	UserID     int    `json:"user_id" db:"user_id" mapstructure:"user_id"`
}

// easyjson:json
type UserWishLists struct {
	Wishlists []Wishlist `json:"wishlists,omitempty"`
}

// easyjson:json
type WishlistHotel struct {
	WishlistID int `json:"wishlist_id" db:"wishlist_id"`
	HotelID    int `json:"hotel_id" db:"hotel_id"`
}

// easyjson:json
type HotelWishlistRequest struct {
	HotelID int `json:"hotel_id"`
}
