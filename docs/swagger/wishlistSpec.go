package swagger

// easyjson:json
type MiniHotel struct {
	HotelID  int     `json:"hotel_id" db:"hotel_id"`
	Name     string  `json:"name" db:"name"`
	Image    string  `json:"image" db:"img"`
	Location string  `json:"location" db:"location"`
	Rating   float64 `json:"rating" db:"curr_rating"`
}

type Wishlist struct {
	WishlistID int    `json:"wishlist_id" db:"wishlist_id"`
	Name       string `json:"name" db:"name"`
	UserID     int    `json:"user_id" db:"user_id"`
}

type UserWishLists struct {
	Wishlists []Wishlist `json:"wishlists,omitempty"`
}

type CreateWishlistRequest struct {
	Name string `json:"name"`
}

type AddHotelToWishlistRequest struct {
	HotelID int `json:"hotel_id"`
}

type DeleteHotelFromWishlistRequest struct {
	HotelID int `json:"hotel_id"`
}

// swagger:parameters addHotelToWishlist
type addHotelToWishlist struct {
	// the id of wishlist
	// in: path
	// required:true
	ID int `json:"wishlist_id"`
	//in: body
	Body AddHotelToWishlistRequest
}

// swagger:response wishlist
type SaveWIshlistResponse struct {
	//in:body
	Body Wishlist
}

// swagger:response wishlists
type GetUserWishListsResponse struct {
	//in:body
	Body UserWishLists
}

// swagger:response wishlisthotels
type GetWishListsHotelResponse struct {
	//in:body
	Body []MiniHotel
}

// // swagger:response
// type SaveHotelToWishlistResponse struct {
// 	//in:body
// 	Body wishlistModel.Wishlist
// }

// swagger:parameters createWishlist
type createWishlistRequest struct {
	//in: body
	Body CreateWishlistRequest
}

// swagger:parameters deleteWishlist
type deleteWishlistRequest struct {
	// the id of wishlist
	// in: path
	// required:true
	ID int `json:"wishlist_id"`
}

// swagger:parameters getWishlist
type getWishlistRequest struct {
	// the id of wishlist
	// in: path
	// required:true
	ID int `json:"wishlist_id"`
}

// swagger:parameters deleteHotelFromWishlist
type deleteHotelFromWishlist struct {
	// the id of wishlist
	// in: path
	// required:true
	ID int `json:"wishlist_id"`
	//in: body
	Body DeleteHotelFromWishlistRequest
}
