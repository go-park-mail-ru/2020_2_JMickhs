package wishlistDelivery

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/go-park-mail-ru/2020_2_JMickhs/package/middlewareApi"

	"github.com/go-park-mail-ru/2020_2_JMickhs/main/configs"
	"github.com/spf13/viper"

	"github.com/mailru/easyjson"

	"github.com/go-park-mail-ru/2020_2_JMickhs/package/clientError"

	"github.com/go-park-mail-ru/2020_2_JMickhs/main/internal/app/hotels"
	hotelModel "github.com/go-park-mail-ru/2020_2_JMickhs/main/internal/app/hotels/models"
	"github.com/go-park-mail-ru/2020_2_JMickhs/main/internal/app/wishlist"
	customerror "github.com/go-park-mail-ru/2020_2_JMickhs/package/error"
	"github.com/go-park-mail-ru/2020_2_JMickhs/package/logger"
	"github.com/go-park-mail-ru/2020_2_JMickhs/package/responses"
	"github.com/gorilla/mux"

	wishlistModel "github.com/go-park-mail-ru/2020_2_JMickhs/main/internal/app/wishlist/models"
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
	r.HandleFunc("/api/v1/wishlists", middlewareApi.CheckCSRFOnHandler(handler.CreateWishlist)).Methods("POST")
	r.HandleFunc("/api/v1/wishlists", handler.GetUserWishlists).Methods("GET")
	r.HandleFunc("/api/v1/wishlists/{wishList_id:[0-9]+}", handler.GetWishlist).Methods("GET")
	r.HandleFunc("/api/v1/wishlists/{wishList_id:[0-9]+}", middlewareApi.CheckCSRFOnHandler(handler.DeleteWishlist)).Methods("DELETE")
	r.HandleFunc("/api/v1/wishlists/{wishList_id:[0-9]+}/hotels", middlewareApi.CheckCSRFOnHandler(handler.AddHotelToWishlist)).Methods("POST")
	r.HandleFunc("/api/v1/wishlists/{wishList_id:[0-9]+}/hotels", middlewareApi.CheckCSRFOnHandler(handler.DeleteHotelFromWishlist)).Methods("DELETE")
}

// swagger:route GET /api/v1/wishlists Wishlist getUserWishlist
// Get list of hotels in wishlist
// responses:
//  200: wishlists
//  400: badrequest
func (wh *WishlistHandler) GetUserWishlists(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(viper.GetString(configs.ConfigFields.RequestUserID)).(int)
	if !ok {
		customerror.PostError(w, r, wh.log, errors.New("user unauthorized"), clientError.Unauthorizied)
		return
	}

	wishlists, err := wh.useCase.GetUserWishlists(userID)
	if err != nil {
		customerror.PostError(w, r, wh.log, err, nil)
		return
	}
	responses.SendDataResponse(w, wishlists)
}

