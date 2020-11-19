package hotelRepository

import (
	"fmt"
	"strconv"
	"strings"

	commModel "github.com/go-park-mail-ru/2020_2_JMickhs/JMickhs_main/internal/app/comment/models"

	"github.com/go-park-mail-ru/2020_2_JMickhs/JMickhs_main/configs"
	hotelmodel "github.com/go-park-mail-ru/2020_2_JMickhs/JMickhs_main/internal/app/hotels/models"
	"github.com/go-park-mail-ru/2020_2_JMickhs/JMickhs_main/pkg/clientError"
	customerror "github.com/go-park-mail-ru/2020_2_JMickhs/JMickhs_main/pkg/error"
	"github.com/go-park-mail-ru/2020_2_JMickhs/JMickhs_main/pkg/serverError"

	"github.com/jmoiron/sqlx"
)

type PostgreHotelRepository struct {
	conn *sqlx.DB
}

func NewPostgresHotelRepository(conn *sqlx.DB) PostgreHotelRepository {
	return PostgreHotelRepository{conn}
}

func (p *PostgreHotelRepository) GetHotels(StartID int) ([]hotelmodel.Hotel, error) {
	hotels := []hotelmodel.Hotel{}
	err := p.conn.Select(&hotels, GetHotelsPostgreRequest, strconv.Itoa(StartID), configs.S3Url)

	if err != nil {
		return hotels, customerror.NewCustomError(err, serverError.ServerInternalError, 1)
	}
	return hotels, nil
}

func (p *PostgreHotelRepository) GetHotelByID(ID int) (hotelmodel.Hotel, error) {
	hotel := hotelmodel.Hotel{}
	err := p.conn.QueryRow(GetHotelByIDPostgreRequest, strconv.Itoa(ID), configs.S3Url).
		Scan(&hotel.HotelID, &hotel.Name, &hotel.Description, &hotel.Image, &hotel.Location,
			&hotel.Rating, &hotel.CommCount, &hotel.Latitude, &hotel.Longitude)
	if err != nil {
		return hotel, customerror.NewCustomError(err, clientError.Gone, 1)
	}

	err = p.conn.Select(&hotel.Photos, GetHotelsPhotosPostgreRequest, strconv.Itoa(ID), configs.S3Url)
	if err != nil {
		return hotel, customerror.NewCustomError(err, serverError.ServerInternalError, 1)
	}

	return hotel, nil
}

func (p *PostgreHotelRepository) BuildQueryForCommentsPercent(filter *hotelmodel.HotelFiltering, param string) string {
	if filter.CommCountConstraint == "" || filter.CommCountPercent == "" {
		return ""
	}
	query := ""
	numbers := strings.Split(filter.CommCountConstraint, ",")
	for _, numberStr := range numbers {
		number, err := strconv.Atoi(numberStr)
		if err != nil {
			continue
		}
		query += fmt.Sprintf(" AND  comm_count_for_each[%d]::real/(case comm_count when 0 then 1 else comm_count end)::real >= %s::real/100::real  ",
			number+1, param)
	}
	return query
}

func (p *PostgreHotelRepository) BuildQueryToFetchHotel(filter *hotelmodel.HotelFiltering) string {
	baseQuery := fmt.Sprint("SELECT hotel_id, name, description, location, concat($4::varchar,img),country,city,curr_rating , " +
		"comm_count,strict_word_similarity($1,name) as t1,strict_word_similarity($1,location) as t2 ")

	baseQuery += fmt.Sprint(" FROM hotels ", SearchHotelsPostgreRequest)

	NearestFilterQuery := ""
	if filter.Radius != "" {
		NearestFilterQuery = fmt.Sprint(" AND ST_Distance(coordinates::geography, $7::geography)<$8")

		baseQuery += p.BuildQueryForCommentsPercent(filter, "$9")

	} else {
		baseQuery += p.BuildQueryForCommentsPercent(filter, "$7")
	}
	baseQuery += NearestFilterQuery

	RatingFilterQuery := fmt.Sprint(" AND curr_rating >= $5 ")
	if filter.RatingFilterStartNumber == "" {
		filter.RatingFilterStartNumber = "0"
	}
	baseQuery += RatingFilterQuery

	CommentFilterQuery := fmt.Sprint(" AND comm_count >= $6")
	if filter.CommentsFilterStartNumber == "" {
		filter.CommentsFilterStartNumber = "0"
	}
	baseQuery += CommentFilterQuery

	rateOrderQuery := fmt.Sprint(" ORDER BY curr_rating DESC,t1 DESC,t2 DESC ")

	baseQuery += rateOrderQuery
	query := fmt.Sprint(baseQuery, " LIMIT $3 OFFSET $2")

	return query
}

