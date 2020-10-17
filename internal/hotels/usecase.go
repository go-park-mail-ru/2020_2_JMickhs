//go:generate mockgen -source usecase.go -destination mocks/hotels_usecase_mock.go -package mocks
package hotels

import hotelmodel "github.com/go-park-mail-ru/2020_2_JMickhs/internal/hotels/models"

type Usecase interface {
	GetHotels(StartID int) ([]hotelmodel.Hotel, error)
	GetHotelByID(ID int) (hotelmodel.Hotel, error)
	FetchHotels(pattern string, cursor hotelmodel.Cursor, limit int) (hotelmodel.SearchData, error)
	DecodeCursor(cursor string) (hotelmodel.FilterData, error)
	EncodeCursor(data hotelmodel.FilterData) string
	UpdateRating(rating hotelmodel.Rating) (int, error)
}
