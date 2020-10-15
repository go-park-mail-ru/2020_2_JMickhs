package hotels

import "github.com/go-park-mail-ru/2020_2_JMickhs/internal/hotels/models"

type Repository interface {
	GetHotels(StartID int) ([]models.Hotel, error)
	GetHotelByID(ID int) (models.Hotel, error)
	SearchHotel(pattern string, data models.FilterData, limit int, nextOrPrev bool) ([]models.Hotel, error)
}