// swagger:route GET /api/v1/wishlists/{wishlist_id} Wishlist getWishlist
// Get list of hotels in wishlist
// responses:
//  200: wishlisthotels
//  400: badrequest
//  410: gone
func (wh *WishlistHandler) GetWishlist(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	WishlistID, err := strconv.Atoi(vars["wishList_id"])

	if err != nil {
		customerror.PostError(w, r, wh.log, err, clientError.BadRequest)
		return
	}

	userID, ok := r.Context().Value(viper.GetString(configs.ConfigFields.RequestUserID)).(int)
	if !ok {
		customerror.PostError(w, r, wh.log, errors.New("user unauthorized"), clientError.Unauthorizied)
		return
	}

	hotelsMeta, err := wh.useCase.GetWishlistMeta(userID, WishlistID)
	if err != nil {
		customerror.PostError(w, r, wh.log, err, nil)
		return
	}

	hotels := make([]hotelModel.MiniHotel, 0, len(hotelsMeta))
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

// swagger:route POST /api/v1/wishlists/{wishlist_id}/hotels Wishlist addHotelToWishlist
// Add hotel to wishlist
// responses:
//  200:
//  400: badrequest
//  423: locked
//  403: Forbidden
func (wh *WishlistHandler) AddHotelToWishlist(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	WishlistID, err := strconv.Atoi(vars["wishList_id"])

	if err != nil {
		customerror.PostError(w, r, wh.log, err, clientError.BadRequest)
		return
	}
	request := wishlistModel.HotelWishlistRequest{}
	err = easyjson.UnmarshalFromReader(r.Body, &request)

	if err != nil {
		customerror.PostError(w, r, wh.log, err, nil)
		return
	}

	userID, ok := r.Context().Value(viper.GetString(configs.ConfigFields.RequestUserID)).(int)
	if !ok {
		customerror.PostError(w, r, wh.log, errors.New("user unauthorized"), clientError.Unauthorizied)
		return
	}

	if err := wh.useCase.AddHotel(userID, request.HotelID, WishlistID); err != nil {
		customerror.PostError(w, r, wh.log, err, nil)
		return
	}
	responses.SendOkResponse(w)
}

// swagger:route DELETE /api/v1/wishlists/{wishlist_id}/hotels Wishlist deleteHotelFromWishlist
// Delete hotel from wishlist
// responses:
//  200:
//  400: badrequest
//  410: gone
//  423: locked
//  403: Forbidden
func (wh *WishlistHandler) DeleteHotelFromWishlist(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	WishlistID, err := strconv.Atoi(vars["wishList_id"])

	if err != nil {
		customerror.PostError(w, r, wh.log, err, clientError.BadRequest)
		return
	}
	request := wishlistModel.HotelWishlistRequest{}
	err = easyjson.UnmarshalFromReader(r.Body, &request)

	if err != nil {
		customerror.PostError(w, r, wh.log, err, nil)
		return
	}

	userID, ok := r.Context().Value(viper.GetString(configs.ConfigFields.RequestUserID)).(int)
	if !ok {
		customerror.PostError(w, r, wh.log, errors.New("user unauthorized"), clientError.Unauthorizied)
		return
	}

	if err := wh.useCase.DeleteHotel(userID, request.HotelID, WishlistID); err != nil {
		customerror.PostError(w, r, wh.log, err, nil)
		return
	}
	responses.SendOkResponse(w)
}

// swagger:route DELETE /api/v1/wishlists/{wishlist_id} Wishlist deleteWishlist
// Delete wishlist
// responses:
//  200:
//  400: badrequest
//  410: gone
//  423: locked
//  403: Forbidden
func (wh *WishlistHandler) DeleteWishlist(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	WishlistID, err := strconv.Atoi(vars["wishList_id"])

	if err != nil {
		customerror.PostError(w, r, wh.log, err, clientError.BadRequest)
		return
	}

	userID, ok := r.Context().Value(viper.GetString(configs.ConfigFields.RequestUserID)).(int)
	if !ok {
		customerror.PostError(w, r, wh.log, errors.New("user unauthorized"), clientError.Unauthorizied)
		return
	}

	if err := wh.useCase.DeleteWishlist(userID, WishlistID); err != nil {
		customerror.PostError(w, r, wh.log, err, nil)
		return
	}
	responses.SendOkResponse(w)
}

// swagger:route POST /api/v1/wishlists Wishlist createWishlist
// Creates a new Wishlist
// responses:
//  200: wishlist
//  400: badrequest
//  423: locked
//  403: Forbidden
func (wh *WishlistHandler) CreateWishlist(w http.ResponseWriter, r *http.Request) {

	newWishlist := wishlistModel.Wishlist{}
	err := easyjson.UnmarshalFromReader(r.Body, &newWishlist)
	if err != nil {
		customerror.PostError(w, r, wh.log, err, clientError.BadRequest)
		return
	}

	userID, ok := r.Context().Value(viper.GetString(configs.ConfigFields.RequestUserID)).(int)
	if !ok {
		customerror.PostError(w, r, wh.log, errors.New("user unauthorized"), clientError.Unauthorizied)
		return
	}

	newWishlist.UserID = userID

	wishlist, err := wh.useCase.CreateWishlist(newWishlist)
	if err != nil {
		customerror.PostError(w, r, wh.log, err, nil)
		return
	}

	responses.SendDataResponse(w, wishlist)
}
