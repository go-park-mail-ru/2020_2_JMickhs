package wishlistDelivery

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	hotels_mock "github.com/go-park-mail-ru/2020_2_JMickhs/main/internal/app/hotels/mocks"
	hotelmodel "github.com/go-park-mail-ru/2020_2_JMickhs/main/internal/app/hotels/models"

	"github.com/mitchellh/mapstructure"

	customerror "github.com/go-park-mail-ru/2020_2_JMickhs/package/error"

	"github.com/go-park-mail-ru/2020_2_JMickhs/package/logger"

	"github.com/go-park-mail-ru/2020_2_JMickhs/package/clientError"

	"github.com/go-park-mail-ru/2020_2_JMickhs/main/configs"
	wishlists_mock "github.com/go-park-mail-ru/2020_2_JMickhs/main/internal/app/wishlist/mocks"
	wishlistModel "github.com/go-park-mail-ru/2020_2_JMickhs/main/internal/app/wishlist/models"
	"github.com/go-park-mail-ru/2020_2_JMickhs/package/responses"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/mailru/easyjson"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestWishlistHandler_AddHotelToWishlist(t *testing.T) {
	userID := 2

	t.Run("AddHotelToWishlist", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockWUseCase := wishlists_mock.NewMockUsecase(ctrl)

		mockWUseCase.EXPECT().
			AddHotel(2, 1, 1).
			Return(nil)

		request := wishlistModel.HotelWishlistRequest{HotelID: 1}
		body, err := easyjson.Marshal(request)
		req, err := http.NewRequest("GET", "/api/v1/wishlists/1", bytes.NewBuffer(body))
		assert.NoError(t, err)

		req = mux.SetURLVars(req, map[string]string{
			"wishList_id": "1",
		})

		req = req.WithContext(context.WithValue(req.Context(), viper.GetString(configs.ConfigFields.RequestUserID), userID))
		rec := httptest.NewRecorder()
		handler := WishlistHandler{
			useCase: mockWUseCase,
		}

		handler.AddHotelToWishlist(rec, req)
		resp := rec.Result()
		body, err = ioutil.ReadAll(resp.Body)
		response := responses.HttpResponse{}

		err = json.Unmarshal(body, &response)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusOK, response.Code)
	})
	t.Run("AddHotelToWishlistErr1", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockWUseCase := wishlists_mock.NewMockUsecase(ctrl)

		request := wishlistModel.HotelWishlistRequest{HotelID: 1}
		body, err := easyjson.Marshal(request)
		req, err := http.NewRequest("GET", "/api/v1/wishlists/1", bytes.NewBuffer(body))
		assert.NoError(t, err)

		req = req.WithContext(context.WithValue(req.Context(), viper.GetString(configs.ConfigFields.RequestUserID), userID))
		rec := httptest.NewRecorder()
		handler := WishlistHandler{
			useCase: mockWUseCase,
			log:     logger.NewLogger(os.Stdout),
		}

		handler.AddHotelToWishlist(rec, req)
		resp := rec.Result()
		body, err = ioutil.ReadAll(resp.Body)
		response := responses.HttpResponse{}

		err = json.Unmarshal(body, &response)
		assert.NoError(t, err)

		assert.Equal(t, clientError.BadRequest, response.Code)
	})
	t.Run("AddHotelToWishlistErr2", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockWUseCase := wishlists_mock.NewMockUsecase(ctrl)

		request := wishlistModel.HotelWishlistRequest{HotelID: 1}
		body, err := easyjson.Marshal(request)
		req, err := http.NewRequest("GET", "/api/v1/wishlists/1", bytes.NewBuffer(body))
		assert.NoError(t, err)

		req = mux.SetURLVars(req, map[string]string{
			"wishList_id": "1",
		})

		rec := httptest.NewRecorder()
		handler := WishlistHandler{
			useCase: mockWUseCase,
			log:     logger.NewLogger(os.Stdout),
		}

		handler.AddHotelToWishlist(rec, req)
		resp := rec.Result()
		body, err = ioutil.ReadAll(resp.Body)
		response := responses.HttpResponse{}

		err = json.Unmarshal(body, &response)
		assert.NoError(t, err)

		assert.Equal(t, clientError.Unauthorizied, response.Code)
	})
	t.Run("AddHotelToWishlist", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockWUseCase := wishlists_mock.NewMockUsecase(ctrl)

		mockWUseCase.EXPECT().
			AddHotel(2, 1, 1).
			Return(customerror.NewCustomError(errors.New("ds"), clientError.BadRequest, 1))

		request := wishlistModel.HotelWishlistRequest{HotelID: 1}
		body, err := easyjson.Marshal(request)
		req, err := http.NewRequest("GET", "/api/v1/wishlists/1", bytes.NewBuffer(body))
		assert.NoError(t, err)

		req = mux.SetURLVars(req, map[string]string{
			"wishList_id": "1",
		})

		req = req.WithContext(context.WithValue(req.Context(), viper.GetString(configs.ConfigFields.RequestUserID), userID))
		rec := httptest.NewRecorder()
		handler := WishlistHandler{
			useCase: mockWUseCase,
			log:     logger.NewLogger(os.Stdout),
		}

		handler.AddHotelToWishlist(rec, req)
		resp := rec.Result()
		body, err = ioutil.ReadAll(resp.Body)
		response := responses.HttpResponse{}

		err = json.Unmarshal(body, &response)
		assert.NoError(t, err)

		assert.Equal(t, clientError.BadRequest, response.Code)
	})
}

