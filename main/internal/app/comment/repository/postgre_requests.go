package commentRepository

const UpdateHotelRatingPostgreRequest = "UPDATE hotels SET curr_rating=$1 where hotel_id = $2"

const GetPrevRatingOnCommentPostgreRequest = "SELECT rating,user_id,hotel_id FROM comments where comm_id = $1"

const GetCurrRatingPostgreRequest = "SELECT round( CAST (curr_rating as numeric),1),comm_count FROM hotels where hotel_id = $1"

const GetCommentsCountPostgreRequest = "SELECT comm_count FROM hotels WHERE hotel_id = $1"

const GetCommentsPostgreRequest = "SELECT comm.user_id,comm_id,message,rating,comm.hotel_id,time FROM comments as comm " +
	"WHERE comm.hotel_id = $3 AND comm.user_id != $4 LIMIT $2 OFFSET $1"

const GetCommentsPhotosPostgreRequest = "SELECT concat($2::varchar,unnest(comm.photos)) FROM comments as comm " +
	" where comm.hotel_id = $1"

const CheckCommentForUpdateRequest = "SELECT user_id FROM comments where comm_id=$1"

const CheckPhotosExistPostgreRequest = "SELECT concat($3::varchar,unnest(c.photos))  " +
	"FROM comments as c  WHERE c.hotel_id=$1 AND c.user_id=$2"

const AddCommentsPostgreRequest = "INSERT INTO comments(comm_id,user_id,hotel_id,message,rating,photos) VALUES (default, $1, $2,$3,$4,$5) RETURNING comm_id,time"

const DeleteCommentsPostgreRequest = "DELETE FROM comments WHERE comm_id=$1"

const UpdateCommentsPostgreRequest = "UPDATE comments SET message=$2,rating=$3,photos=COALESCE($4,photos) WHERE comm_id=$1 RETURNING time"

const GetOneCommentPhotosPostgreRequest = "SELECT concat($2::varchar,unnest(comm.photos)) FROM comments as comm " +
	" where comm.comm_id = $1"

const CheckRateExistForCommentsRequest = "SELECT * FROM comments where hotel_id=$1 and user_id=$2"
