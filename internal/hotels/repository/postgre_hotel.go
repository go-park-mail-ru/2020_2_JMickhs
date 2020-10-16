package hotelRepository

import (
	"fmt"
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
	rows := p.conn.QueryRow("SELECT hotel_id,name,description,img,location FROM hotels WHERE hotel_id=$1", strconv.Itoa(ID))
	hotel := models.Hotel{}
	err := rows.Scan(&hotel.HotelID, &hotel.Name, &hotel.Description, &hotel.Image, &hotel.Location)
	if err != nil {
		return hotel, customerror.NewCustomError("such hotel doesn't exist", http.StatusGone)
	}
	return hotel, nil
}

func (p *PostgreHotelRepository) FetchHotels(pattern string, filter models.FilterData, limit int, nextOrPrev bool) ([]models.Hotel, error) {
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
	fmt.Println(filter.Rating)
	rows, err := p.conn.Query(fmt.Sprint("SELECT hotel_id, name, description, location, img, rating FROM hotels WHERE (name % $1"+
		"or location % $1 or name LIKE '%' || $1 || '%' or location LIKE '%' || $1 || '%')  AND (rating ", comprasion, " $4 OR (rating = $4 AND hotel_id ", id,
		" $3)) ORDER BY rating ", order, ", hotel_id ", orderId, " LIMIT $2"), pattern, strconv.Itoa(limit), filter.ID, filter.Rating)
	hotels := []models.Hotel{}
	if err != nil {
		return hotels, customerror.NewCustomError(err.Error(), http.StatusInternalServerError)
	}
	hotel := models.Hotel{}
	for rows.Next() {
		err := rows.Scan(&hotel.HotelID, &hotel.Name, &hotel.Description, &hotel.Location, &hotel.Image, &hotel.Rating)
		if err != nil {
			return hotels, customerror.NewCustomError(err.Error(), http.StatusInternalServerError)
		}
		hotels = append(hotels, hotel)
	}
	return hotels, nil
}
