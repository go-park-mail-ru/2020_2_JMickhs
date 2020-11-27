package hotelUsecase

import (
	"context"
	"errors"
	"mime/multipart"

	"github.com/go-park-mail-ru/2020_2_JMickhs/main/configs"
	commModel "github.com/go-park-mail-ru/2020_2_JMickhs/main/internal/app/comment/models"
	"github.com/go-park-mail-ru/2020_2_JMickhs/main/internal/app/hotels"
	hotelmodel "github.com/go-park-mail-ru/2020_2_JMickhs/main/internal/app/hotels/models"

	"github.com/go-park-mail-ru/2020_2_JMickhs/package/clientError"

	"github.com/spf13/viper"

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

func (p *HotelUseCase) UploadPhoto(hotel *hotelmodel.Hotel, file multipart.File, contentType string, mainImage bool, iterator int) error {
	filePath, err := p.hotelRepo.UploadPhoto(file, contentType)
	if err != nil {
		return err
	}
	if mainImage {
		hotel.Image = filePath
	} else {
		hotel.Photos = append(hotel.Photos, filePath)
	}
	return nil
}

func (p *HotelUseCase) AddHotel(hotel hotelmodel.Hotel, userID int) error {

	user, err := p.userService.GetUserByID(context.Background(), &userService.UserID{UserID: int64(userID)})
	if err != nil {
		return customerror.NewCustomError(err, clientError.BadRequest, 1)
	}

	geo, err := p.hotelRepo.GetLatitudeLongitudeByLocation(hotel.Location)
	if err != nil {
		return customerror.NewCustomError(err, clientError.BadRequest, 1)
	}
	hotel.Latitude = geo.Lat
	hotel.Longitude = geo.Lon
	err = p.hotelRepo.AddHotel(hotel, userID, user.Email)
	if err != nil {
		return err
	}
	return nil
}

func (p *HotelUseCase) GetHotels(StartID int) ([]hotelmodel.Hotel, error) {
	return p.hotelRepo.GetHotels(StartID)
}

func (p *HotelUseCase) GetHotelByID(ID int) (hotelmodel.Hotel, error) {
	return p.hotelRepo.GetHotelByID(ID)
}

func (p *HotelUseCase) FetchHotels(filter hotelmodel.HotelFiltering, pattern string, page int) (hotelmodel.SearchData, error) {
	pag := hotelmodel.SearchData{}

	offset := page * viper.GetInt(configs.ConfigFields.BaseItemPerPage)
	data, err := p.hotelRepo.FetchHotels(filter, pattern, offset)
	if err != nil {
		return pag, err
	}
	pag.Hotels = data

	if page > 0 && page <= viper.GetInt(configs.ConfigFields.BasePageCount) {
		pag.PagInfo.PrevLink = ""
	}
	if page >= 0 && page < viper.GetInt(configs.ConfigFields.BasePageCount) {
		pag.PagInfo.NextLink = ""
	}

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
		return comment, errors.New("wrong  user id ")
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
