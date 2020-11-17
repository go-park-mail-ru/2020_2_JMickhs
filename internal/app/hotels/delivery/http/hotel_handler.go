package hotelDelivery

import (
	"net/http"
	"strconv"

	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/pkg/clientError"

	hotelmodel "github.com/go-park-mail-ru/2020_2_JMickhs/internal/app/hotels/models"

	"github.com/go-park-mail-ru/2020_2_JMickhs/configs"
	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/app/user/models"

	customerror "github.com/go-park-mail-ru/2020_2_JMickhs/internal/pkg/error"

	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/pkg/logger"
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
	r.HandleFunc("/api/v1/hotels/{id:[0-9]+}", handler.Hotel).Methods("GET")
	r.Path("/api/v1/hotels/search").Queries("pattern", "{pattern}", "page", "{page:[0-9]+}").
		HandlerFunc(handler.FetchHotels).Methods("GET")
	r.Path("/api/v1/hotels").Queries("from", "{from:[0-9]+}").HandlerFunc(handler.ListHotels).Methods("GET")
	r.Path("/api/v1/hotels/previewSearch").Queries("pattern", "{pattern}").
		HandlerFunc(handler.FetchHotelsPreview).Methods("GET")

	r.Path("/api/v1/hotels/radiusSearch").Queries("latitude", "{latitude}",
		"longitude", "{longitude}", "radius", "{radius}").HandlerFunc(handler.FetchHotelsByRadius).Methods("GET")
}

// swagger:route GET /api/v1/hotels/radiusSearch hotel hotelsByRadius
// GetList of hotels by radius
// responses:
//  200: hotels
//  400: badrequest
func (hh *HotelHandler) FetchHotelsByRadius(w http.ResponseWriter, r *http.Request) {

	radius := r.FormValue("radius")
	latitude := r.FormValue("latitude")
	longitude := r.FormValue("longitude")

	hotels, err := hh.HotelUseCase.GetHotelsByRadius(latitude, longitude, radius)

	if err != nil {
		customerror.PostError(w, r, hh.log, err, nil)
		return
	}
	responses.SendDataResponse(w, hotelmodel.Hotels{hotels})
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
		customerror.PostError(w, r, hh.log, err, clientError.BadRequest)
		return
	}

	hotel, err := hh.HotelUseCase.GetHotels(startId)

	if err != nil {
		customerror.PostError(w, r, hh.log, err, nil)
		return
	}
	hotels := hotelmodel.Hotels{Hotels: hotel}
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
		customerror.PostError(w, r, hh.log, err, clientError.BadRequest)
		return
	}

	hotel, err := hh.HotelUseCase.GetHotelByID(id)

	if err != nil {
		customerror.PostError(w, r, hh.log, err, nil)
		return
	}
	data := hotelmodel.HotelData{Hotel: hotel}

	usr, ok := r.Context().Value(configs.RequestUser).(models.User)
	if !ok {
		responses.SendDataResponse(w, data)
		return
	}

	comment, err := hh.HotelUseCase.CheckRateExist(usr.ID, hotel.HotelID)
	if err != nil {
		responses.SendDataResponse(w, data)
		return
	}
	data.Comment = &comment

	responses.SendDataResponse(w, data)
}

// swagger:route GET /api/v1/hotels/search hotel searchHotel
// Search hotels
// responses:
//  200: searchHotel
//  400: badrequest
func (hh *HotelHandler) FetchHotels(w http.ResponseWriter, r *http.Request) {

	pattern := r.FormValue("pattern")
	pageNum := r.FormValue("page")
	page, err := strconv.Atoi(pageNum)
	if err != nil {
		customerror.PostError(w, r, hh.log, err, clientError.BadRequest)
		return
	}

	rateStart := r.FormValue("rateStart")
	commStart := r.FormValue("commentStart")

	radius := r.FormValue("radius")
	latitude := r.FormValue("latitude")
	longitude := r.FormValue("longitude")

	commCountConstraint := r.FormValue("commCount")
	commCountPercent := r.FormValue("commPercent")

	orderData := hotelmodel.HotelFiltering{rateStart, commStart,
		longitude, latitude, radius, commCountConstraint, commCountPercent}
	if err != nil {
		customerror.PostError(w, r, hh.log, err, clientError.BadRequest)
		return
	}

	hotels, err := hh.HotelUseCase.FetchHotels(orderData, pattern, page)

	if err != nil {
		customerror.PostError(w, r, hh.log, err, nil)
		return
	}

	responses.SendDataResponse(w, hotels)
}

// swagger:route GET /api/v1/hotels/previewSearch hotel hotelPreview
// Search hotels preview
// responses:
//  200: hotelsPreview
//  400: badrequest
func (hh *HotelHandler) FetchHotelsPreview(w http.ResponseWriter, r *http.Request) {
	pattern := r.FormValue("pattern")

	hotels, err := hh.HotelUseCase.GetHotelsPreview(pattern)

	if err != nil {
		customerror.PostError(w, r, hh.log, err, nil)
		return
	}

	responses.SendDataResponse(w, hotelmodel.HotelsPreview{hotels})
}
