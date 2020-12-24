package hotelDelivery

import (
	"net/http"
	"strconv"

	recommModels "github.com/go-park-mail-ru/2020_2_JMickhs/main/internal/app/recommendation/models"

	"github.com/go-park-mail-ru/2020_2_JMickhs/main/internal/app/recommendation"

	packageConfig "github.com/go-park-mail-ru/2020_2_JMickhs/package/configs"

	"github.com/go-park-mail-ru/2020_2_JMickhs/main/internal/app/hotels"
	hotelmodel "github.com/go-park-mail-ru/2020_2_JMickhs/main/internal/app/hotels/models"

	customerror "github.com/go-park-mail-ru/2020_2_JMickhs/package/error"

	"github.com/go-park-mail-ru/2020_2_JMickhs/package/clientError"
	"github.com/go-park-mail-ru/2020_2_JMickhs/package/logger"
	"github.com/go-park-mail-ru/2020_2_JMickhs/package/responses"

	"github.com/gorilla/mux"
)

type HotelHandler struct {
	HotelUseCase          hotels.Usecase
	RecommendationUseCase recommendation.Usecase
	log                   *logger.CustomLogger
}

func NewHotelHandler(r *mux.Router, hs hotels.Usecase, RecommendationUseCase recommendation.Usecase, lg *logger.CustomLogger) {
	handler := HotelHandler{
		HotelUseCase:          hs,
		log:                   lg,
		RecommendationUseCase: RecommendationUseCase,
	}
	r.HandleFunc("/api/v1/hotels/{id:[0-9]+}", handler.Hotel).Methods("GET")
	r.Path("/api/v1/hotels/search").Queries("pattern", "{pattern}", "page", "{page:[0-9]+}").
		HandlerFunc(handler.FetchHotels).Methods("GET")
	r.Path("/api/v1/hotels/previewSearch").Queries("pattern", "{pattern}").
		HandlerFunc(handler.FetchHotelsPreview).Methods("GET")

	r.Path("/api/v1/hotels/radiusSearch").Queries("latitude", "{latitude}",
		"longitude", "{longitude}", "radius", "{radius}").HandlerFunc(handler.FetchHotelsByRadius).Methods("GET")

	r.HandleFunc("/api/v1/hotels/recommendations", handler.Recommendations).Methods("GET")
}

// swagger:route GET /api/v1/hotels/recommendations hotel hotelsRecommendation
// GetUserRecommendations
// responses:
//  200: recommendations
func (hh *HotelHandler) Recommendations(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(packageConfig.RequestUserID).(int)
	if !ok {
		userID = -1
	}

	hotels, err := hh.RecommendationUseCase.GetHotelsRecommendations(userID)
	if err != nil {
		customerror.PostError(w, r, hh.log, err, nil)
		return
	}

	responses.SendDataResponse(w, recommModels.Hotels{Hotels: hotels})
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

	responses.SendDataResponse(w, hotelmodel.Hotels{Hotels: hotels})
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

	hotel := hotelmodel.Hotel{}
	userID, ok := r.Context().Value(packageConfig.RequestUserID).(int)
	if !ok {
		userID = -1
	}
	hotel, err = hh.HotelUseCase.GetHotelByID(id, userID)
	if err != nil {
		customerror.PostError(w, r, hh.log, err, nil)
		return
	}

	data := hotelmodel.HotelData{Hotel: hotel}

	comment, err := hh.HotelUseCase.CheckRateExist(userID, hotel.HotelID)
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
	rateEnd := r.FormValue("rateEnd")
	commStart := r.FormValue("commentStart")

	radius := r.FormValue("radius")
	latitude := r.FormValue("latitude")
	longitude := r.FormValue("longitude")

	commCountConstraint := r.FormValue("commCount")
	commCountPercent := r.FormValue("commPercent")

	orderData := hotelmodel.HotelFiltering{RatingFilterStartNumber: rateStart, RatingFilterEndNumber: rateEnd, CommentsFilterStartNumber: commStart,
		Longitude: longitude, Latitude: latitude, Radius: radius, CommCountConstraint: commCountConstraint, CommCountPercent: commCountPercent}

	userID, ok := r.Context().Value(packageConfig.RequestUserID).(int)
	if !ok {
		userID = -1
	}
	if userID != -1 && pattern != "" {
		err := hh.RecommendationUseCase.AddInSearchHistory(userID, pattern)
		if err != nil {
			customerror.PostError(w, r, hh.log, err, nil)
			return
		}

	}
	hotels, err := hh.HotelUseCase.FetchHotels(orderData, pattern, page, userID)

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

	responses.SendDataResponse(w, hotelmodel.HotelsPreview{Hotels: hotels})
}
