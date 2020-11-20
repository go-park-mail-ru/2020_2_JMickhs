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

type WishlistByIdRequest struct {
	ID int `json:"id"`
}

type AddHotelToWishlistRequest struct {
	HotelID    int `json:"hotel_id"`
	WishlistID int `json:"wishlist_id"`
}

type DeleteHotelFromWishlistRequest struct {
	HotelID    int `json:"hotel_id"`
	WishlistID int `json:"wishlist_id"`
}

type DeleteWishlistRequest struct {
	ID int `json:"id"`
}
