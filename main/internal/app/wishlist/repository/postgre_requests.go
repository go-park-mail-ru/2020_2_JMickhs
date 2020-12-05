package wishlistrepository

const GetWishlistMeta = "SELECT wishlist_id, hotel_id from wishlistshotels where wishlist_id = $1"

const CreateWishlistPostgreRequest = "INSERT INTO wishlists (wishlist_id, name, user_id) VALUES (default, $1, $2) RETURNING wishlist_id"

const DeleteWishlistPostgreRequest = "DELETE FROM wishlists WHERE wishlist_id = $1"

const AddHotelToWishlistPostgreRequest = "INSERT INTO wishlistshotels (wishlist_id, hotel_id) VALUES ($1, $2)"

const DeleteHotelFromWishlistPostgreRequest = "DELETE FROM wishlistshotels WHERE wishlist_id = $1 AND hotel_id = $2"

const CheckWishListOwnerPostgreRequest = "SELECT user_id from wishlists where wishlist_id = $1"

const GetUserWithListsPostgreRequest = "SELECT wishlist_id,name from wishlists where user_id=$1"

const GetWishlistsBYHotelPostgreRequest = "SELECT DISTINCT wl.wishlist_id, wl.name from wishlists as wl inner join wishlistshotels as wh on wl.wishlist_id = wh.wishlist_id " +
	" where wl.user_id = $1 and wh.hotel_id = $2"

const CheckHotelInWishlistsPostgreRequest = "SELECT DISTINCT hotel_id from wishlists as wl inner join wishlistshotels as wh on wl.wishlist_id = wh.wishlist_id" +
	" where wl.user_id = $1 and wh.hotel_id = $2"
