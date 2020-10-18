package sqlrequests

const GetCommentsPostgreRequest = "SELECT us.user_id,comm_id,message,rt.rate,avatar,username,comm.hotel_id,time FROM comments as comm " +
	"INNER JOIN users as us on us.user_id = comm.user_id " +
	"INNER JOIN rating as rt on rt.user_id = us.user_id AND rt.hotel_id = $1 WHERE comm.hotel_id = $1  LIMIT 4 OFFSET $2"

const AddCommentsPostgreRequest = "INSERT INTO comments VALUES (default, $1, $2,$3) RETURNING comm_id,time"

const DeleteCommentsPostgreRequest = "DELETE FROM comments WHERE comm_id=$1"

const UpdateCommentsPostgreRequest = "UPDATE comments SET message=$2 WHERE comm_id=$1"

const GetHotelsPostgreRequest = "SELECT hotel_id,name,description,img,location,curr_rating FROM hotels LIMIT 4 OFFSET $1"

const GetHotelByIDPostgreRequest = "SELECT hotel_id,name,description,img,location,curr_rating,photos[:] FROM hotels WHERE hotel_id=$1"

const InsertRatingPostgreRequest = "INSERT INTO rating VALUES (default, $1, $2, $3)"

const UpdateHotelRatingPostgreRequest = "UPDATE hotels SET curr_rating=$1 where hotel_id = $2"

const GetRatingCountOnHotelPostgreRequest = "SELECT COUNT(*) FROM rating where hotel_id = $1"

const GetCurrRatingPostgreRequest = "SELECT curr_rating FROM hotels where hotels.hotel_id = $1"

const AddUserPostgreRequest = "INSERT INTO users VALUES (default, $1, $2,$3,$4) RETURNING user_id"

const GetUserByNamePostgreRequest = "SELECT user_id,username,email,password,avatar FROM users WHERE username=$1"

const GetUserByIDPostgreRequest = "SELECT user_id,username,email,password,avatar FROM users WHERE user_id=$1"

const UpdateUserCredPostgreRequest = "UPDATE users SET username=$2,email=$3 WHERE user_id=$1"

const UpdateUserAvatarPostgreRequest = "UPDATE users SET avatar=$2 WHERE user_id=$1"

const UpdateUserPasswordPostgreRequest = "UPDATE users SET  password=$2 WHERE user_id=$1"

const CheckRateIfExistPostgreRequest = "SELECT rate FROM rating WHERE user_id=$1 AND hotel_id=$2"
