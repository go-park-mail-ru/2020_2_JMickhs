package commentRepository

const UpdateHotelRatingPostgreRequest = "UPDATE hotels SET curr_rating=$1 where hotel_id = $2"

const GetPrevRatingOnCommentPostgreRequest = "SELECT rating,user_id,hotel_id FROM comments where comm_id = $1"

const GetCurrRatingPostgreRequest = "SELECT round( CAST (curr_rating as numeric),1),comm_count FROM hotels where hotel_id = $1"

const GetCommentsCountPostgreRequest = "SELECT comm_count FROM hotels WHERE hotel_id = $1"

const GetCommentsPostgreRequest = "SELECT DISTINCT us.user_id,comm_id,message,rating,concat($5::varchar,avatar),username,comm.hotel_id,time FROM comments as comm " +
	"INNER JOIN users as us on us.user_id = comm.user_id WHERE comm.hotel_id = $3 AND comm.user_id != $4 LIMIT $2 OFFSET $1"

const AddCommentsPostgreRequest = "INSERT INTO comments VALUES (default, $1, $2,$3,$4) RETURNING comm_id,time"

const DeleteCommentsPostgreRequest = "DELETE FROM comments WHERE comm_id=$1"

const UpdateCommentsPostgreRequest = "UPDATE comments SET message=$2,rating=$3 WHERE comm_id=$1 RETURNING time"

const CheckRateExistForCommentsRequest = "SELECT * FROM comments where hotel_id=$1 and user_id=$2"
