package wishlistDelivery

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/app/wishlist"
	wishlistModel "github.com/go-park-mail-ru/2020_2_JMickhs/internal/app/wishlist/models"
	customerror "github.com/go-park-mail-ru/2020_2_JMickhs/internal/pkg/error"
	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/pkg/logger"
	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/pkg/responses"
	"github.com/gorilla/mux"
)

type WishlistHandler struct {
	useCase wishlist.Usecase
	log     *logger.CustomLogger
}

func NewWishlistHandler(r *mux.Router, useCase wishlist.Usecase, lg *logger.CustomLogger) {
	handler := WishlistHandler{
		useCase: useCase,
		log:     lg,
	}
	r.HandleFunc("/api/v1/wishlist", handler.GetWishlist)
	r.HandleFunc("/api/v1/addhotel", handler.AddHotelToWishlist)
	r.HandleFunc("/api/v1/deletehotel", handler.DeleteHotelFromWishlist)
	r.HandleFunc("/api/v1/deletewishlist", handler.DeleteWishlist)
	r.HandleFunc("/api/v1/createwishlist", handler.CreateWishlist)
}

func (wh *WishlistHandler) GetWishlist(w http.ResponseWriter, r *http.Request) {
	hotels, err := wh.useCase.GetWishlist(42)
	if err != nil {
		customerror.PostError(w, r, wh.log, err, nil)
	}
	responses.SendDataResponse(w, hotels)
}

func (wh *WishlistHandler) AddHotelToWishlist(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	wishlist_id, _ := strconv.Atoi(query["wishlist_id"][0])
	hotel_id, _ := strconv.Atoi(query["hotel_id"][0])
	if err := wh.useCase.AddHotel(hotel_id, wishlist_id); err != nil {
		customerror.PostError(w, r, wh.log, err, nil)
	}
	w.Write([]byte("okay"))
	responses.SendOkResponse(w)
}

func (wh *WishlistHandler) DeleteHotelFromWishlist(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	wishlist_id, _ := strconv.Atoi(query["wishlist_id"][0])
	hotel_id, _ := strconv.Atoi(query["hotel_id"][0])
	if err := wh.useCase.DeleteHotel(hotel_id, wishlist_id); err != nil {
		customerror.PostError(w, r, wh.log, err, nil)
	}
	w.Write([]byte("okay"))
	responses.SendOkResponse(w)
}

func (wh *WishlistHandler) DeleteWishlist(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	wishlist_id, _ := strconv.Atoi(query["wishlist_id"][0])
	if err := wh.useCase.DeleteWishlist(wishlist_id); err != nil {
		customerror.PostError(w, r, wh.log, err, nil)
	}
	w.Write([]byte("okay"))
	responses.SendOkResponse(w)
}

func (wh *WishlistHandler) CreateWishlist(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)
	newWishlist := new(wishlistModel.Wishlist)
	err := decoder.Decode(newWishlist)
	if err != nil {
		customerror.PostError(w, r, wh.log, err, nil)
		w.Write([]byte("{}"))
		return
	}

	// query := r.URL.Query()
	// wishlist_id := query["wishlist_id"][0]
	// name := query["name"][0]
	// user_id := query["user_id"][0]
	// wishlist := wishlistModel.Wishlist{wishlist_id, name, user_id}
	if err := wh.useCase.CreateWishlist(*newWishlist); err != nil {
		customerror.PostError(w, r, wh.log, err, nil)
	}
	w.Write([]byte("okay"))
	responses.SendOkResponse(w)
}
