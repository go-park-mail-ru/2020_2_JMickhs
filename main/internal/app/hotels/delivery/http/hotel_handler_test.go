package hotelDelivery

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"

	packageConfig "github.com/go-park-mail-ru/2020_2_JMickhs/package/configs"

	"github.com/go-park-mail-ru/2020_2_JMickhs/main/configs"
	"github.com/spf13/viper"

	commModel "github.com/go-park-mail-ru/2020_2_JMickhs/main/internal/app/comment/models"
	paginationModel "github.com/go-park-mail-ru/2020_2_JMickhs/main/internal/app/paginator/model"
	"github.com/go-park-mail-ru/2020_2_JMickhs/package/clientError"
	customerror "github.com/go-park-mail-ru/2020_2_JMickhs/package/error"
	"github.com/go-park-mail-ru/2020_2_JMickhs/package/logger"
	"github.com/go-park-mail-ru/2020_2_JMickhs/package/responses"
	"github.com/go-park-mail-ru/2020_2_JMickhs/package/serverError"

	"github.com/mitchellh/mapstructure"

	hotels_mock "github.com/go-park-mail-ru/2020_2_JMickhs/main/internal/app/hotels/mocks"
	"github.com/gorilla/mux"

	hotelmodel "github.com/go-park-mail-ru/2020_2_JMickhs/main/internal/app/hotels/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestHotelHandler_Hotel(t *testing.T) {
	testHotel := hotelmodel.Hotel{
		HotelID: 3, Name: "kek", Description: "kekw hotel", Image: "src/image.png", Location: "moscow", Email: "ewrsd@mail.u",
		Country: "russia", City: "moscow", Rating: 3.4, Photos: []string{"fds", "fsd"},
		CommCount: 3, Latitude: 54.33, Longitude: 43.4, WishListExist: viper.GetString(configs.ConfigFields.WishListOut),
	}
	comment := commModel.FullCommentInfo{
		UserID: 3, CommID: 2, HotelID: 3, Message: "kekw", Rating: 3, Avatar: "src/kek.jpg", Username: "kostikan", Time: "22:03:12",
	}
	userID := 3
	t.Run("Hotel", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockHUseCase := hotels_mock.NewMockUsecase(ctrl)

		mockHUseCase.EXPECT().
			GetHotelByID(testHotel.HotelID, userID).
			Return(testHotel, nil)

		mockHUseCase.EXPECT().
			CheckRateExist(3, testHotel.HotelID).
			Return(comment, nil)

		req, err := http.NewRequest("GET", "/api/v1/hotels", nil)
		assert.NoError(t, err)

		req = mux.SetURLVars(req, map[string]string{
			"id": strconv.Itoa(testHotel.HotelID),
		})

		req = req.WithContext(context.WithValue(req.Context(), packageConfig.RequestUserID, userID))
		rec := httptest.NewRecorder()
		handler := HotelHandler{
			HotelUseCase: mockHUseCase,
		}

		handler.Hotel(rec, req)
		resp := rec.Result()
		hotel := hotelmodel.HotelData{}
		body, _ := ioutil.ReadAll(resp.Body)
		response := responses.HttpResponse{}

		err = json.Unmarshal(body, &response)
		assert.NoError(t, err)
		err = mapstructure.Decode(response.Data.(map[string]interface{}), &hotel)
		assert.NoError(t, err)

		assert.Equal(t, hotel.Hotel, testHotel)
		assert.Equal(t, http.StatusOK, response.Code)
	})

	t.Run("HotelErr", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockHUseCase := hotels_mock.NewMockUsecase(ctrl)

		mockHUseCase.EXPECT().
			GetHotelByID(testHotel.HotelID, -1).
			Return(testHotel, customerror.NewCustomError(errors.New(""), clientError.Gone, 1))

		req, err := http.NewRequest("GET", "/api/v1/hotels", nil)
		assert.NoError(t, err)
		req = mux.SetURLVars(req, map[string]string{
			"id": strconv.Itoa(testHotel.HotelID),
		})

		rec := httptest.NewRecorder()
		handler := HotelHandler{
			HotelUseCase: mockHUseCase,
			log:          logger.NewLogger(os.Stdout),
		}

		handler.Hotel(rec, req)
		resp := rec.Result()
		body, _ := ioutil.ReadAll(resp.Body)
		response := responses.HttpResponse{}

		err = json.Unmarshal(body, &response)
		assert.NoError(t, err)

		assert.Equal(t, response.Code, clientError.Gone)
	})

	t.Run("HotelErr2", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockHUseCase := hotels_mock.NewMockUsecase(ctrl)

		req, err := http.NewRequest("GET", "/api/v1/hotels", nil)
		assert.NoError(t, err)

		rec := httptest.NewRecorder()
		handler := HotelHandler{
			HotelUseCase: mockHUseCase,
			log:          logger.NewLogger(os.Stdout),
		}

		handler.Hotel(rec, req)
		resp := rec.Result()
		body, _ := ioutil.ReadAll(resp.Body)
		response := responses.HttpResponse{}

		err = json.Unmarshal(body, &response)
		assert.NoError(t, err)

		assert.Equal(t, response.Code, clientError.BadRequest)
	})

	t.Run("HotelErr2", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockHUseCase := hotels_mock.NewMockUsecase(ctrl)

		mockHUseCase.EXPECT().
			GetHotelByID(testHotel.HotelID, userID).
			Return(testHotel, nil)

		mockHUseCase.EXPECT().
			CheckRateExist(3, testHotel.HotelID).
			Return(comment, customerror.NewCustomError(errors.New(""), serverError.ServerInternalError, 1))

		req, err := http.NewRequest("GET", "/api/v1/hotels", nil)
		assert.NoError(t, err)

		req = mux.SetURLVars(req, map[string]string{
			"id": strconv.Itoa(testHotel.HotelID),
		})
		req = req.WithContext(context.WithValue(req.Context(), packageConfig.RequestUserID, userID))

		rec := httptest.NewRecorder()
		handler := HotelHandler{
			HotelUseCase: mockHUseCase,
			log:          logger.NewLogger(os.Stdout),
		}

		handler.Hotel(rec, req)
		resp := rec.Result()
		hotel := hotelmodel.HotelData{}
		body, _ := ioutil.ReadAll(resp.Body)
		response := responses.HttpResponse{}

		err = json.Unmarshal(body, &response)
		assert.NoError(t, err)

		err = mapstructure.Decode(response.Data.(map[string]interface{}), &hotel)
		assert.NoError(t, err)
		assert.Equal(t, hotel.Hotel, testHotel)
		assert.Equal(t, response.Code, http.StatusOK)
	})
}

