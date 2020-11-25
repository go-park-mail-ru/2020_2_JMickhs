package wishlistDelivery

import (
	"encoding/json"
	"net/http"

	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/app/hotels"
	hotelModel "github.com/go-park-mail-ru/2020_2_JMickhs/internal/app/hotels/models"
	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/app/wishlist"
	wishlistModel "github.com/go-park-mail-ru/2020_2_JMickhs/internal/app/wishlist/models"
	customerror "github.com/go-park-mail-ru/2020_2_JMickhs/internal/pkg/error"
	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/pkg/logger"
	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/pkg/responses"
	"github.com/gorilla/mux"
)

type WishlistHandler struct {
	useCase      wishlist.Usecase
	hotelUseCase hotels.Usecase
	log          *logger.CustomLogger
}

func NewWishlistHandler(r *mux.Router, useCase wishlist.Usecase, hotelUseCase hotels.Usecase, lg *logger.CustomLogger) {
	handler := WishlistHandler{
		useCase:      useCase,
		hotelUseCase: hotelUseCase,
		log:          lg,
	}
	r.HandleFunc("/api/v1/wishlist", handler.CreateWishlist).Methods("POST")
	r.HandleFunc("/api/v1/wishlist", handler.GetWishlist).Methods("GET")
	r.HandleFunc("/api/v1/wishlist", handler.DeleteWishlist).Methods("DELETE")
	r.HandleFunc("/api/v1/wishlist", handler.AddHotelToWishlist).Methods("PUT")
	r.HandleFunc("/api/v1/wishlistshotel", handler.DeleteHotelFromWishlist).Methods("DELETE")
}

// swagger:route GET /api/v1/wishlist Wishlist getWishlist
// Get list of hotels in wishlist
// responses:
//  200: hotels
//  400: badrequest
func (wh *WishlistHandler) GetWishlist(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)
	WishlistID := wishlistModel.WishlistByIdRequest{}
	err := decoder.Decode(&WishlistID)
	if err != nil {
		customerror.PostError(w, r, wh.log, err, nil)
		return
	}
	hotelsMeta, err := wh.useCase.GetWishlistMeta(WishlistID.ID)
	if err != nil {
		customerror.PostError(w, r, wh.log, err, nil)
		return
	}
	hotels := []hotelModel.MiniHotel{}
	for _, hotel := range hotelsMeta {
		tmpMiniHotel := hotelModel.MiniHotel{}
		tmpMiniHotel, err = wh.hotelUseCase.GetMiniHotelByID(hotel.HotelID)
		if err != nil {
			customerror.PostError(w, r, wh.log, err, nil)
			return
		}
		hotels = append(hotels, tmpMiniHotel)
	}

	responses.SendDataResponse(w, hotels)
}

// swagger:route PUT /api/v1/wishlist Wishlist addHotelToWishlist
// Add hotel to wishlist
// responses:
//  200:
//  400: badrequest
func (wh *WishlistHandler) AddHotelToWishlist(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)
	Data := wishlistModel.HotelWishlistRequest{}
	err := decoder.Decode(&Data)
	if err != nil {
		customerror.PostError(w, r, wh.log, err, nil)
		return
	}

	if err := wh.useCase.AddHotel(Data.HotelID, Data.WishlistID); err != nil {
		customerror.PostError(w, r, wh.log, err, nil)
		return
	}
	responses.SendOkResponse(w)
}

// swagger:route DELETE /api/v1/deletehotel Wishlist deleteHotelFromWishlist
// Delete hotel from wishlist
// responses:
//  200:
//  400: badrequest
func (wh *WishlistHandler) DeleteHotelFromWishlist(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)
	Data := wishlistModel.HotelWishlistRequest{}
	err := decoder.Decode(&Data)
	if err != nil {
		customerror.PostError(w, r, wh.log, err, nil)
		return
	}
	if err := wh.useCase.DeleteHotel(Data.HotelID, Data.WishlistID); err != nil {
		customerror.PostError(w, r, wh.log, err, nil)
		return
	}
	responses.SendOkResponse(w)
}

// swagger:route DELETE /api/v1/wishlist Wishlist deleteWishlist
// Delete wishlist
// responses:
//  200:
//  400: badrequest
func (wh *WishlistHandler) DeleteWishlist(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)
	WishlistID := wishlistModel.WishlistByIdRequest{}
	err := decoder.Decode(&WishlistID)
	if err != nil {
		customerror.PostError(w, r, wh.log, err, nil)
		return
	}
	if err := wh.useCase.DeleteWishlist(WishlistID.ID); err != nil {
		customerror.PostError(w, r, wh.log, err, nil)
		return
	}
	responses.SendOkResponse(w)
}

// swagger:route POST /api/v1/wishlist Wishlist createWishlist
// Creates a new Wishlist
// responses:
//  200: saveWishlist
//  400: badrequest
//  409: conflict
func (wh *WishlistHandler) CreateWishlist(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)
	newWishlist := wishlistModel.Wishlist{}
	err := decoder.Decode(&newWishlist)
	if err != nil {
		customerror.PostError(w, r, wh.log, err, nil)
		return
	}

	if err := wh.useCase.CreateWishlist(newWishlist); err != nil {
		customerror.PostError(w, r, wh.log, err, nil)
		return
	}
	responses.SendOkResponse(w)
}
