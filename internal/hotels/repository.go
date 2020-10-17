//go:generate mockgen -source repository.go -destination mocks/hotels_repository_mock.go -package mocks
package hotels

import hotelmodel "github.com/go-park-mail-ru/2020_2_JMickhs/internal/hotels/models"

type Repository interface {
	GetHotels(StartID int) ([]hotelmodel.Hotel, error)
	GetHotelByID(ID int) (hotelmodel.Hotel, error)
	FetchHotels(pattern string, data hotelmodel.FilterData, limit int, nextOrPrev bool) ([]hotelmodel.Hotel, error)
	InsertRating(rating hotelmodel.Rating) error
	UpdateHotelRating(hotelID int, NewRate int) error
	GetCurrentRating(hotelID int) (hotelmodel.RateInfo, error)
}
