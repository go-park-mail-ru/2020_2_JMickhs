package swagger

type CreateWishlistRequest struct {
	WishlistID int    `json:"wishlist_id"`
	Name       string `json:"name"`
	UserID     int    `json:"user_id"`
}

type GetWishlistRequest struct {
	ID int `json:"id"`
}

type DeleteWishlistRequest struct {
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

// swagger:parameters addHotelToWishlist
type addHotelToWishlist struct {
	//in: body
	Body AddHotelToWishlistRequest
}

// swagger:parameters getWishlist
type getWishlistRequest struct {
	//in: body
	Body GetWishlistRequest
}

// swagger:response saveWishlist
type SaveWIshlistResponse struct {
	//in:body
	Body wishlistModel.Wishlist
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
	//in: body
	Body DeleteWishlistRequest
}

// swagger:parameters deleteHotelFromWishlist
type deleteHotelFromWishlist struct {
	//in: body
	Body DeleteHotelFromWishlistRequest
}
