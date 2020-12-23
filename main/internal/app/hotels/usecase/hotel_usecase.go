package hotelUsecase

import (
	"context"
	"errors"

	"github.com/go-park-mail-ru/2020_2_JMickhs/main/internal/app/wishlist"

	"github.com/go-park-mail-ru/2020_2_JMickhs/main/configs"
	commModel "github.com/go-park-mail-ru/2020_2_JMickhs/main/internal/app/comment/models"
	"github.com/go-park-mail-ru/2020_2_JMickhs/main/internal/app/hotels"
	hotelmodel "github.com/go-park-mail-ru/2020_2_JMickhs/main/internal/app/hotels/models"

	"github.com/spf13/viper"

	customerror "github.com/go-park-mail-ru/2020_2_JMickhs/package/error"
	userService "github.com/go-park-mail-ru/2020_2_JMickhs/package/proto/user"
	"github.com/go-park-mail-ru/2020_2_JMickhs/package/serverError"
)

type HotelUseCase struct {
	hotelRepo    hotels.Repository
	wishlistRepo wishlist.Repository
	userService  userService.UserServiceClient
}

func NewHotelUsecase(r hotels.Repository, userService userService.UserServiceClient, wishlistRepo wishlist.Repository) *HotelUseCase {
	return &HotelUseCase{
		hotelRepo:    r,
		userService:  userService,
		wishlistRepo: wishlistRepo,
	}
}

func (p *HotelUseCase) GetHotelByID(ID int, userID int) (hotelmodel.Hotel, error) {
	hotel, err := p.hotelRepo.GetHotelByID(ID)
	if err != nil {
		return hotel, err
	}
	if userID == -1 {
		return hotel, nil
	}
	check, err := p.wishlistRepo.CheckHotelInWishlists(userID, hotel.HotelID)
	if err != nil {
		return hotel, err
	}
	hotel.WishListExist = check

	return hotel, nil
}

func (p *HotelUseCase) FetchHotels(filter hotelmodel.HotelFiltering, pattern string, page int, userID int) (hotelmodel.SearchData, error) {
	pag := hotelmodel.SearchData{}

	offset := page * viper.GetInt(configs.ConfigFields.BaseItemPerPage)
	data, err := p.hotelRepo.FetchHotels(filter, pattern, offset)
	if err != nil {
		return pag, err
	}

	if page > 0 && page <= viper.GetInt(configs.ConfigFields.BasePageCount) {
		pag.PagInfo.PrevLink = ""
	}
	if page >= 0 && page < viper.GetInt(configs.ConfigFields.BasePageCount) {
		pag.PagInfo.NextLink = ""
	}

	if userID == -1 {
		pag.Hotels = data
		return pag, nil
	}

	for pos, hotel := range data {
		check, err := p.wishlistRepo.CheckHotelInWishlists(userID, hotel.HotelID)
		if err != nil {
			return pag, err
		}
		data[pos].WishListExist = check
	}
	pag.Hotels = data
	return pag, nil
}

func (p *HotelUseCase) GetHotelsPreview(pattern string) ([]hotelmodel.HotelPreview, error) {
	return p.hotelRepo.GetHotelsPreview(pattern)
}

func (p *HotelUseCase) CheckRateExist(UserID int, HotelID int) (commModel.FullCommentInfo, error) {
	comment, err := p.hotelRepo.CheckRateExist(UserID, HotelID)

	if err != nil {
		return comment, err
	}

	if comment.UserID != UserID {
		return comment, customerror.NewCustomError(errors.New("wrong user id"), serverError.ServerInternalError, 1)
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

func (p *HotelUseCase) GetMiniHotelByID(HotelID int) (hotelmodel.MiniHotel, error) {
	return p.hotelRepo.GetMiniHotelByID(HotelID)
}
