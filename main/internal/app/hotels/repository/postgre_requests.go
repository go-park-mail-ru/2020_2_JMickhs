package hotelRepository

const GetHotelsPostgreRequest = "SELECT hotel_id,name,description,concat($2::varchar,img),location,curr_rating,comm_count FROM hotels LIMIT 4 OFFSET $1"

const GetHotelByIDPostgreRequest = "SELECT hotel_id,name,description,concat($2::varchar,img),location,curr_rating,comm_count," +
	"(coordinates::POINT)[0], (coordinates::POINT)[1] FROM hotels WHERE hotel_id=$1 GROUP BY hotel_id"

const GetHotelsPhotosPostgreRequest = "SELECT concat($2::varchar,unnest(photos)) FROM hotels WHERE hotel_id=$1"

const CheckRateIfExistPostgreRequest = "SELECT message,time,c.hotel_id,c.user_id,comm_id,rating " +
	"FROM comments as c  WHERE c.hotel_id=$1"

const AddHotelPostgreRequest = "INSERT INTO hotels VALUES(default,$1,$2,$3,$4,$5)"

const SearchHotelsPostgreRequest = "WHERE (lower(name) % lower($1) or lower(location) % lower($1) " +
	"or lower(name) ~* lower($1) or lower(location) ~* lower($1))"

const GetHotelsByRadiusPostgreRequest = "SELECT hotel_id,name,description,concat($4::varchar,img)," +
	"location,curr_rating,comm_count,(coordinates::POINT)[0] as x, (coordinates::POINT)[1] as y  FROM hotels WHERE ST_Distance(coordinates::geography, $1::geography)<$2 ORDER BY coordinates <-> $1 LIMIT $3"