func TestWishlistHandler_DeleteHotelFromWishlist(t *testing.T) {
	userID := 2

	t.Run("DeleteHotelFromWishlist", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockWUseCase := wishlists_mock.NewMockUsecase(ctrl)

		mockWUseCase.EXPECT().
			DeleteHotel(2, 1, 1).
			Return(nil)

		request := wishlistModel.HotelWishlistRequest{HotelID: 1}
		body, err := easyjson.Marshal(request)
		req, err := http.NewRequest("GET", "/api/v1/wishlists/1", bytes.NewBuffer(body))
		assert.NoError(t, err)

		req = mux.SetURLVars(req, map[string]string{
			"wishList_id": "1",
		})

		req = req.WithContext(context.WithValue(req.Context(), viper.GetString(configs.ConfigFields.RequestUserID), userID))
		rec := httptest.NewRecorder()
		handler := WishlistHandler{
			useCase: mockWUseCase,
		}

		handler.DeleteHotelFromWishlist(rec, req)
		resp := rec.Result()
		body, err = ioutil.ReadAll(resp.Body)
		response := responses.HttpResponse{}

		err = json.Unmarshal(body, &response)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusOK, response.Code)
	})
	t.Run("DeleteHotelFromWishlistErr1", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockWUseCase := wishlists_mock.NewMockUsecase(ctrl)

		request := wishlistModel.HotelWishlistRequest{HotelID: 1}
		body, err := easyjson.Marshal(request)
		req, err := http.NewRequest("GET", "/api/v1/wishlists/1", bytes.NewBuffer(body))
		assert.NoError(t, err)

		req = req.WithContext(context.WithValue(req.Context(), viper.GetString(configs.ConfigFields.RequestUserID), userID))
		rec := httptest.NewRecorder()
		handler := WishlistHandler{
			useCase: mockWUseCase,
			log:     logger.NewLogger(os.Stdout),
		}

		handler.DeleteHotelFromWishlist(rec, req)
		resp := rec.Result()
		body, err = ioutil.ReadAll(resp.Body)
		response := responses.HttpResponse{}

		err = json.Unmarshal(body, &response)
		assert.NoError(t, err)

		assert.Equal(t, clientError.BadRequest, response.Code)
	})
	t.Run("DeleteHotelFromWishlistErr2", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockWUseCase := wishlists_mock.NewMockUsecase(ctrl)

		request := wishlistModel.HotelWishlistRequest{HotelID: 1}
		body, err := easyjson.Marshal(request)
		req, err := http.NewRequest("GET", "/api/v1/wishlists/1", bytes.NewBuffer(body))
		assert.NoError(t, err)

		req = mux.SetURLVars(req, map[string]string{
			"wishList_id": "1",
		})

		rec := httptest.NewRecorder()
		handler := WishlistHandler{
			useCase: mockWUseCase,
			log:     logger.NewLogger(os.Stdout),
		}

		handler.DeleteHotelFromWishlist(rec, req)
		resp := rec.Result()
		body, err = ioutil.ReadAll(resp.Body)
		response := responses.HttpResponse{}

		err = json.Unmarshal(body, &response)
		assert.NoError(t, err)

		assert.Equal(t, clientError.Unauthorizied, response.Code)
	})
	t.Run("DeleteHotelFromWishlistErr3", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockWUseCase := wishlists_mock.NewMockUsecase(ctrl)

		mockWUseCase.EXPECT().
			DeleteHotel(2, 1, 1).
			Return(customerror.NewCustomError(errors.New("ds"), clientError.BadRequest, 1))

		request := wishlistModel.HotelWishlistRequest{HotelID: 1}
		body, err := easyjson.Marshal(request)
		req, err := http.NewRequest("GET", "/api/v1/wishlists/1", bytes.NewBuffer(body))
		assert.NoError(t, err)

		req = mux.SetURLVars(req, map[string]string{
			"wishList_id": "1",
		})

		req = req.WithContext(context.WithValue(req.Context(), viper.GetString(configs.ConfigFields.RequestUserID), userID))
		rec := httptest.NewRecorder()
		handler := WishlistHandler{
			useCase: mockWUseCase,
			log:     logger.NewLogger(os.Stdout),
		}

		handler.DeleteHotelFromWishlist(rec, req)
		resp := rec.Result()
		body, err = ioutil.ReadAll(resp.Body)
		response := responses.HttpResponse{}

		err = json.Unmarshal(body, &response)
		assert.NoError(t, err)

		assert.Equal(t, clientError.BadRequest, response.Code)
	})
}

