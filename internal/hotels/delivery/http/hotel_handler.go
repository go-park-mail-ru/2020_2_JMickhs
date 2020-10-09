package hotelDelivery

import (
	"encoding/json"
	permissions "github.com/go-park-mail-ru/2020_2_JMickhs/internal/permission"
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

	r.HandleFunc("/api/v1/hotels",  permissions.SetCSRF(handler.ListHotels)).Methods("GET")
	r.HandleFunc("/api/v1/hotel/{id:[0-9]+}",  permissions.SetCSRF(handler.Hotel)).Methods("GET")
}

// swagger:route GET /api/v1/hotels hotels listHotel
// GetList of hotels
// responses:
//  200: listHotel
func (hh *HotelHandler) ListHotels(w http.ResponseWriter, r *http.Request) {

	hotels, err := hh.HotelUseCase.GetHotels()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(hotels)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// swagger:route GET /api/v1/hotel/{id} hotels onehotel
// Get single hotel by id
// responses:
//  200: Hotel
func (hh *HotelHandler) Hotel(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		hh.log.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	hotel, err := hh.HotelUseCase.GetHotelByID(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	err = json.NewEncoder(w).Encode(hotel)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
