package recommendRepository

const GetRecommendationsForNonUnouthoriziedRequest = "SELECT hotel_id,name,concat($1::varchar,img),location,curr_rating FROM hotels " +
	"ORDER BY curr_rating DESC LIMIT $2"

const GetRecommendationsMatrixRows = "SELECT user_id, hotel_id , rating  FROM comments"

const GetBestRecommendationsRequest = "SELECT hotel_id,name,concat($1::varchar,img),location,curr_rating FROM hotels " +
	" WHERE hotel_id = any($3)  LIMIT $2"
