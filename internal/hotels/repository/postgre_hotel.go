package hotelRepository

import (
	"net/http"
	"strconv"

	customerror "github.com/go-park-mail-ru/2020_2_JMickhs/internal/error"
	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/hotels/models"
	"github.com/jmoiron/sqlx"
)

type PostgreHotelRepository struct {
	conn *sqlx.DB
}

func NewPostgresHotelRepository(conn *sqlx.DB) PostgreHotelRepository {
	return PostgreHotelRepository{conn}
}

func (p *PostgreHotelRepository) GetHotels(StartID int) ([]models.Hotel, error) {
	rows, err := p.conn.Query("SELECT hotel_id,name,description,img,location FROM hotels LIMIT 4 OFFSET $1", strconv.Itoa(StartID))
	defer rows.Close()
	hotels := []models.Hotel{}
	if err != nil {
		return hotels, customerror.NewCustomError(err.Error(), http.StatusInternalServerError)
	}
	hotel := models.Hotel{}
	for rows.Next() {
		err := rows.Scan(&hotel.HotelID, &hotel.Name, &hotel.Description, &hotel.Image, &hotel.Location)
		if err != nil {
			return hotels, customerror.NewCustomError(err.Error(), http.StatusInternalServerError)
		}
		hotels = append(hotels, hotel)
	}
	return hotels, nil
}

func (p *PostgreHotelRepository) GetHotelByID(ID int) (models.Hotel, error) {
	rows, err := p.conn.Query("SELECT hotel_id,name,description,img,location FROM hotels WHERE hotel_id=$1", strconv.Itoa(ID))
	defer rows.Close()
	hotel := models.Hotel{}
	if err != nil {
		return hotel, customerror.NewCustomError(err.Error(), http.StatusInternalServerError)
	}
	for rows.Next() {
		err := rows.Scan(&hotel.HotelID, &hotel.Name, &hotel.Description, &hotel.Image, &hotel.Location)
		if err != nil {
			return hotel, customerror.NewCustomError(err.Error(), http.StatusInternalServerError)
		}
	}
	return hotel, nil
}

func (p *PostgreHotelRepository) SearchHotel(pattern string, startID int, limit int) ([]models.Hotel, error) {
	rows, err := p.conn.Query("SELECT hotel_id, name, description, location, img FROM hotels WHERE name % $1"+
		"or location % $1 or name LIKE '%' || $1 || '%' or location LIKE '%' || $1 || '%' LIMIT $2 OFFSET $3", pattern, strconv.Itoa(limit), strconv.Itoa(startID))
	hotels := []models.Hotel{}
	if err != nil {
		return hotels, customerror.NewCustomError(err.Error(), http.StatusInternalServerError)
	}
	hotel := models.Hotel{}
	for rows.Next() {
		err := rows.Scan(&hotel.HotelID, &hotel.Name, &hotel.Description, &hotel.Location, &hotel.Image)
		if err != nil {
			return hotels, customerror.NewCustomError(err.Error(), http.StatusInternalServerError)
		}
		hotels = append(hotels, hotel)
	}
	return hotels, nil
}
