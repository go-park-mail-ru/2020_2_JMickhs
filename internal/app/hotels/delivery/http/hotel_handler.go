package hotelDelivery

import (
	"context"
	"net/http"
	"strconv"

	hotelmodel "github.com/go-park-mail-ru/2020_2_JMickhs/internal/app/hotels/models"

	"github.com/go-park-mail-ru/2020_2_JMickhs/configs"
	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/app/user/models"

	customerror "github.com/go-park-mail-ru/2020_2_JMickhs/internal/pkg/error"

	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/pkg/logger"
	permissions "github.com/go-park-mail-ru/2020_2_JMickhs/internal/pkg/permission"
	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/pkg/responses"

	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/app/hotels"
	"github.com/gorilla/mux"
)

type HotelHandler struct {
	HotelUseCase hotels.Usecase
	log          *logger.CustomLogger
}

func NewHotelHandler(r *mux.Router, hs hotels.Usecase, lg *logger.CustomLogger) {
	handler := HotelHandler{
		HotelUseCase: hs,
		log:          lg,
	}
	r.HandleFunc("/api/v1/hotels/{id:[0-9]+}", permissions.SetCSRF(handler.Hotel)).Methods("GET")
	r.Path("/api/v1/hotels/search").Queries("pattern", "{pattern}", "prev", "{prev}", "next", "{next}", "limit", "{limit:[0-9]+}").
		HandlerFunc(permissions.SetCSRF(handler.FetchHotels)).Methods("GET")
	r.Path("/api/v1/hotels").Queries("from", "{from:[0-9]+}").HandlerFunc(permissions.SetCSRF(handler.ListHotels)).Methods("GET")
}

// swagger:route GET /api/v1/hotels hotel hotels
// GetList of hotels
// responses:
//  200: hotels
//  400: badrequest
func (hh *HotelHandler) ListHotels(w http.ResponseWriter, r *http.Request) {

	from := r.FormValue("from")
	startId, err := strconv.Atoi(from)

	if err != nil {
		customerror.PostError(w, r, hh.log, err, http.StatusBadRequest)
		return
	}

	hotels, err := hh.HotelUseCase.GetHotels(startId)

	if err != nil {
		customerror.PostError(w, r, hh.log, err, nil)
	}

	responses.SendDataResponse(w, hotels)
}

// swagger:route GET /api/v1/hotels/{id} hotel hotel
// Get Single hotel by id,
// if don't rate by curr user, field "rate" should be empty
// responses:
//  200: HotelData
//  400: badrequest
//  410:  gone
func (hh *HotelHandler) Hotel(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		customerror.PostError(w, r, hh.log, err, http.StatusBadRequest)
	}

	hotel, err := hh.HotelUseCase.GetHotelByID(id)

	if err != nil {
		r = r.WithContext(context.WithValue(r.Context(), configs.DeliveryError, err))
		return
	}
	data := hotelmodel.HotelData{Hotel: hotel}

	usr, ok := r.Context().Value(configs.RequestUser).(models.User)
	if !ok {
		responses.SendDataResponse(w, data)
		return
	}

	rate, err := hh.HotelUseCase.CheckRateExist(usr.ID, hotel.HotelID)
	if err != nil {
		responses.SendDataResponse(w, data)
		return
	}
	data.CurrRate = rate

	responses.SendDataResponse(w, data)
}

// swagger:route GET /api/v1/hotels/search hotel searchHotel
// Search hotels
// responses:
//  200: searchHotel
//  400: badrequest
func (hh *HotelHandler) FetchHotels(w http.ResponseWriter, r *http.Request) {
	next := r.FormValue("next")
	before := r.FormValue("prev")

	cursor := hotelmodel.Cursor{next, before}

	pattern := r.FormValue("pattern")
	limits := r.FormValue("limit")
	limit, err := strconv.Atoi(limits)

	if err != nil {
		err := customerror.NewCustomError(err, http.StatusBadRequest, nil)
		r = r.WithContext(context.WithValue(r.Context(), configs.DeliveryError, err))
		return
	}

	hotels, err := hh.HotelUseCase.FetchHotels(pattern, cursor, limit)

	if err != nil {
		customerror.PostError(w, r, hh.log, err, nil)
	}

	responses.SendDataResponse(w, hotels)
}