func TestHotelHandler_FetchHotels(t *testing.T) {
	testHotels := []hotelmodel.Hotel{
		{HotelID: 3, Name: "kek", Description: "kekw hotel", Image: "src/image.png", Location: "moscow", Email: "ewrsd@mail.u",
			Country: "russia", City: "moscow", Rating: 3.4, Photos: []string{"fds", "fsd"},
			CommCount: 3, Latitude: 54.33, Longitude: 43.4, WishListExist: viper.GetString(configs.ConfigFields.WishListOut)},
		{HotelID: 4, Name: "kek", Description: "kekw hotel", Image: "src/image.png", Location: "moscow", Email: "dsaxcds@mail.ru",
			Country: "russia", City: "moscow", Rating: 3.4, Photos: []string{"fds", "fsd"},
			CommCount: 3, Latitude: 54.33, Longitude: 43.4, WishListExist: viper.GetString(configs.ConfigFields.WishListOut)},
	}

	pagInfo := paginationModel.PaginationInfo{NextLink: "", PrevLink: "", ItemsCount: 3}
	searchData := hotelmodel.SearchData{Hotels: testHotels, PagInfo: pagInfo}

	t.Run("FetchHotels", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockHUseCase := hotels_mock.NewMockUsecase(ctrl)

		filter := hotelmodel.HotelFiltering{}
		mockHUseCase.EXPECT().
			FetchHotels(filter, "kekw", 0, -1).
			Return(searchData, nil)

		req, err := http.NewRequest("GET", "/api/v1/hotels/search?pattern=kekw&page=0", nil)
		assert.NoError(t, err)

		rec := httptest.NewRecorder()
		handler := HotelHandler{
			HotelUseCase: mockHUseCase,
			log:          logger.NewLogger(os.Stdout),
		}

		handler.FetchHotels(rec, req)
		resp := rec.Result()
		var hotels []hotelmodel.Hotel
		body, _ := ioutil.ReadAll(resp.Body)
		response := responses.HttpResponse{}

		err = json.Unmarshal(body, &response)
		assert.NoError(t, err)
		err = mapstructure.Decode(response.Data.(map[string]interface{})["hotels"], &hotels)
		assert.NoError(t, err)

		assert.Equal(t, hotels, searchData.Hotels)
		assert.Equal(t, http.StatusOK, response.Code)
	})

	t.Run("FetchHotelsErr1", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockHUseCase := hotels_mock.NewMockUsecase(ctrl)

		req, err := http.NewRequest("GET", "/api/v1/hotels/search?pattern=kekw", nil)
		assert.NoError(t, err)

		rec := httptest.NewRecorder()
		handler := HotelHandler{
			HotelUseCase: mockHUseCase,
			log:          logger.NewLogger(os.Stdout),
		}

		handler.FetchHotels(rec, req)
		resp := rec.Result()
		body, _ := ioutil.ReadAll(resp.Body)
		response := responses.HttpResponse{}

		err = json.Unmarshal(body, &response)
		assert.NoError(t, err)

		assert.Equal(t, clientError.BadRequest, response.Code)
	})

	t.Run("FetchHotelsErr2", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockHUseCase := hotels_mock.NewMockUsecase(ctrl)
		filter := hotelmodel.HotelFiltering{}
		mockHUseCase.EXPECT().
			FetchHotels(filter, "kekw", 0, -1).
			Return(searchData, customerror.NewCustomError(errors.New("fds"), serverError.ServerInternalError, 1))

		req, err := http.NewRequest("GET", "/api/v1/hotels/search?pattern=kekw&page=0", nil)
		assert.NoError(t, err)

		rec := httptest.NewRecorder()
		handler := HotelHandler{
			HotelUseCase: mockHUseCase,
			log:          logger.NewLogger(os.Stdout),
		}

		handler.FetchHotels(rec, req)
		resp := rec.Result()
		body, _ := ioutil.ReadAll(resp.Body)
		response := responses.HttpResponse{}

		err = json.Unmarshal(body, &response)
		assert.NoError(t, err)
		assert.Equal(t, serverError.ServerInternalError, response.Code)
	})
}

