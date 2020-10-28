//go:generate mockgen -source repository.go -destination mocks/hotels_repository_mock.go -package hotels_mock
package hotels

import hotelmodel "github.com/go-park-mail-ru/2020_2_JMickhs/internal/app/hotels/models"

type Repository interface {
	GetHotels(StartID int) ([]hotelmodel.Hotel, error)
	GetHotelByID(ID int) (hotelmodel.Hotel, error)
	FetchHotels(pattern string, offset int) ([]hotelmodel.Hotel, error)
	CheckRateExist(UserID int, HotelID int) (int, error)
	GetHotelsPreview(pattern string) ([]hotelmodel.HotelPreview, error)
}
