package hotelDelivery

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/app/hotels/mocks"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestHotelHandler_ListHotels(t *testing.T) {
	var mockHotels models.Hotel
	err := faker.FakeData(&mockHotels)
	assert.NoError(t, err)
	mocksListHotels := make([]models.Hotel, 0)
	mocksListHotels = append(mocksListHotels, mockHotels)

	mockUCase := new(mocks.HotelsUsecase)

	mockUCase.On("GetHotels").Return(mocksListHotels, nil)

	req, err := http.NewRequest("GET", " /api/v1/hotels", strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	handler := HotelHandler{
		HotelUseCase: mockUCase,
	}

	handler.ListHotels(rec, req)
	resp := rec.Result()
	hotels := []models.Hotel{}
	body, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &hotels)
	assert.Equal(t, mocksListHotels, hotels)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockUCase.AssertExpectations(t)
}

func TestHotelHandler_Hotel(t *testing.T) {
	var mockHotel models.Hotel
	err := faker.FakeData(&mockHotel)
	assert.NoError(t, err)
	mockHotel.ID = 1
	mockUCase := new(mocks.HotelsUsecase)

	mockUCase.On("GetHotelByID", int(mockHotel.ID)).Return(mockHotel, nil)

	req, err := http.NewRequest("GET", " /api/v1/hotel/1", strings.NewReader(""))
	req = mux.SetURLVars(req, map[string]string{
		"id": "1",
	})

	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	handler := HotelHandler{
		HotelUseCase: mockUCase,
	}

	handler.Hotel(rec, req)
	resp := rec.Result()
	hotel := models.Hotel{}
	body, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &hotel)
	assert.Equal(t, mockHotel, hotel)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockUCase.AssertExpectations(t)
}
