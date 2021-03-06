//go:generate mockgen -source usecase.go -destination mocks/hotels_usecase_mock.go -package hotels_mock
package hotels

import (
	commModel "github.com/go-park-mail-ru/2020_2_JMickhs/main/internal/app/comment/models"
	hotelmodel "github.com/go-park-mail-ru/2020_2_JMickhs/main/internal/app/hotels/models"
)

type Usecase interface {
	GetHotelByID(ID int, userID int) (hotelmodel.Hotel, error)
	FetchHotels(filter hotelmodel.HotelFiltering, pattern string, page int, userID int) (hotelmodel.SearchData, error)
	CheckRateExist(UserID int, HotelID int) (commModel.FullCommentInfo, error)
	GetHotelsPreview(pattern string) ([]hotelmodel.HotelPreview, error)
	GetHotelsByRadius(latitude string, longitude string, radius string) ([]hotelmodel.Hotel, error)
	GetMiniHotelByID(HotelID int) (hotelmodel.MiniHotel, error)
}
