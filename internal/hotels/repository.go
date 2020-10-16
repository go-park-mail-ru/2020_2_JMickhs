//go:generate mockgen -source repository.go -destination mocks/hotels_repository_mock.go -package mocks
package hotels

import "github.com/go-park-mail-ru/2020_2_JMickhs/internal/hotels/models"

type Repository interface {
	GetHotels(StartID int) ([]models.Hotel, error)
	GetHotelByID(ID int) (models.Hotel, error)
	FetchHotels(pattern string, data models.FilterData, limit int, nextOrPrev bool) ([]models.Hotel, error)
	InsertRating(rating models.Rating) error
	UpdateHotelRating(hotelID int, NewRate int) error
	GetCurrentRating(hotelID int) (models.RateInfo, error)
}
