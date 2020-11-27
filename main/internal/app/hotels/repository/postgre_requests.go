package hotelRepository

const AddHotelByOwner = "INSERT INTO hotels(hotel_id,name,description,email,city,country,location,img,photos) VALUES(default,$1,$2,$3,$4,$5,$6,$7,$8,ST_GeomFromEWKT($9)) RETURNING hotel_id"

const GetHotelsPostgreRequest = "SELECT hotel_id,name,description,concat($2::varchar,img),location,curr_rating,comm_count FROM hotels LIMIT 4 OFFSET $1"

const GetHotelByIDPostgreRequest = "SELECT hotel_id,name,description,concat($2::varchar,img),location,curr_rating,comm_count," +
	"(coordinates::POINT)[0], (coordinates::POINT)[1] FROM hotels WHERE hotel_id=$1 GROUP BY hotel_id"

const GetHotelsPhotosPostgreRequest = "SELECT concat($2::varchar,unnest(photos)) FROM hotels WHERE hotel_id=$1"

const CheckRateIfExistPostgreRequest = "SELECT message,time,c.hotel_id,c.user_id,comm_id,rating " +
	"FROM comments as c  WHERE c.hotel_id=$1"

const SearchHotelsPostgreRequest = "WHERE (lower(name) % lower($1) or lower(location) % lower($1) " +
	"or lower(name) ~* lower($1) or lower(location) ~* lower($1))"

const GetHotelsByRadiusPostgreRequest = "SELECT hotel_id,name,description,concat($4::varchar,img)," +
	"location,curr_rating,comm_count,(coordinates::POINT)[0] as x, (coordinates::POINT)[1] as y  FROM hotels WHERE ST_Distance(coordinates::geography, $1::geography)<$2 ORDER BY coordinates <-> $1 LIMIT $3"

const GetMiniHotelPostgreRequest = "SELECT h.hotel_id,h.name,h.description,h.img,h.location,h.curr_rating FROM hotels AS h WHERE hotel_id = $1"
