package hotels

import "github.com/go-park-mail-ru/2020_2_JMickhs/internal/hotels/models"

type Usecase interface {
	GetHotels() ([]models.Hotel, error)
	GetHotelByID(ID int) (models.Hotel, error)
}
