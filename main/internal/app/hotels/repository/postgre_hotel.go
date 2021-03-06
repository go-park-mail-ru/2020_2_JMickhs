package hotelRepository

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/go-park-mail-ru/2020_2_JMickhs/main/configs"
	commModel "github.com/go-park-mail-ru/2020_2_JMickhs/main/internal/app/comment/models"
	hotelmodel "github.com/go-park-mail-ru/2020_2_JMickhs/main/internal/app/hotels/models"

	"github.com/aws/aws-sdk-go/service/s3"

	"github.com/spf13/viper"

	customerror "github.com/go-park-mail-ru/2020_2_JMickhs/package/error"

	"github.com/go-park-mail-ru/2020_2_JMickhs/package/clientError"
	"github.com/go-park-mail-ru/2020_2_JMickhs/package/serverError"

	"github.com/jmoiron/sqlx"
)

type PostgreHotelRepository struct {
	conn *sqlx.DB
	s3   *s3.S3
}

func NewPostgresHotelRepository(conn *sqlx.DB, s3 *s3.S3) PostgreHotelRepository {
	return PostgreHotelRepository{conn, s3}
}

func (p *PostgreHotelRepository) GetHotelByID(ID int) (hotelmodel.Hotel, error) {
	hotel := hotelmodel.Hotel{}
	err := p.conn.QueryRow(GetHotelByIDPostgreRequest, strconv.Itoa(ID), viper.GetString(configs.ConfigFields.S3Url)).
		Scan(&hotel.HotelID, &hotel.Name, &hotel.Description, &hotel.Image, &hotel.Location,
			&hotel.Rating, &hotel.CommCount, &hotel.Latitude, &hotel.Longitude)
	if err != nil {
		return hotel, customerror.NewCustomError(err, clientError.Gone, 1)
	}

	err = p.conn.Select(&hotel.Photos, GetHotelsPhotosPostgreRequest, strconv.Itoa(ID), viper.GetString(configs.ConfigFields.S3Url))
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
	query += " AND ("
	numbers := strings.Split(filter.CommCountConstraint, ",")
	for pos, numberStr := range numbers {
		number, err := strconv.Atoi(numberStr)
		if err != nil {
			continue
		}
		if pos != len(numbers)-1 {
			query += fmt.Sprintf(" comm_count_for_each[%d]::real/(case comm_count when 0 then 1 else comm_count end)::real >= %s::real/100::real  OR ",
				number+1, param)
		} else {
			query += fmt.Sprintf(" comm_count_for_each[%d]::real/(case comm_count when 0 then 1 else comm_count end)::real >= %s::real/100::real  )",
				number+1, param)
		}

	}
	return query
}

func (p *PostgreHotelRepository) BuildQueryToFetchHotel(filter *hotelmodel.HotelFiltering) string {
	baseQuery := "SELECT hotel_id, name, description, location, concat($4::varchar,img),country,city,curr_rating , " +
		"comm_count,strict_word_similarity($1,name) as t1,strict_word_similarity($1,location) as t2 "

	baseQuery += fmt.Sprint(" FROM hotels ", SearchHotelsPostgreRequest)

	NearestFilterQuery := ""
	if filter.Radius != "" {
		NearestFilterQuery += " AND ST_Distance(coordinates::geography, $8::geography)<$9"
		baseQuery += p.BuildQueryForCommentsPercent(filter, "$10")
	} else {
		baseQuery += p.BuildQueryForCommentsPercent(filter, "$8")
	}
	baseQuery += NearestFilterQuery

	RatingFilterQuery := " AND (curr_rating BETWEEN $5 AND $6 OR curr_rating BETWEEN $6 AND $5) "
	if filter.RatingFilterStartNumber == "" {
		filter.RatingFilterStartNumber = "0"
	}
	baseQuery += RatingFilterQuery

	CommentFilterQuery := " AND comm_count >= $7"
	if filter.CommentsFilterStartNumber == "" {
		filter.CommentsFilterStartNumber = "0"
	}
	baseQuery += CommentFilterQuery

	rateOrderQuery := " ORDER BY curr_rating DESC,t1 DESC,t2 DESC "

	baseQuery += rateOrderQuery
	query := baseQuery + " LIMIT $3 OFFSET $2"

	return query
}

