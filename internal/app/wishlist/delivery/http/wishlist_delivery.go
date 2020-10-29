package wishlistDelivery

import (
	"net/http"

	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/app/wishlist"
	customerror "github.com/go-park-mail-ru/2020_2_JMickhs/internal/pkg/error"
	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/pkg/logger"
	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/pkg/responses"
	"github.com/gorilla/mux"
)

type WishlistHandler struct {
	useCase wishlist.Usecase
	log     *logger.CustomLogger
}

func NewWishlistHandler(r *mux.Router, useCase wishlist.Usecase, lg *logger.CustomLogger) *WishlistHandler {
	return &WishlistHandler{
		useCase: useCase,
		log:     lg,
	}
}

func (wh *WishlistHandler) GetWishlist(w http.ResponseWriter, r *http.Request) {
	// как в реквесте передать значение wishListID?
	hotels, err := wh.useCase.GetWishlist(1337)

	if err != nil {
		customerror.PostError(w, r, wh.log, err, nil)
	}
	responses.SendDataResponse(w, hotels)
}

func (wh *WishlistHandler) AddHotelToWishlist(w http.ResponseWriter, r *http.Request) {
	// TODO:
}

func (wh *WishlistHandler) DeleteHotelFromWishlist(w http.ResponseWriter, r *http.Request) {
	// TODO:
}

func (wh *WishlistHandler) DeleteWishlist(w http.ResponseWriter, r *http.Request) {
	// TODO:
}

func (wh *WishlistHandler) CreateWishlist(w http.ResponseWriter, r *http.Request) {
	// TODO:
}
