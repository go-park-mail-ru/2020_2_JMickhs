package sqlrequests

const SearchHotelsPostgreRequest = "WHERE (name % $1 or location % $1 or name LIKE '%' || $1 || '%' or location LIKE '%' || $1 || '%')"

const GetCommentsCountPostgreRequest = "SELECT comm_count FROM hotels WHERE hotel_id = $1"

const GetCommentsPostgreRequest = "SELECT us.user_id,comm_id,message,rating,avatar,username,comm.hotel_id,time FROM comments as comm " +
	"INNER JOIN users as us on us.user_id = comm.user_id WHERE comm.hotel_id = $3  LIMIT $2 OFFSET $1"

const AddCommentsPostgreRequest = "INSERT INTO comments VALUES (default, $1, $2,$3,$4) RETURNING comm_id,time"

const DeleteCommentsPostgreRequest = "DELETE FROM comments WHERE comm_id=$1"

const UpdateCommentsPostgreRequest = "UPDATE comments SET message=$2,rating=$3 WHERE comm_id=$1 RETURNING time"

const GetHotelsPostgreRequest = "SELECT hotel_id,name,description,img,location,curr_rating FROM hotels LIMIT 4 OFFSET $1"

const GetHotelByIDPostgreRequest = "SELECT hotel_id,name,description,img,location,curr_rating,comm_count FROM hotels WHERE hotel_id=$1"

const GetHotelsPhotosPostgreRequest = "SELECT unnest(photos) FROM hotels WHERE hotel_id=$1"

const UpdateHotelRatingPostgreRequest = "UPDATE hotels SET curr_rating=$1 where hotel_id = $2"

const GetPrevRatingOnCommentPostgreRequest = "SELECT rating,user_id,hotel_id FROM comments where comm_id = $1"

const GetCurrRatingPostgreRequest = "SELECT round( CAST (curr_rating as numeric),1),comm_count FROM hotels where hotel_id = $1"

const AddUserPostgreRequest = "INSERT INTO users VALUES (default, $1, $2,$3,$4) RETURNING user_id"

const GetUserByNamePostgreRequest = "SELECT user_id,username,email,password,avatar FROM users WHERE username=$1"

const GetUserByIDPostgreRequest = "SELECT user_id,username,email,password,avatar FROM users WHERE user_id=$1"

const UpdateUserCredPostgreRequest = "UPDATE users SET username=$2,email=$3 WHERE user_id=$1"

const UpdateUserAvatarPostgreRequest = "UPDATE users SET avatar=$2 WHERE user_id=$1"

const UpdateUserPasswordPostgreRequest = "UPDATE users SET  password=$2 WHERE user_id=$1"

const CheckRateIfExistPostgreRequest = "SELECT rating FROM comments WHERE user_id=$1 AND hotel_id=$2"

const GetWishlistPostgreRequest = "SELECT h.hotel_id,h.name,h.description,h.img,h.location,h.curr_rating FROM hotels AS h LEFT JOIN wishlistshotels AS wh ON h.hotel_id = wh.hotel_id LEFT JOIN wishlists AS w ON w.wishlist_id = wh.wishlist_id WHERE wh.wishlist_id = $1"

const CreateWishlistPostgreRequest = "INSERT INTO wishlists (wishlist_id, name, user_id) VALUES ($1, $2, $3)"

const DeleteWishlistPostgreRequest = "DELETE FROM wishlists WHERE wishlist_id = $1"

const AddHotelToWishlistPostgreRequest = "INSERT INTO wishlistshotels (wishlist_id, hotel_id) VALUES ($1, $2)"

const DeleteHotelFromWishlistPostgreRequest = "DELETE FROM wishlistshotels WHERE wishlist_id = $1 AND hotel_id = $2"
