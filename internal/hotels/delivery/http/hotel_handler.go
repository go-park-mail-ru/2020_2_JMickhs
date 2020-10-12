package hotelDelivery

import (
	permissions "github.com/go-park-mail-ru/2020_2_JMickhs/internal/permission"
	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/responses"
	"net/http"
	"strconv"

	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/hotels"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type HotelHandler struct {
	HotelUseCase hotels.Usecase
	log          *logrus.Logger
}

func NewHotelHandler(r *mux.Router, hs hotels.Usecase, lg *logrus.Logger) {
	handler := HotelHandler{
		HotelUseCase: hs,
		log:          lg,
	}

	r.HandleFunc("/api/v1/hotels/{id:[0-9]+}",  permissions.SetCSRF(handler.Hotel)).Methods("GET")
	r.Path("/api/v1/hotels/search").Queries("pattern" , "{pattern}", "from", "{from:[0-9]+}", "limit","{limit:[0-9]+}").
		HandlerFunc(permissions.SetCSRF(handler.SearchHotels)).Methods("GET")
	r.Path("/api/v1/hotels").Queries("from", "{from:[0-9]+}").HandlerFunc( permissions.SetCSRF(handler.ListHotels)).Methods("GET")
}

// swagger:route GET /api/v1/hotels hotel hotels
// GetList of hotels
// responses:
//  200: hotels
func (hh *HotelHandler) ListHotels(w http.ResponseWriter, r *http.Request) {

	from := r.FormValue("from")
	startId, err := strconv.Atoi(from)

	hotels, err := hh.HotelUseCase.GetHotels(startId)

	if err != nil {
		responses.SendErrorResponse(w,http.StatusInternalServerError,err)
		return
	}

	responses.SendOkResponse(w,hotels)
}

// swagger:route GET /api/v1/hotels/{id} hotel hotel
// Get single hotel by id
// responses:
//  200: Hotel
func (hh *HotelHandler) Hotel(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	
	if err != nil {
		hh.log.Error(err.Error())
		responses.SendErrorResponse(w,http.StatusBadRequest,err)
		return
	}

	hotel, err := hh.HotelUseCase.GetHotelByID(id)

	if err != nil {
		hh.log.Error(err.Error())
		responses.SendErrorResponse(w,http.StatusBadRequest,err)
	}

	responses.SendOkResponse(w,hotel)
}


// swagger:route GET /api/v1/hotels/search hotel searchHotel
// Search hotels
// responses:
//  200: searchHotel
func (hh *HotelHandler) SearchHotels (w http.ResponseWriter, r *http.Request) {
	from := r.FormValue("from")
	startId, err := strconv.Atoi(from)

	pattern := r.FormValue("pattern")
	limits := r.FormValue("limit")
	limit, err := strconv.Atoi(limits)

	if err != nil {
		hh.log.Error(err.Error())
		responses.SendErrorResponse(w,http.StatusBadRequest,err)
		return
	}

	hotels, err := hh.HotelUseCase.SearchHotel(pattern,startId,limit)

	if err != nil {
		hh.log.Error(err.Error())
		responses.SendErrorResponse(w,http.StatusBadRequest,err)
	}

	responses.SendOkResponse(w,hotels)
}