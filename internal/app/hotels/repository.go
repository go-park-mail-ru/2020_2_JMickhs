//go:generate mockgen -source repository.go -destination mocks/hotels_repository_mock.go -package hotels_mock
package hotels

import (
	commModel "github.com/go-park-mail-ru/2020_2_JMickhs/internal/app/comment/models"
	hotelmodel "github.com/go-park-mail-ru/2020_2_JMickhs/internal/app/hotels/models"
)

type Repository interface {
	GetHotels(StartID int) ([]hotelmodel.Hotel, error)
	GetHotelByID(ID int) (hotelmodel.Hotel, error)
	FetchHotels(filter hotelmodel.HotelFiltering, pattern string, offset int) ([]hotelmodel.Hotel, error)
	BuildQueryForCommentsPercent(filter *hotelmodel.HotelFiltering, param string) string
	BuildQueryToFetchHotel(filter *hotelmodel.HotelFiltering) string
	CheckRateExist(UserID int, HotelID int) (commModel.FullCommentInfo, error)
	GetHotelsPreview(pattern string) ([]hotelmodel.HotelPreview, error)
	GetHotelsByRadius(latitude string, longitude string, radius string) ([]hotelmodel.Hotel, error)
}
