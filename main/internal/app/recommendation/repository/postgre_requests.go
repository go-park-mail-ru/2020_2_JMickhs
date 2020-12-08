package recommendRepository

const GetRecommendationsForNonUnouthoriziedRequest = "SELECT hotel_id,name,concat($1::varchar,img),location,curr_rating FROM hotels " +
	" WHERE curr_rating >= 4 ORDER BY random() DESC LIMIT $2"

const GetUsersFromHotelRequest = "SELECT user_id FROM comments where hotel_id = $1"

const GetRecommendationsMatrixRows = "SELECT user_id, hotel_id , rating  FROM comments"

const GetBestRecommendationsRequest = "SELECT hotel_id,name,concat($1::varchar,img),location,curr_rating FROM hotels " +
	" WHERE hotel_id = any($3)  LIMIT $2"

const UpdateRecommendationsForUser = "INSERT INTO recommendations VALUES($1,$2,default) ON CONFLICT(user_id) DO UPDATE set user_id=$1,hotel_id=$2,time=default"

const GetRecommendationsForUser = "SELECT user_id,hotel_id as h,time from recommendations where user_id = $1"
