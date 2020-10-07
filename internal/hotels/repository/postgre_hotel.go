package hotelRepository

import (
	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/hotels/models"
	"github.com/jmoiron/sqlx"
	"strconv"
)

type PostgreHotelRepository struct {
	conn *sqlx.DB
}

func NewPostgresHotelRepository(conn *sqlx.DB) PostgreHotelRepository {
	return PostgreHotelRepository{conn}
}

func (p *PostgreHotelRepository) GetHotels() ([]models.Hotel, error) {
	rows, err := p.conn.Query("SELECT id,name,description,img FROM hotels")
	defer rows.Close()
	hotels := []models.Hotel{}
	if err != nil {
		return hotels, err
	}
	hotel := models.Hotel{}
	for rows.Next() {
		err := rows.Scan(&hotel.ID, &hotel.Name, &hotel.Description, &hotel.Image)
		if err != nil {
			return hotels, err
		}
		hotels = append(hotels, hotel)
	}
	return hotels, nil
}

func (p *PostgreHotelRepository) GetHotelByID(ID int) (models.Hotel, error) {
	rows, err := p.conn.Query("SELECT id,name,description,img FROM hotels WHERE id=$1", strconv.Itoa(ID))
	defer rows.Close()
	hotel := models.Hotel{}
	if err != nil {
		return hotel, err
	}
	for rows.Next() {
		err := rows.Scan(&hotel.ID, &hotel.Name, &hotel.Description, &hotel.Image)
		if err != nil {
			return hotel, err
		}
	}
	return hotel, nil
}
