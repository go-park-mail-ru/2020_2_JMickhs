package hotelUsecase

import (
	"encoding/base64"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"strings"

	customerror "github.com/go-park-mail-ru/2020_2_JMickhs/internal/error"
	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/hotels"
	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/hotels/models"
)

type HotelUseCase struct {
	hotelRepo hotels.Repository
}

func NewHotelUsecase(r hotels.Repository) *HotelUseCase {
	return &HotelUseCase{
		hotelRepo: r,
	}
}

func (p *HotelUseCase) GetHotels(StartID int) ([]models.Hotel, error) {
	return p.hotelRepo.GetHotels(StartID)
}
func (p *HotelUseCase) GetHotelByID(ID int) (models.Hotel, error) {
	return p.hotelRepo.GetHotelByID(ID)
}

func (p *HotelUseCase) SearchHotel(pattern string, cursor models.Cursor, limit int) (models.SearchData, error) {
	DataWithCursor := models.SearchData{}
	currCursor := ""
	next := false

	if cursor.PrevCursor != "" {
		currCursor = cursor.PrevCursor
		next = false
		fmt.Println("fdsfsd")
	} else {
		currCursor = cursor.NextCursor
		next = true
	}

	prevCursor, err := p.DecodeCursor(currCursor)
	if err != nil {
		return DataWithCursor, err
	}

	hotels, err := p.hotelRepo.SearchHotel(pattern, prevCursor, limit, next)
	if err != nil {
		return DataWithCursor, err
	}
	if len(hotels) == 0 {
		return DataWithCursor, nil
	}
	lastHotel := hotels[len(hotels)-1]
	FilterData := models.FilterData{lastHotel.Rating, strconv.Itoa(lastHotel.HotelID)}

	nextCursor := p.EncodeCursor(FilterData)
	cursor.NextCursor = nextCursor
	cursor.PrevCursor = currCursor

	DataWithCursor.Hotels = hotels
	DataWithCursor.Cursor = cursor

	return DataWithCursor, nil
}

func (p *HotelUseCase) DecodeCursor(cursor string) (models.FilterData, error) {
	filter := models.FilterData{}
	if cursor == "" {
		filter.ID = "0"
		filter.Rating = strconv.Itoa(math.MaxInt32)
		return filter, nil
	}
	byt, err := base64.StdEncoding.DecodeString(cursor)
	if err != nil {
		return filter, customerror.NewCustomError(err.Error(), http.StatusBadRequest)
	}

	arrStr := strings.Split(string(byt), ",")
	if len(arrStr) != 2 {
		return filter, customerror.NewCustomError("unvalid cursor", http.StatusBadRequest)
	}

	filter.Rating = arrStr[0]
	filter.ID = arrStr[1]
	return filter, nil
}

func (p *HotelUseCase) EncodeCursor(data models.FilterData) string {
	key := fmt.Sprintf("%s,%s", data.Rating, data.ID)
	return base64.StdEncoding.EncodeToString([]byte(key))
}
