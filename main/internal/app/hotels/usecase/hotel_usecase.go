package hotelUsecase

import (
	"context"
	"fmt"

	"github.com/go-park-mail-ru/2020_2_JMickhs/JMickhs_main/configs"
	commModel "github.com/go-park-mail-ru/2020_2_JMickhs/JMickhs_main/internal/app/comment/models"
	"github.com/go-park-mail-ru/2020_2_JMickhs/JMickhs_main/internal/app/hotels"
	hotelmodel "github.com/go-park-mail-ru/2020_2_JMickhs/JMickhs_main/internal/app/hotels/models"
	customerror "github.com/go-park-mail-ru/2020_2_JMickhs/package/error"
	userService "github.com/go-park-mail-ru/2020_2_JMickhs/package/proto/user"
	"github.com/go-park-mail-ru/2020_2_JMickhs/package/serverError"
)

type HotelUseCase struct {
	hotelRepo   hotels.Repository
	userService userService.UserServiceClient
}

func NewHotelUsecase(r hotels.Repository, userService userService.UserServiceClient) *HotelUseCase {
	return &HotelUseCase{
		hotelRepo:   r,
		userService: userService,
	}
}

func (p *HotelUseCase) GetHotels(StartID int) ([]hotelmodel.Hotel, error) {
	return p.hotelRepo.GetHotels(StartID)
}
func (p *HotelUseCase) GetHotelByID(ID int) (hotelmodel.Hotel, error) {
	return p.hotelRepo.GetHotelByID(ID)
}

func (p *HotelUseCase) FetchHotels(filter hotelmodel.HotelFiltering, pattern string, page int) (hotelmodel.SearchData, error) {
	pag := hotelmodel.SearchData{}

	offset := page * configs.BaseItemPerPage
	data, err := p.hotelRepo.FetchHotels(filter, pattern, offset)
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
	fmt.Println("here")
	comment, err := p.hotelRepo.CheckRateExist(UserID, HotelID)
	fmt.Println(err)
	if err != nil {
		return comment, err
	}
	user, err := p.userService.GetUserByID(context.Background(), &userService.UserID{UserID: int64(UserID)})
	if err != nil {
		return comment, customerror.NewCustomError(err, serverError.ServerInternalError, 1)
	}
	comment.UserID = int(user.UserID)
	comment.Username = user.Username
	comment.Avatar = user.Avatar
	return comment, nil
}

func (p *HotelUseCase) GetHotelsByRadius(latitude string, longitude string, radius string) ([]hotelmodel.Hotel, error) {
	return p.hotelRepo.GetHotelsByRadius(latitude, longitude, radius)
}