func (p *PostgreHotelRepository) FetchHotels(filter hotelmodel.HotelFiltering, pattern string, offset int) ([]hotelmodel.Hotel, error) {
	query := p.BuildQueryToFetchHotel(&filter)
	fmt.Println(query)
	point := p.GeneratePointToGeo(filter.Latitude, filter.Longitude)
	hotels := []hotelmodel.Hotel{}
	p.conn.Exec("Select set_limit(0.18)")

	udb := p.conn.Unsafe()
	var err error

	if filter.Radius == "" {
		if filter.CommCountPercent == "" {
			err = udb.Select(&hotels, query, pattern, offset,
				configs.BaseItemPerPage, configs.S3Url, filter.RatingFilterStartNumber,
				filter.CommentsFilterStartNumber)
		} else {
			err = udb.Select(&hotels, query, pattern, offset,
				configs.BaseItemPerPage, configs.S3Url, filter.RatingFilterStartNumber,
				filter.CommentsFilterStartNumber, filter.CommCountPercent)
		}

	} else {
		if filter.CommCountPercent == "" {
			err = udb.Select(&hotels, query, pattern, offset,
				configs.BaseItemPerPage, configs.S3Url, filter.RatingFilterStartNumber, filter.CommentsFilterStartNumber,
				point, filter.Radius)
		} else {
			err = udb.Select(&hotels, query, pattern, offset,
				configs.BaseItemPerPage, configs.S3Url, filter.RatingFilterStartNumber, filter.CommentsFilterStartNumber,
				point, filter.Radius, filter.CommCountPercent)
		}
	}

	if err != nil {
		return hotels, customerror.NewCustomError(err, serverError.ServerInternalError, 1)
	}

	return hotels, nil
}

func (p *PostgreHotelRepository) CheckRateExist(UserID int, HotelID int) (commModel.FullCommentInfo, error) {
	comment := commModel.FullCommentInfo{}

	err := p.conn.QueryRow(CheckRateIfExistPostgreRequest, UserID, HotelID).Scan(&comment.Message, &comment.Time, &comment.HotelID,
		&comment.Avatar, &comment.UserID, &comment.CommID, &comment.Username, &comment.Rating)
	if err != nil {
		return comment, customerror.NewCustomError(err, serverError.ServerInternalError, 1)
	}
	return comment, nil
}

func (p *PostgreHotelRepository) GetHotelsPreview(pattern string) ([]hotelmodel.HotelPreview, error) {
	query := fmt.Sprint("SELECT hotel_id, name, location, concat($3::varchar,img) FROM hotels ",
		SearchHotelsPostgreRequest, " ORDER BY curr_rating DESC LIMIT $2")

	hotels := []hotelmodel.HotelPreview{}

	err := p.conn.Select(&hotels, query, pattern, configs.PreviewItemLimit, configs.S3Url)
	if err != nil {
		return hotels, customerror.NewCustomError(err, serverError.ServerInternalError, 1)
	}

	return hotels, nil
}
func (p *PostgreHotelRepository) GeneratePointToGeo(latitude string, longitude string) string {
	return fmt.Sprintf("SRID=4326;POINT(%s %s)", latitude, longitude)
}

func (p *PostgreHotelRepository) GetHotelsByRadius(latitude string, longitude string, radius string) ([]hotelmodel.Hotel, error) {
	point := p.GeneratePointToGeo(latitude, longitude)
	hotels := []hotelmodel.Hotel{}

	err := p.conn.Select(&hotels, GetHotelsByRadiusPostgreRequest, point, radius, configs.BaseItemPerPage, configs.S3Url)
	if err != nil {
		return hotels, customerror.NewCustomError(err, serverError.ServerInternalError, 1)
	}

	return hotels, nil
}