func TestHotelHandler_FetchHotelsPreview(t *testing.T) {
	previews := []hotelmodel.HotelPreview{
		{HotelID: 3, Name: "kekw hotel", Image: "src/image.png", Location: "moscow"},
		{HotelID: 3, Name: "kekw hotel", Image: "src/image.png", Location: "moscow"}}

	t.Run("FetchHotelsPreviews", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockHUseCase := hotels_mock.NewMockUsecase(ctrl)

		mockHUseCase.EXPECT().
			GetHotelsPreview("kekw").
			Return(previews, nil)

		req, err := http.NewRequest("GET", "/api/v1/hotels/searchPreview?pattern=kekw", nil)
		assert.NoError(t, err)

		rec := httptest.NewRecorder()
		handler := HotelHandler{
			HotelUseCase: mockHUseCase,
			log:          logger.NewLogger(os.Stdout),
		}

		handler.FetchHotelsPreview(rec, req)
		resp := rec.Result()
		var hotels []hotelmodel.HotelPreview
		body, _ := ioutil.ReadAll(resp.Body)
		response := responses.HttpResponse{}

		err = json.Unmarshal(body, &response)
		assert.NoError(t, err)
		err = mapstructure.Decode(response.Data.(map[string]interface{})["hotels_preview"], &hotels)
		assert.NoError(t, err)

		assert.Equal(t, hotels, previews)
		assert.Equal(t, http.StatusOK, response.Code)
	})

	t.Run("FetchHotelsPreviews2", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockHUseCase := hotels_mock.NewMockUsecase(ctrl)

		mockHUseCase.EXPECT().
			GetHotelsPreview("kekw").
			Return(previews, customerror.NewCustomError(errors.New("f"), serverError.ServerInternalError, 1))

		req, err := http.NewRequest("GET", "/api/v1/hotels/searchPreview?pattern=kekw", nil)
		assert.NoError(t, err)

		rec := httptest.NewRecorder()
		handler := HotelHandler{
			HotelUseCase: mockHUseCase,
			log:          logger.NewLogger(os.Stdout),
		}

		handler.FetchHotelsPreview(rec, req)
		resp := rec.Result()
		body, _ := ioutil.ReadAll(resp.Body)
		response := responses.HttpResponse{}

		err = json.Unmarshal(body, &response)
		assert.NoError(t, err)
		assert.Equal(t, serverError.ServerInternalError, response.Code)
	})

}

