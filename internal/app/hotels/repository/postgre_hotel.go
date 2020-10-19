package hotelRepository

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/app/sqlrequests"

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
	rows, err := p.conn.Query(sqlrequests.GetHotelsPostgreRequest, strconv.Itoa(StartID))
	defer rows.Close()
	hotels := []hotelmodel.Hotel{}
	if err != nil {
		return hotels, customerror.NewCustomError(err, http.StatusInternalServerError, nil)
	}
	hotel := hotelmodel.Hotel{}
	for rows.Next() {
		err := rows.Scan(&hotel.HotelID, &hotel.Name, &hotel.Description, &hotel.Image, &hotel.Location, &hotel.Rating)
		if err != nil {
			return hotels, customerror.NewCustomError(err, http.StatusInternalServerError, nil)
		}
		hotels = append(hotels, hotel)
	}
	return hotels, nil
}

func (p *PostgreHotelRepository) GetHotelByID(ID int) (hotelmodel.Hotel, error) {

	rows := p.conn.QueryRow(sqlrequests.GetHotelByIDPostgreRequest, strconv.Itoa(ID))
	hotel := hotelmodel.Hotel{}
	err := rows.Scan(&hotel.HotelID, &hotel.Name, &hotel.Location, &hotel.Description, &hotel.Image, &hotel.Rating)
	if err != nil {
		return hotel, customerror.NewCustomError(err, http.StatusGone, nil)
	}

	photosRows, err := p.conn.Query(sqlrequests.GetHotelsPhotosPostgreRequest, strconv.Itoa(ID))
	if err != nil {
		return hotel, customerror.NewCustomError(err, http.StatusGone, nil)
	}
	photo := ""
	for photosRows.Next() {
		err := photosRows.Scan(&photo)
		if err != nil {
			return hotel, customerror.NewCustomError(err, http.StatusInternalServerError, nil)
		}
		hotel.Photos = append(hotel.Photos, photo)
	}

	return hotel, nil
}

func (p *PostgreHotelRepository) FetchHotels(pattern string, filter hotelmodel.FilterData, limit int, nextOrPrev bool) ([]hotelmodel.Hotel, error) {
	comprasion := ""
	id := ""
	order := "DESC"
	orderId := "DESC"
	if nextOrPrev == true {
		comprasion = "<"
		id = "<"
	} else {
		comprasion = ">"
		id = ">"
		order = "ASC"
		orderId = "ASC"
	}
	rows, err := p.conn.Query(fmt.Sprint("SELECT hotel_id, name, description, location, img, round( CAST (curr_rating as numeric),1) FROM hotels WHERE (name % $1"+
		"or location % $1 or name LIKE '%' || $1 || '%' or location LIKE '%' || $1 || '%')  AND (curr_rating ", comprasion, " $4 OR (curr_rating = $4 AND hotel_id ", id,
		" $3)) ORDER BY curr_rating ", order, ", hotel_id ", orderId, " LIMIT $2"), pattern, strconv.Itoa(limit), filter.ID, filter.Rating)
	hotels := []hotelmodel.Hotel{}
	if err != nil {
		return hotels, customerror.NewCustomError(err, http.StatusInternalServerError, nil)
	}
	hotel := hotelmodel.Hotel{}
	for rows.Next() {
		err := rows.Scan(&hotel.HotelID, &hotel.Name, &hotel.Description, &hotel.Location, &hotel.Image, &hotel.Rating)
		if err != nil {
			return hotels, customerror.NewCustomError(err, http.StatusInternalServerError, nil)
		}
		hotels = append(hotels, hotel)
	}
	return hotels, nil
}

func (p *PostgreHotelRepository) CheckRateExist(UserID int, HotelID int) (int, error) {
	rate := -1
	err := p.conn.QueryRow(sqlrequests.CheckRateIfExistPostgreRequest, UserID, HotelID).Scan(&rate)
	if err != nil {
		return rate, customerror.NewCustomError(err, http.StatusInternalServerError, nil)
	}
	return rate, nil
}