func (p *PostgreHotelRepository) FetchHotels(filter hotelmodel.HotelFiltering, pattern string, offset int) ([]hotelmodel.Hotel, error) {
	query := p.BuildQueryToFetchHotel(&filter)

	point := p.GeneratePointToGeo(filter.Latitude, filter.Longitude)
	hotels := []hotelmodel.Hotel{}
	_, err := p.conn.Exec("Select set_limit(0.18)")
	if err != nil {
		return hotels, err
	}
	udb := p.conn.Unsafe()

	if filter.Radius == "" {
		if filter.CommCountPercent == "" {
			err = udb.Select(&hotels, query, pattern, offset,
				viper.GetInt(configs.ConfigFields.BaseItemPerPage), viper.GetString(configs.ConfigFields.S3Url), filter.RatingFilterStartNumber,
				filter.RatingFilterEndNumber, filter.CommentsFilterStartNumber)
		} else {
			err = udb.Select(&hotels, query, pattern, offset,
				viper.GetInt(configs.ConfigFields.BaseItemPerPage), viper.GetString(configs.ConfigFields.S3Url), filter.RatingFilterStartNumber,
				filter.RatingFilterEndNumber, filter.CommentsFilterStartNumber, filter.CommCountPercent)
		}

	} else {
		if filter.CommCountPercent == "" {
			err = udb.Select(&hotels, query, pattern, offset,
				viper.GetInt(configs.ConfigFields.BaseItemPerPage), viper.GetString(configs.ConfigFields.S3Url),
				filter.RatingFilterStartNumber, filter.RatingFilterEndNumber, filter.CommentsFilterStartNumber,
				point, filter.Radius)
		} else {
			err = udb.Select(&hotels, query, pattern, offset,
				viper.GetInt(configs.ConfigFields.BaseItemPerPage), viper.GetString(configs.ConfigFields.S3Url),
				filter.RatingFilterStartNumber, filter.RatingFilterEndNumber, filter.CommentsFilterStartNumber,
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

	err := p.conn.QueryRow(CheckRateIfExistPostgreRequest, HotelID, UserID).Scan(&comment.Message, &comment.Time, &comment.HotelID,
		&comment.UserID, &comment.CommID, &comment.Rating)

	if err != nil {
		return comment, customerror.NewCustomError(err, serverError.ServerInternalError, 1)
	}

	err = p.conn.Select(&comment.Photos, CheckPhotosExistPostgreRequest, HotelID, UserID, viper.GetString(configs.ConfigFields.S3Url))
	if err != nil {
		return comment, customerror.NewCustomError(err, serverError.ServerInternalError, 1)
	}
	return comment, nil
}

func (p *PostgreHotelRepository) GetHotelsPreview(pattern string) ([]hotelmodel.HotelPreview, error) {
	query := fmt.Sprint("SELECT hotel_id, name, location, concat($3::varchar,img) FROM hotels ",
		SearchHotelsPostgreRequest, " ORDER BY curr_rating DESC LIMIT $2")

	hotels := []hotelmodel.HotelPreview{}

	err := p.conn.Select(&hotels, query, pattern, viper.GetInt(configs.ConfigFields.PreviewItemLimit), viper.GetString(configs.ConfigFields.S3Url))
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
	err := p.conn.Select(&hotels, GetHotelsByRadiusPostgreRequest, point, radius, viper.GetInt(configs.ConfigFields.BaseItemPerPage), viper.GetString(configs.ConfigFields.S3Url))
	if err != nil {
		return hotels, customerror.NewCustomError(err, serverError.ServerInternalError, 1)
	}

	return hotels, nil
}

func (p *PostgreHotelRepository) GetMiniHotelByID(HotelID int) (hotelmodel.MiniHotel, error) {
	rows := p.conn.QueryRow(GetMiniHotelPostgreRequest, strconv.Itoa(HotelID), viper.GetString(configs.ConfigFields.S3Url))
	hotel := hotelmodel.MiniHotel{}
	err := rows.Scan(&hotel.HotelID, &hotel.Name, &hotel.Image, &hotel.Location, &hotel.Rating)
	if err != nil {
		return hotel, customerror.NewCustomError(err, clientError.Gone, 1)
	}
	return hotel, nil
}
