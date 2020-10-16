package hotelUsecase

import (
	"encoding/base64"
	"fmt"
	"math"
	"net/http"
	"reflect"
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

func reverse(v reflect.Value) reflect.Value {
	result := reflect.MakeSlice(v.Type(), 0, v.Cap())
	for i := v.Len() - 1; i >= 0; i-- {
		result = reflect.Append(result, v.Index(i))
	}
	return result
}

func (p *HotelUseCase) FetchHotels(pattern string, cursor models.Cursor, limit int) (models.SearchData, error) {
	DataWithCursor := models.SearchData{}
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
		hotels = reverse(reflect.ValueOf(hotels)).Interface().([]models.Hotel)
	}
	lastHotel := hotels[len(hotels)-1]
	FilterData := models.FilterData{lastHotel.Rating, strconv.Itoa(lastHotel.HotelID)}
	nextCursor := p.EncodeCursor(FilterData)

	firstHotel := hotels[0]
	FilterData = models.FilterData{firstHotel.Rating, strconv.Itoa(firstHotel.HotelID)}
	fmt.Println(firstHotel.Rating)
	prevNewCursor := p.EncodeCursor(FilterData)
	cursor.NextCursor = nextCursor
	cursor.PrevCursor = prevNewCursor

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
