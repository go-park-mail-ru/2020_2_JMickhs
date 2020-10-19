package hotelUsecase

import (
	"encoding/base64"
	"fmt"
	"math"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/app/hotels"
	hotelmodel "github.com/go-park-mail-ru/2020_2_JMickhs/internal/app/hotels/models"
	customerror "github.com/go-park-mail-ru/2020_2_JMickhs/internal/pkg/error"
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

func reverse(v reflect.Value) reflect.Value {
	result := reflect.MakeSlice(v.Type(), 0, v.Cap())
	for i := v.Len() - 1; i >= 0; i-- {
		result = reflect.Append(result, v.Index(i))
	}
	return result
}

func (p *HotelUseCase) FetchHotels(pattern string, cursor hotelmodel.Cursor, limit int) (hotelmodel.SearchData, error) {
	DataWithCursor := hotelmodel.SearchData{}
	currCursor := ""
	next := false

	if cursor.PrevCursor != "" {
		currCursor = cursor.PrevCursor
		next = false
	} else {
		currCursor = cursor.NextCursor
		next = true
	}

	prevCursor, err := p.DecodeCursor(currCursor)
	if err != nil {
		return DataWithCursor, err
	}

	hotels, err := p.hotelRepo.FetchHotels(pattern, prevCursor, limit, next)
	if err != nil {
		return DataWithCursor, err
	}
	if len(hotels) == 0 {
		return DataWithCursor, nil
	}
	if cursor.PrevCursor != "" {
		hotels = reverse(reflect.ValueOf(hotels)).Interface().([]hotelmodel.Hotel)
	}
	lastHotel := hotels[len(hotels)-1]
	FilterData := hotelmodel.FilterData{lastHotel.Rating, strconv.Itoa(lastHotel.HotelID)}
	nextCursor := p.EncodeCursor(FilterData)

	firstHotel := hotels[0]
	FilterData = hotelmodel.FilterData{firstHotel.Rating, strconv.Itoa(firstHotel.HotelID)}
	fmt.Println(firstHotel.Rating)
	prevNewCursor := p.EncodeCursor(FilterData)
	cursor.NextCursor = nextCursor
	cursor.PrevCursor = prevNewCursor

	DataWithCursor.Hotels = hotels
	DataWithCursor.Cursor = cursor

	return DataWithCursor, nil
}

func (p *HotelUseCase) DecodeCursor(cursor string) (hotelmodel.FilterData, error) {
	filter := hotelmodel.FilterData{}
	if cursor == "" {
		filter.ID = "0"
		filter.Rating = math.MaxFloat64
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

	rate, _ := strconv.ParseFloat(arrStr[0], 64)
	filter.Rating = rate
	filter.ID = arrStr[1]
	return filter, nil
}

func (p *HotelUseCase) EncodeCursor(data hotelmodel.FilterData) string {
	key := fmt.Sprintf("%s,%s", data.Rating, data.ID)
	return base64.StdEncoding.EncodeToString([]byte(key))
}

func (p *HotelUseCase) CheckRateExist(UserID int, HotelID int) (int, error) {
	return p.hotelRepo.CheckRateExist(UserID, HotelID)
}