func TestWishlistHandler_DeleteWishlist(t *testing.T) {
	userID := 2

	t.Run("DeleteHotelFromWishlist", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockWUseCase := wishlists_mock.NewMockUsecase(ctrl)

		mockWUseCase.EXPECT().
			DeleteWishlist(2, 1).
			Return(nil)

		req, err := http.NewRequest("GET", "/api/v1/wishlists/1", nil)
		assert.NoError(t, err)

		req = mux.SetURLVars(req, map[string]string{
			"wishList_id": "1",
		})

		req = req.WithContext(context.WithValue(req.Context(), viper.GetString(configs.ConfigFields.RequestUserID), userID))
		rec := httptest.NewRecorder()
		handler := WishlistHandler{
			useCase: mockWUseCase,
			log:     logger.NewLogger(os.Stdout),
		}

		handler.DeleteWishlist(rec, req)
		resp := rec.Result()
		body, err := ioutil.ReadAll(resp.Body)
		response := responses.HttpResponse{}

		err = json.Unmarshal(body, &response)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusOK, response.Code)
	})
	t.Run("DeleteHotelFromWishlistErr1", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockWUseCase := wishlists_mock.NewMockUsecase(ctrl)

		req, err := http.NewRequest("GET", "/api/v1/wishlists/1", nil)
		assert.NoError(t, err)

		req = mux.SetURLVars(req, map[string]string{
			"wishList_id": "1",
		})

		rec := httptest.NewRecorder()
		handler := WishlistHandler{
			useCase: mockWUseCase,
			log:     logger.NewLogger(os.Stdout),
		}

		handler.DeleteWishlist(rec, req)
		resp := rec.Result()
		body, err := ioutil.ReadAll(resp.Body)
		response := responses.HttpResponse{}

		err = json.Unmarshal(body, &response)
		assert.NoError(t, err)

		assert.Equal(t, clientError.Unauthorizied, response.Code)
	})
	t.Run("DeleteHotelFromWishlistErr2", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockWUseCase := wishlists_mock.NewMockUsecase(ctrl)

		req, err := http.NewRequest("GET", "/api/v1/wishlists/1", nil)
		assert.NoError(t, err)

		req = req.WithContext(context.WithValue(req.Context(), viper.GetString(configs.ConfigFields.RequestUserID), userID))
		rec := httptest.NewRecorder()
		handler := WishlistHandler{
			useCase: mockWUseCase,
			log:     logger.NewLogger(os.Stdout),
		}

		handler.DeleteWishlist(rec, req)
		resp := rec.Result()
		body, err := ioutil.ReadAll(resp.Body)
		response := responses.HttpResponse{}

		err = json.Unmarshal(body, &response)
		assert.NoError(t, err)

		assert.Equal(t, clientError.BadRequest, response.Code)
	})
	t.Run("DeleteHotelFromWishlistErr3", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockWUseCase := wishlists_mock.NewMockUsecase(ctrl)

		mockWUseCase.EXPECT().
			DeleteWishlist(2, 1).
			Return(customerror.NewCustomError(errors.New(""), clientError.BadRequest, 1))

		req, err := http.NewRequest("GET", "/api/v1/wishlists/1", nil)
		assert.NoError(t, err)

		req = mux.SetURLVars(req, map[string]string{
			"wishList_id": "1",
		})

		req = req.WithContext(context.WithValue(req.Context(), viper.GetString(configs.ConfigFields.RequestUserID), userID))
		rec := httptest.NewRecorder()
		handler := WishlistHandler{
			useCase: mockWUseCase,
			log:     logger.NewLogger(os.Stdout),
		}

		handler.DeleteWishlist(rec, req)
		resp := rec.Result()
		body, err := ioutil.ReadAll(resp.Body)
		response := responses.HttpResponse{}

		err = json.Unmarshal(body, &response)
		assert.NoError(t, err)

		assert.Equal(t, clientError.BadRequest, response.Code)
	})
}

