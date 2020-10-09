package hotelUsecase

import (
	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/hotels"
	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/hotels/models"
)

type HotelUseCase struct {
	hotelRepo     hotels.Repository
}

func NewHotelUsecase(r hotels.Repository) *HotelUseCase {
	return &HotelUseCase{
		hotelRepo:     r,
	}
}

func (p *HotelUseCase) GetHotels(StartID int) ([]models.Hotel, error) {
	return p.hotelRepo.GetHotels(StartID)
}
func (p *HotelUseCase) GetHotelByID(ID int) (models.Hotel, error) {
	return p.hotelRepo.GetHotelByID(ID)
}
