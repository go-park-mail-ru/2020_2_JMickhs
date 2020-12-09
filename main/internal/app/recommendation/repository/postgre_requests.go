package recommendRepository

const GetRecommendationsForNonUnouthoriziedRequest = "SELECT hotel_id,name,concat($1::varchar,img),location,curr_rating FROM hotels " +
	" WHERE curr_rating >= 4 ORDER BY random() DESC LIMIT $2"

const GetUserCommentsRequest = "SELECT hotel_id FROM comments where user_id = $1"

const GetRecommendationsMatrixRows = " SELECT user_id, unnest(subquery.hotel) as hotel1,unnest(subquery.rating) as rating1 from " +
	"(SELECT user_id,array_agg(hotel_id) as hotel,array_agg(rating) as rating  FROM comments group by user_id  having array_agg(hotel_id) && $1) as subquery"

const GetRecommendationFromSearchHistory = "SELECT hotel_id,name,concat($1::varchar,img),location,curr_rating FROM hotels" +
	" WHERE lower(name) ~* $3 OR lower(location) ~* $3 order by curr_rating LIMIT $2"

const GetBestRecommendationsRequest = "SELECT hotel_id,name,concat($1::varchar,img),location,curr_rating FROM hotels " +
	" WHERE hotel_id = any($3)  LIMIT $2"

const UpdateRecommendationsForUser = "INSERT INTO recommendations VALUES($1,$2,default) ON CONFLICT(user_id) DO UPDATE set user_id=$1,hotel_id=$2,time=default"

const GetRecommendationsForUser = "SELECT user_id,hotel_id as h,time from recommendations where user_id = $1"