func TestWishlistHandler_CreateWishlist(t *testing.T) {
	userID := 2
	wishList := wishlistModel.Wishlist{Name: "kek", WishlistID: 3, UserID: 2}
	t.Run("CreateWishlist", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockWUseCase := wishlists_mock.NewMockUsecase(ctrl)

		mockWUseCase.EXPECT().
			CreateWishlist(wishList).
			Return(wishList, nil)

		body, err := easyjson.Marshal(wishList)
		req, err := http.NewRequest("GET", "/api/v1/wishlists/1", bytes.NewBuffer(body))
		assert.NoError(t, err)

		req = mux.SetURLVars(req, map[string]string{
			"wishList_id": "1",
		})

		req = req.WithContext(context.WithValue(req.Context(), viper.GetString(configs.ConfigFields.RequestUserID), userID))
		rec := httptest.NewRecorder()
		handler := WishlistHandler{
			useCase: mockWUseCase,
			log:     logger.NewLogger(os.Stdout),
		}

		handler.CreateWishlist(rec, req)
		resp := rec.Result()
		body, err = ioutil.ReadAll(resp.Body)
		response := responses.HttpResponse{}
		wishListTest := wishlistModel.Wishlist{}
		err = json.Unmarshal(body, &response)
		assert.NoError(t, err)
		err = mapstructure.Decode(response.Data.(map[string]interface{}), &wishListTest)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, wishListTest, wishList)
	})
	t.Run("CreateWishlistErr1", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockWUseCase := wishlists_mock.NewMockUsecase(ctrl)

		req, err := http.NewRequest("GET", "/api/v1/wishlists/1", bytes.NewBuffer([]byte("fdsfsd")))
		assert.NoError(t, err)

		req = mux.SetURLVars(req, map[string]string{
			"wishList_id": "1",
		})

		req = req.WithContext(context.WithValue(req.Context(), viper.GetString(configs.ConfigFields.RequestUserID), userID))
		rec := httptest.NewRecorder()
		handler := WishlistHandler{
			useCase: mockWUseCase,
			log:     logger.NewLogger(os.Stdout),
		}

		handler.CreateWishlist(rec, req)
		resp := rec.Result()
		body, err := ioutil.ReadAll(resp.Body)
		response := responses.HttpResponse{}
		err = json.Unmarshal(body, &response)
		assert.NoError(t, err)

		assert.Equal(t, clientError.BadRequest, response.Code)

	})

	t.Run("CreateWishlistErr2", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockWUseCase := wishlists_mock.NewMockUsecase(ctrl)

		body, err := easyjson.Marshal(wishList)
		req, err := http.NewRequest("GET", "/api/v1/wishlists/1", bytes.NewBuffer(body))
		assert.NoError(t, err)

		req = mux.SetURLVars(req, map[string]string{
			"wishList_id": "1",
		})

		rec := httptest.NewRecorder()
		handler := WishlistHandler{
			useCase: mockWUseCase,
			log:     logger.NewLogger(os.Stdout),
		}

		handler.CreateWishlist(rec, req)
		resp := rec.Result()
		body, err = ioutil.ReadAll(resp.Body)
		response := responses.HttpResponse{}
		err = json.Unmarshal(body, &response)
		assert.NoError(t, err)

		assert.Equal(t, clientError.Unauthorizied, response.Code)

	})

	t.Run("CreateWishlistErr3", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockWUseCase := wishlists_mock.NewMockUsecase(ctrl)
		mockWUseCase.EXPECT().
			CreateWishlist(wishList).
			Return(wishList, customerror.NewCustomError(errors.New("fd"), clientError.BadRequest, 1))

		body, err := easyjson.Marshal(wishList)
		req, err := http.NewRequest("GET", "/api/v1/wishlists/1", bytes.NewBuffer(body))
		assert.NoError(t, err)

		req = mux.SetURLVars(req, map[string]string{
			"wishList_id": "1",
		})
		req = req.WithContext(context.WithValue(req.Context(), viper.GetString(configs.ConfigFields.RequestUserID), userID))
		rec := httptest.NewRecorder()
		handler := WishlistHandler{
			useCase: mockWUseCase,
			log:     logger.NewLogger(os.Stdout),
		}

		handler.CreateWishlist(rec, req)
		resp := rec.Result()
		body, err = ioutil.ReadAll(resp.Body)
		response := responses.HttpResponse{}
		err = json.Unmarshal(body, &response)
		assert.NoError(t, err)

		assert.Equal(t, clientError.BadRequest, response.Code)

	})
}

