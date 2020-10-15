package hotels

import "github.com/go-park-mail-ru/2020_2_JMickhs/internal/hotels/models"

type Usecase interface {
	GetHotels(StartID int) ([]models.Hotel, error)
	GetHotelByID(ID int) (models.Hotel, error)
	SearchHotel(pattern string, cursor models.Cursor, limit int) (models.SearchData, error)
	DecodeCursor(cursor string) (models.FilterData, error)
	EncodeCursor(data models.FilterData) string
}
