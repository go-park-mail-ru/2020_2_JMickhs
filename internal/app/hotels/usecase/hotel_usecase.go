package hotelUsecase

import (
	"github.com/go-park-mail-ru/2020_2_JMickhs/configs"

	paginationModel "github.com/go-park-mail-ru/2020_2_JMickhs/internal/app/paginator/model"

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

func (p *HotelUseCase) FetchHotels(pattern string, page int) (paginationModel.PaginationModel, error) {
	pag := paginationModel.PaginationModel{}

	pag.PagInfo.NumPages = configs.BasePageCount
	pag.PagInfo.PageNum = page
	offset := page * configs.BaseItemsPerPage
	data, err := p.hotelRepo.FetchHotels(pattern, offset)
	if err != nil {
		return pag, err
	}
	pag.List = data

	if page > 0 && page <= configs.BasePageCount {
		pag.PagInfo.HasPrev = true
	}
	if page >= 0 && page < configs.BasePageCount {
		pag.PagInfo.HasNext = true
	}

	return pag, nil
}

func (p *HotelUseCase) CheckRateExist(UserID int, HotelID int) (int, error) {
	return p.hotelRepo.CheckRateExist(UserID, HotelID)
}