func TestWishlistHandler_GetWishlist(t *testing.T) {
	userID := 2
	hotelsMeta := []wishlistModel.WishlistHotel{{3, 1},
		{3, 2}}
	hotels := []hotelmodel.MiniHotel{
		{HotelID: 1, Name: "kekw", Location: "moscow russia", Image: "img.jpeg", Rating: 4},
		{HotelID: 2, Name: "kekw", Location: "moscow russia", Image: "img.jpeg", Rating: 4},
	}

	t.Run("GetWishlist", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockWUseCase := wishlists_mock.NewMockUsecase(ctrl)
		mockHUseCase := hotels_mock.NewMockUsecase(ctrl)

		mockWUseCase.EXPECT().
			GetWishlistMeta(userID, 3).
			Return(hotelsMeta, nil)
		mockHUseCase.EXPECT().
			GetMiniHotelByID(1).
			Return(hotels[0], nil)
		mockHUseCase.EXPECT().
			GetMiniHotelByID(2).
			Return(hotels[1], nil)

		req, err := http.NewRequest("GET", "/api/v1/wishlists/1", nil)
		assert.NoError(t, err)

		req = mux.SetURLVars(req, map[string]string{
			"wishList_id": "3",
		})

		req = req.WithContext(context.WithValue(req.Context(), viper.GetString(configs.ConfigFields.RequestUserID), userID))
		rec := httptest.NewRecorder()
		handler := WishlistHandler{
			useCase:      mockWUseCase,
			hotelUseCase: mockHUseCase,
			log:          logger.NewLogger(os.Stdout),
		}

		handler.GetWishlist(rec, req)
		resp := rec.Result()
		body, err := ioutil.ReadAll(resp.Body)
		response := responses.HttpResponse{}
		hotelsResp := []hotelmodel.MiniHotel{}
		err = json.Unmarshal(body, &response)
		assert.NoError(t, err)
		err = mapstructure.Decode(response.Data.([]interface{}), &hotelsResp)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, hotels, hotelsResp)
	})

	t.Run("GetWishlistErr1", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockWUseCase := wishlists_mock.NewMockUsecase(ctrl)
		mockHUseCase := hotels_mock.NewMockUsecase(ctrl)

		req, err := http.NewRequest("GET", "/api/v1/wishlists/1", nil)
		assert.NoError(t, err)

		req = req.WithContext(context.WithValue(req.Context(), viper.GetString(configs.ConfigFields.RequestUserID), userID))
		rec := httptest.NewRecorder()
		handler := WishlistHandler{
			useCase:      mockWUseCase,
			hotelUseCase: mockHUseCase,
			log:          logger.NewLogger(os.Stdout),
		}

		handler.GetWishlist(rec, req)
		resp := rec.Result()
		body, err := ioutil.ReadAll(resp.Body)
		response := responses.HttpResponse{}
		err = json.Unmarshal(body, &response)
		assert.NoError(t, err)

		assert.Equal(t, clientError.BadRequest, response.Code)
	})

	t.Run("GetWishlistErr2", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockWUseCase := wishlists_mock.NewMockUsecase(ctrl)
		mockHUseCase := hotels_mock.NewMockUsecase(ctrl)

		req, err := http.NewRequest("GET", "/api/v1/wishlists/1", nil)
		assert.NoError(t, err)

		req = mux.SetURLVars(req, map[string]string{
			"wishList_id": "3",
		})
		rec := httptest.NewRecorder()
		handler := WishlistHandler{
			useCase:      mockWUseCase,
			hotelUseCase: mockHUseCase,
			log:          logger.NewLogger(os.Stdout),
		}

		handler.GetWishlist(rec, req)
		resp := rec.Result()
		body, err := ioutil.ReadAll(resp.Body)
		response := responses.HttpResponse{}
		err = json.Unmarshal(body, &response)
		assert.NoError(t, err)

		assert.Equal(t, clientError.Unauthorizied, response.Code)
	})

	t.Run("GetWishlistErr3", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockWUseCase := wishlists_mock.NewMockUsecase(ctrl)
		mockHUseCase := hotels_mock.NewMockUsecase(ctrl)

		mockWUseCase.EXPECT().
			GetWishlistMeta(userID, 3).
			Return(hotelsMeta, customerror.NewCustomError(errors.New(""), clientError.BadRequest, 1))
		req, err := http.NewRequest("GET", "/api/v1/wishlists/1", nil)
		assert.NoError(t, err)

		req = mux.SetURLVars(req, map[string]string{
			"wishList_id": "3",
		})

		req = req.WithContext(context.WithValue(req.Context(), viper.GetString(configs.ConfigFields.RequestUserID), userID))
		rec := httptest.NewRecorder()
		handler := WishlistHandler{
			useCase:      mockWUseCase,
			hotelUseCase: mockHUseCase,
			log:          logger.NewLogger(os.Stdout),
		}

		handler.GetWishlist(rec, req)
		resp := rec.Result()
		body, err := ioutil.ReadAll(resp.Body)
		response := responses.HttpResponse{}

		err = json.Unmarshal(body, &response)
		assert.NoError(t, err)

		assert.Equal(t, clientError.BadRequest, response.Code)
	})
	t.Run("GetWishlistErr3", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockWUseCase := wishlists_mock.NewMockUsecase(ctrl)
		mockHUseCase := hotels_mock.NewMockUsecase(ctrl)

		mockWUseCase.EXPECT().
			GetWishlistMeta(userID, 3).
			Return(hotelsMeta, nil)
		mockHUseCase.EXPECT().
			GetMiniHotelByID(1).
			Return(hotels[0], customerror.NewCustomError(errors.New(""), clientError.BadRequest, 1))

		req, err := http.NewRequest("GET", "/api/v1/wishlists/1", nil)
		assert.NoError(t, err)

		req = mux.SetURLVars(req, map[string]string{
			"wishList_id": "3",
		})

		req = req.WithContext(context.WithValue(req.Context(), viper.GetString(configs.ConfigFields.RequestUserID), userID))
		rec := httptest.NewRecorder()
		handler := WishlistHandler{
			useCase:      mockWUseCase,
			hotelUseCase: mockHUseCase,
			log:          logger.NewLogger(os.Stdout),
		}

		handler.GetWishlist(rec, req)
		resp := rec.Result()
		body, err := ioutil.ReadAll(resp.Body)
		response := responses.HttpResponse{}

		err = json.Unmarshal(body, &response)
		assert.NoError(t, err)

		assert.Equal(t, clientError.BadRequest, response.Code)
	})

}

