//go:generate mockgen -source usecase.go -destination mocks/hotels_usecase_mock.go -package mocks
package hotels

import "github.com/go-park-mail-ru/2020_2_JMickhs/internal/hotels/models"

type Usecase interface {
	GetHotels(StartID int) ([]models.Hotel, error)
	GetHotelByID(ID int) (models.Hotel, error)
	FetchHotels(pattern string, cursor models.Cursor, limit int) (models.SearchData, error)
	DecodeCursor(cursor string) (models.FilterData, error)
	EncodeCursor(data models.FilterData) string
	UpdateRating(rating models.Rating) (int, error)
}
