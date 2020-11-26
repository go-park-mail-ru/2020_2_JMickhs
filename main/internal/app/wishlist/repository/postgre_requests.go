package wishlistrepository

const GetMiniHotelPostgreRequest = "SELECT h.hotel_id,h.name,h.description,h.img,h.location,h.curr_rating FROM hotels AS h WHERE hotel_id = $1"

const GetWishlistMeta = "SELECT wishlist_id, hotel_id from wishlistshotels where wishlist_id =$1"

const CreateWishlistPostgreRequest = "INSERT INTO wishlists (wishlist_id, name, user_id) VALUES ($1, $2, $3)"

const DeleteWishlistPostgreRequest = "DELETE FROM wishlists WHERE wishlist_id = $1"

const AddHotelToWishlistPostgreRequest = "INSERT INTO wishlistshotels (wishlist_id, hotel_id) VALUES ($1, $2)"

const DeleteHotelFromWishlistPostgreRequest = "DELETE FROM wishlistshotels WHERE wishlist_id = $1 AND hotel_id = $2"