func TestWishlistHandler_GetUserWishlists(t *testing.T) {
	userID := 2
	wishLists := wishlistModel.UserWishLists{Wishlists: []wishlistModel.Wishlist{
		{WishlistID: 1, UserID: 2, Name: "kekws"},
		{WishlistID: 2, UserID: 2, Name: "kekwss"},
	}}
	t.Run("GetUserWishlists", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockWUseCase := wishlists_mock.NewMockUsecase(ctrl)

		mockWUseCase.EXPECT().
			GetUserWishlists(userID).
			Return(wishLists, nil)

		req, err := http.NewRequest("GET", "/api/v1/wishlists/1", nil)
		assert.NoError(t, err)

		req = req.WithContext(context.WithValue(req.Context(), viper.GetString(configs.ConfigFields.RequestUserID), userID))
		rec := httptest.NewRecorder()
		handler := WishlistHandler{
			useCase: mockWUseCase,
			log:     logger.NewLogger(os.Stdout),
		}

		handler.GetUserWishlists(rec, req)
		resp := rec.Result()
		body, err := ioutil.ReadAll(resp.Body)
		response := responses.HttpResponse{}
		wishListTest := wishlistModel.UserWishLists{}
		err = json.Unmarshal(body, &response)
		assert.NoError(t, err)
		err = mapstructure.Decode(response.Data.(map[string]interface{}), &wishListTest)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, wishListTest, wishLists)
	})
	t.Run("GetUserWishlistsErr1", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockWUseCase := wishlists_mock.NewMockUsecase(ctrl)

		req, err := http.NewRequest("GET", "/api/v1/wishlists/1", nil)
		assert.NoError(t, err)

		rec := httptest.NewRecorder()
		handler := WishlistHandler{
			useCase: mockWUseCase,
			log:     logger.NewLogger(os.Stdout),
		}

		handler.GetUserWishlists(rec, req)
		resp := rec.Result()
		body, err := ioutil.ReadAll(resp.Body)
		response := responses.HttpResponse{}
		err = json.Unmarshal(body, &response)
		assert.NoError(t, err)

		assert.Equal(t, clientError.Unauthorizied, response.Code)
	})

	t.Run("GetUserWishlistsErr1", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockWUseCase := wishlists_mock.NewMockUsecase(ctrl)
		mockWUseCase.EXPECT().
			GetUserWishlists(userID).
			Return(wishLists, customerror.NewCustomError(errors.New(""), clientError.BadRequest, 1))

		req, err := http.NewRequest("GET", "/api/v1/wishlists/1", nil)
		assert.NoError(t, err)
		req = req.WithContext(context.WithValue(req.Context(), viper.GetString(configs.ConfigFields.RequestUserID), userID))

		rec := httptest.NewRecorder()
		handler := WishlistHandler{
			useCase: mockWUseCase,
			log:     logger.NewLogger(os.Stdout),
		}

		handler.GetUserWishlists(rec, req)
		resp := rec.Result()
		body, err := ioutil.ReadAll(resp.Body)
		response := responses.HttpResponse{}
		err = json.Unmarshal(body, &response)
		assert.NoError(t, err)

		assert.Equal(t, clientError.BadRequest, response.Code)
	})

}
