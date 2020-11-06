package hotelRepository

import (
	"fmt"
	commModel "github.com/go-park-mail-ru/2020_2_JMickhs/internal/app/comment/models"
	"strconv"

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

func (p *PostgreHotelRepository) FetchHotels(pattern string, offset int) ([]hotelmodel.Hotel, error) {
	query := fmt.Sprint("SELECT hotel_id, name, description, location, concat($4::varchar,img), curr_rating , comm_count FROM hotels ",
		SearchHotelsPostgreRequest, " ORDER BY curr_rating DESC LIMIT $3 OFFSET $2")

	hotels := []hotelmodel.Hotel{}

	err := p.conn.Select(&hotels, query, pattern, offset, configs.BaseItemPerPage, configs.S3Url)
	if err != nil {
		return hotels, customerror.NewCustomError(err, serverError.ServerInternalError, 1)
	}

	return hotels, nil
}

func (p *PostgreHotelRepository) CheckRateExist(UserID int, HotelID int) (commModel.FullCommentInfo, error) {
	comment := commModel.FullCommentInfo{}

	err := p.conn.QueryRow(CheckRateIfExistPostgreRequest, UserID, HotelID).Scan(&comment.Message,&comment.Time,&comment.HotelID,
		&comment.Avatar,&comment.UserID,&comment.CommID,&comment.Username,&comment.Rating)
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

func (p *PostgreHotelRepository) AddHotel(hotel hotelmodel.Hotel) error{
	_, err := p.conn.Exec(AddHotelPostgreRequest,hotel.Name,hotel.Location,hotel.Description,hotel.Image,hotel.Photos)
	if err != nil{
		return customerror.NewCustomError(err, serverError.ServerInternalError, 1)
	}
	return nil
}
