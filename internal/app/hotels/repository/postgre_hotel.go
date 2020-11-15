package hotelRepository

import (
	"fmt"
	"strconv"

	commModel "github.com/go-park-mail-ru/2020_2_JMickhs/internal/app/comment/models"

	"github.com/go-park-mail-ru/2020_2_JMickhs/configs"

	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/pkg/clientError"
	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/pkg/serverError"

	hotelmodel "github.com/go-park-mail-ru/2020_2_JMickhs/internal/app/hotels/models"
	customerror "github.com/go-park-mail-ru/2020_2_JMickhs/internal/pkg/error"
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
		Scan(&hotel.HotelID, &hotel.Name, &hotel.Description, &hotel.Image, &hotel.Location, &hotel.Rating, &hotel.CommCount)
	if err != nil {
		return hotel, customerror.NewCustomError(err, clientError.Gone, 1)
	}

	err = p.conn.Select(&hotel.Photos, GetHotelsPhotosPostgreRequest, strconv.Itoa(ID), configs.S3Url)
	if err != nil {
		return hotel, customerror.NewCustomError(err, serverError.ServerInternalError, 1)
	}

	return hotel, nil
}

func (p *PostgreHotelRepository) BuildQueryToFetchHotel(filter hotelmodel.HotelFiltering) string {
	baseQuery := fmt.Sprint("SELECT hotel_id, name, description, location, concat($4::varchar,img), curr_rating , "+
		"comm_count,strict_word_similarity($1,name) as t1,strict_word_similarity($1,location) as t2 FROM hotels ",
		SearchHotelsPostgreRequest)

	RatingFilterQuery := ""
	if filter.RatingFilterStartNumber != -1 {
		if filter.RatingFilterMoreThenStartNumber {
			RatingFilterQuery = fmt.Sprint(" AND curr_rating > $5 ")
		} else {
			RatingFilterQuery = fmt.Sprint(" AND curr_rating < $5 ")
		}
	}

	baseQuery += RatingFilterQuery
	baseQuery += " ORDER BY  t1 DESC,t2 DESC,"

	rateOrderQuery := ""
	if filter.RatingOrdering == true {
		rateOrderQuery = fmt.Sprint("curr_rating DESC")
	} else {
		rateOrderQuery = fmt.Sprint("curr_rating ASC")
	}

	baseQuery += rateOrderQuery

	commentOrderQuery := ""
	if filter.CommentsCountOrdering == true {
		commentOrderQuery = fmt.Sprint(",comm_count DESC")
	} else {
		commentOrderQuery = fmt.Sprint(",comm_count ASC")
	}

	baseQuery += commentOrderQuery
	query := fmt.Sprint(baseQuery, " LIMIT $3 OFFSET $2")

	return query
}

func (p *PostgreHotelRepository) FetchHotels(filter hotelmodel.HotelFiltering, pattern string, offset int) ([]hotelmodel.Hotel, error) {
	query := p.BuildQueryToFetchHotel(filter)
	fmt.Println(query)
	hotels := []hotelmodel.Hotel{}
	p.conn.Exec("Select set_limit(0.25)")
	udb := p.conn.Unsafe()
	var err error
	if filter.RatingFilterStartNumber == -1 {
		err = udb.Select(&hotels, query, pattern, offset, configs.BaseItemPerPage, configs.S3Url)
	} else {
		err = udb.Select(&hotels, query, pattern, offset, configs.BaseItemPerPage, configs.S3Url, filter.RatingFilterStartNumber)
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
