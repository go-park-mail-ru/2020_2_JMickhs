package hotelDelivery

import (
	"net/http"
	"strconv"

	customerror "github.com/go-park-mail-ru/2020_2_JMickhs/internal/error"

	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/logger"
	permissions "github.com/go-park-mail-ru/2020_2_JMickhs/internal/permission"
	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/responses"

	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/hotels"
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
	r.Path("/api/v1/hotels/search").Queries("pattern", "{pattern}", "from", "{from:[0-9]+}", "limit", "{limit:[0-9]+}").
		HandlerFunc(permissions.SetCSRF(handler.SearchHotels)).Methods("GET")
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
		err := customerror.NewCustomError(err.Error(), http.StatusBadRequest)
		hh.log.LogError(r.Context(), err)
		responses.SendErrorResponse(w, customerror.ParseCode(err))
		return
	}

	hotels, err := hh.HotelUseCase.GetHotels(startId)

	if err != nil {
		hh.log.LogError(r.Context(), err)
		responses.SendErrorResponse(w, http.StatusInternalServerError)
		return
	}

	responses.SendOkResponse(w, hotels)
}

// swagger:route GET /api/v1/hotels/{id} hotel hotel
// Get single hotel by id
// responses:
//  200: Hotel
//  400: badrequest
//  410:  gone
func (hh *HotelHandler) Hotel(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		err := customerror.NewCustomError(err.Error(), http.StatusBadRequest)
		hh.log.LogError(r.Context(), err)
		responses.SendErrorResponse(w, customerror.ParseCode(err))
		return
	}

	hotel, err := hh.HotelUseCase.GetHotelByID(id)

	if err != nil {
		err := err.(*customerror.CustomError)

		hh.log.LogError(r.Context(), err)
		responses.SendErrorResponse(w, customerror.ParseCode(err))
	}

	responses.SendOkResponse(w, hotel)
}

// swagger:route GET /api/v1/hotels/search hotel searchHotel
// Search hotels
// responses:
//  200: searchHotel
//  400: badrequest
func (hh *HotelHandler) SearchHotels(w http.ResponseWriter, r *http.Request) {
	from := r.FormValue("from")
	startId, err := strconv.Atoi(from)

	pattern := r.FormValue("pattern")
	limits := r.FormValue("limit")
	limit, err := strconv.Atoi(limits)

	if err != nil {
		err := customerror.NewCustomError(err.Error(), http.StatusBadRequest)
		hh.log.LogError(r.Context(), err)
		responses.SendErrorResponse(w, customerror.ParseCode(err))
		return
	}

	hotels, err := hh.HotelUseCase.SearchHotel(pattern, startId, limit)

	if err != nil {
		hh.log.LogError(r.Context(), err)
		responses.SendErrorResponse(w, customerror.ParseCode(err))
	}

	responses.SendOkResponse(w, hotels)
}
