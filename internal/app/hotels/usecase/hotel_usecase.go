package hotelUsecase

import (
	"github.com/go-park-mail-ru/2020_2_JMickhs/configs"
	commModel "github.com/go-park-mail-ru/2020_2_JMickhs/internal/app/comment/models"
	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/app/hotels"
	hotelmodel "github.com/go-park-mail-ru/2020_2_JMickhs/internal/app/hotels/models"
)

type HotelUseCase struct {
	hotelRepo hotels.Repository
}

func NewHotelUsecase(r hotels.Repository) *HotelUseCase {
	return &HotelUseCase{
		hotelRepo: r,
	}
}

func (p *HotelUseCase) GetHotels(StartID int) ([]hotelmodel.Hotel, error) {
	return p.hotelRepo.GetHotels(StartID)
}
func (p *HotelUseCase) GetHotelByID(ID int) (hotelmodel.Hotel, error) {
	return p.hotelRepo.GetHotelByID(ID)
}

func (p *HotelUseCase) FetchHotels(pattern string, page int) (hotelmodel.SearchData, error) {
	pag := hotelmodel.SearchData{}

	offset := page * configs.BaseItemPerPage
	data, err := p.hotelRepo.FetchHotels(pattern, offset)
	if err != nil {
		return pag, err
	}
	pag.Hotels = data

	if page > 0 && page <= configs.BasePageCount {
		pag.PagInfo.PrevLink = ""
	}
	if page >= 0 && page < configs.BasePageCount {
		pag.PagInfo.NextLink = ""
	}

	return pag, nil
}

func (p *HotelUseCase) GetHotelsPreview(pattern string) ([]hotelmodel.HotelPreview, error) {
	return p.hotelRepo.GetHotelsPreview(pattern)
}

func (p *HotelUseCase) CheckRateExist(UserID int, HotelID int) (commModel.FullCommentInfo, error) {
	return p.hotelRepo.CheckRateExist(UserID, HotelID)
}
