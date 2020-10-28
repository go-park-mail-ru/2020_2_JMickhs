//go:generate mockgen -source usecase.go -destination mocks/hotels_usecase_mock.go -package hotels_mock
package hotels

import (
	hotelmodel "github.com/go-park-mail-ru/2020_2_JMickhs/internal/app/hotels/models"
	paginationModel "github.com/go-park-mail-ru/2020_2_JMickhs/internal/app/paginator/model"
)

type Usecase interface {
	GetHotels(StartID int) ([]hotelmodel.Hotel, error)
	GetHotelByID(ID int) (hotelmodel.Hotel, error)
	FetchHotels(pattern string, page int) (paginationModel.PaginationModel, error)
	CheckRateExist(UserID int, HotelID int) (int, error)
	GetHotelsPreview(pattern string) ([]hotelmodel.HotelPreview, error)
}
