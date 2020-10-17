package hotelRepository

import (
	"fmt"
	"net/http"
	"strconv"

	customerror "github.com/go-park-mail-ru/2020_2_JMickhs/internal/error"
	hotelmodel "github.com/go-park-mail-ru/2020_2_JMickhs/internal/hotels/models"
	"github.com/jmoiron/sqlx"
)

type PostgreHotelRepository struct {
	conn *sqlx.DB
}

func NewPostgresHotelRepository(conn *sqlx.DB) PostgreHotelRepository {
	return PostgreHotelRepository{conn}
}

func (p *PostgreHotelRepository) GetHotels(StartID int) ([]hotelmodel.Hotel, error) {
	rows, err := p.conn.Query("SELECT hotel_id,name,description,img,location,curr_rating FROM hotels LIMIT 4 OFFSET $1", strconv.Itoa(StartID))
	defer rows.Close()
	hotels := []hotelmodel.Hotel{}
	if err != nil {
		return hotels, customerror.NewCustomError(err.Error(), http.StatusInternalServerError)
	}
	hotel := hotelmodel.Hotel{}
	for rows.Next() {
		err := rows.Scan(&hotel.HotelID, &hotel.Name, &hotel.Description, &hotel.Image, &hotel.Location, &hotel.Rating)
		if err != nil {
			return hotels, customerror.NewCustomError(err.Error(), http.StatusInternalServerError)
		}
		hotels = append(hotels, hotel)
	}
	return hotels, nil
}

func (p *PostgreHotelRepository) GetHotelByID(ID int) (hotelmodel.Hotel, error) {
	rows := p.conn.QueryRow("SELECT hotel_id,name,description,img,location,curr_rating FROM hotels WHERE hotel_id=$1", strconv.Itoa(ID))
	hotel := hotelmodel.Hotel{}
	err := rows.Scan(&hotel.HotelID, &hotel.Name, &hotel.Description, &hotel.Image, &hotel.Location, &hotel.Rating)
	if err != nil {
		return hotel, customerror.NewCustomError("such hotel doesn't exist", http.StatusGone)
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
	rows, err := p.conn.Query(fmt.Sprint("SELECT hotel_id, name, description, location, img,  curr_rating FROM hotels WHERE (name % $1"+
		"or location % $1 or name LIKE '%' || $1 || '%' or location LIKE '%' || $1 || '%')  AND (curr_rating ", comprasion, " $4 OR (curr_rating = $4 AND hotel_id ", id,
		" $3)) ORDER BY curr_rating ", order, ", hotel_id ", orderId, " LIMIT $2"), pattern, strconv.Itoa(limit), filter.ID, filter.Rating)
	hotels := []hotelmodel.Hotel{}
	if err != nil {
		return hotels, customerror.NewCustomError(err.Error(), http.StatusInternalServerError)
	}
	hotel := hotelmodel.Hotel{}
	for rows.Next() {
		err := rows.Scan(&hotel.HotelID, &hotel.Name, &hotel.Description, &hotel.Location, &hotel.Image, &hotel.Rating)
		if err != nil {
			return hotels, customerror.NewCustomError(err.Error(), http.StatusInternalServerError)
		}
		hotels = append(hotels, hotel)
	}
	return hotels, nil
}

func (p *PostgreHotelRepository) InsertRating(rating hotelmodel.Rating) error {
	err := p.conn.QueryRow("INSERT INTO rating VALUES (default, $1, $2, $3)", rating.HotelID, rating.UserID, rating.Rate).Err()
	if err != nil {
		return customerror.NewCustomError(err.Error(), http.StatusLocked)
	}
	return nil
}

func (p *PostgreHotelRepository) UpdateHotelRating(hotelID int, NewRate int) error {
	err := p.conn.QueryRow("UPDATE hotels SET curr_rating=$1 where hotel_id = $2", NewRate, hotelID).Err()
	if err != nil {
		return customerror.NewCustomError(err.Error(), http.StatusBadRequest)
	}
	return nil
}

func (p *PostgreHotelRepository) GetCurrentRating(hotelID int) (hotelmodel.RateInfo, error) {
	rateInfo := hotelmodel.RateInfo{}
	err := p.conn.QueryRow("SELECT COUNT(*) FROM rating where hotel_id = $1", hotelID).Scan(&rateInfo.RatesCount)
	if err != nil {
		return rateInfo, customerror.NewCustomError(err.Error(), http.StatusInternalServerError)
	}
	err = p.conn.QueryRow("SELECT curr_rating FROM hotels where hotels.hotel_id = $1", hotelID).Scan(&rateInfo.CurrRating)
	if err != nil {
		return rateInfo, customerror.NewCustomError(err.Error(), http.StatusInternalServerError)
	}
	return rateInfo, nil
}
