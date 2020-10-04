package hotelUsecase

import (
	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/hotels"
	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/hotels/models"
)

type HotelUseCase struct {
	hotelRepo     hotels.Repository
	staticImgPath string
}

func NewHotelUsecase(r hotels.Repository, path string) *HotelUseCase {
	return &HotelUseCase{
		hotelRepo:     r,
		staticImgPath: path,
	}
}

func (p *HotelUseCase) GetHotels() ([]models.Hotel, error) {
	return p.hotelRepo.GetHotels()
}
func (p *HotelUseCase) GetHotelByID(ID int) (models.Hotel, error) {
	return p.hotelRepo.GetHotelByID(ID)
}
