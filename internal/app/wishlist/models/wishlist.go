package wishlistmodel

type Wishlist struct {
	WishistID string `json:"wishlist_id" db:"wishlist_id"`
	Name      string `json:"name" db:"name"`
	UserID    string `json:"user_id" db:"user_id"`
}

type WishlisstHotel struct {
	WishistID string `json:"wishlist_id" db:"wishlist_id"`
	HotelID   int    `json:"hotel_id" db:"hotel_id"`
}