func TestHotelHandler_FetchHotelsByRadius(t *testing.T) {
	testHotels := []hotelmodel.Hotel{
		{HotelID: 3, Name: "kek", Description: "kekw hotel", Image: "src/image.png", Location: "moscow", Email: "ewrsd@mail.u",
			Country: "russia", City: "moscow", Rating: 3.4, Photos: []string{"fds", "fsd"},
			CommCount: 3, Latitude: 54.33, Longitude: 43.4, WishListExist: viper.GetString(configs.ConfigFields.WishListOut)},
		{HotelID: 4, Name: "kek", Description: "kekw hotel", Image: "src/image.png", Location: "moscow", Email: "dsaxcds@mail.ru",
			Country: "russia", City: "moscow", Rating: 3.4, Photos: []string{"fds", "fsd"},
			CommCount: 3, Latitude: 54.33, Longitude: 43.4, WishListExist: viper.GetString(configs.ConfigFields.WishListOut)},
	}

	t.Run("FetchHotelsByRadius", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockHUseCase := hotels_mock.NewMockUsecase(ctrl)

		mockHUseCase.EXPECT().
			GetHotelsByRadius("54.33", "43.4", "5000").
			Return(testHotels, nil)

		req, err := http.NewRequest("GET", "/api/v1/hotels/search?radius=5000&longitude=43.4&latitude=54.33", nil)
		assert.NoError(t, err)

		rec := httptest.NewRecorder()
		handler := HotelHandler{
			HotelUseCase: mockHUseCase,
			log:          logger.NewLogger(os.Stdout),
		}

		handler.FetchHotelsByRadius(rec, req)
		resp := rec.Result()
		var hotels []hotelmodel.Hotel
		body, _ := ioutil.ReadAll(resp.Body)
		response := responses.HttpResponse{}

		err = json.Unmarshal(body, &response)
		assert.NoError(t, err)
		err = mapstructure.Decode(response.Data.(map[string]interface{})["hotels"], &hotels)
		assert.NoError(t, err)

		assert.Equal(t, hotels, testHotels)
		assert.Equal(t, http.StatusOK, response.Code)
	})

	t.Run("FetchHotelsByRadius", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockHUseCase := hotels_mock.NewMockUsecase(ctrl)

		mockHUseCase.EXPECT().
			GetHotelsByRadius("54.33", "43.4", "5000").
			Return(testHotels, customerror.NewCustomError(errors.New("f"), serverError.ServerInternalError, 1))

		req, err := http.NewRequest("GET", "/api/v1/hotels/search?radius=5000&longitude=43.4&latitude=54.33", nil)
		assert.NoError(t, err)

		rec := httptest.NewRecorder()
		handler := HotelHandler{
			HotelUseCase: mockHUseCase,
			log:          logger.NewLogger(os.Stdout),
		}

		handler.FetchHotelsByRadius(rec, req)
		resp := rec.Result()
		body, _ := ioutil.ReadAll(resp.Body)
		response := responses.HttpResponse{}

		err = json.Unmarshal(body, &response)
		assert.NoError(t, err)
		assert.Equal(t, serverError.ServerInternalError, response.Code)
	})

}
