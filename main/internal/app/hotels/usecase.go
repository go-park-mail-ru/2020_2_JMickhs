//go:generate mockgen -source usecase.go -destination mocks/hotels_usecase_mock.go -package hotels_mock
package hotels

import (
	"mime/multipart"

	commModel "github.com/go-park-mail-ru/2020_2_JMickhs/main/internal/app/comment/models"
	hotelmodel "github.com/go-park-mail-ru/2020_2_JMickhs/main/internal/app/hotels/models"
)

type Usecase interface {
	UploadPhoto(hotel *hotelmodel.Hotel, file multipart.File, contentType string, mainImage bool, iterator int) error
	AddHotel(hotel hotelmodel.Hotel, userID int) error
	GetHotels(StartID int) ([]hotelmodel.Hotel, error)
	GetHotelByID(ID int) (hotelmodel.Hotel, error)
	FetchHotels(filter hotelmodel.HotelFiltering, pattern string, page int) (hotelmodel.SearchData, error)
	CheckRateExist(UserID int, HotelID int) (commModel.FullCommentInfo, error)
	GetHotelsPreview(pattern string) ([]hotelmodel.HotelPreview, error)
	GetHotelsByRadius(latitude string, longitude string, radius string) ([]hotelmodel.Hotel, error)
}
