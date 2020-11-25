package hotelDelivery

import (
	"errors"
	"mime/multipart"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-park-mail-ru/2020_2_JMickhs/main/configs"
	"github.com/go-park-mail-ru/2020_2_JMickhs/main/internal/app/hotels"
	hotelmodel "github.com/go-park-mail-ru/2020_2_JMickhs/main/internal/app/hotels/models"

	"github.com/mailru/easyjson"

	"github.com/spf13/viper"

	customerror "github.com/go-park-mail-ru/2020_2_JMickhs/package/error"

	"github.com/go-park-mail-ru/2020_2_JMickhs/package/clientError"
	"github.com/go-park-mail-ru/2020_2_JMickhs/package/logger"
	"github.com/go-park-mail-ru/2020_2_JMickhs/package/responses"

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

	r.Path("/api/v1/hotels").HandlerFunc(handler.AddHotelByOwner).Methods("POST")
}

func (hh *HotelHandler) DetectFileContentType(file multipart.File) (string, error) {
	fileHeader := make([]byte, 512)
	contentType := ""
	if _, err := file.Read(fileHeader); err != nil {
		return contentType, customerror.NewCustomError(err, clientError.BadRequest, 1)
	}

	if _, err := file.Seek(0, 0); err != nil {
		return contentType, customerror.NewCustomError(err, clientError.BadRequest, 1)
	}
	contentTypeStr := http.DetectContentType(fileHeader)
	contentType = strings.Split(contentTypeStr, "/")[1]
	if contentType != "jpg" && contentType != "png" && contentType != "jpeg" {
		return contentType, customerror.NewCustomError(errors.New("Wrong file type"), clientError.UnsupportedMediaType, 1)
	}
	return contentType, nil
}

// swagger:route GET /api/v1/hotels hotel AddHotel
// AddHotel
// responses:
//  400: badrequest
func (hh *HotelHandler) AddHotelByOwner(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(viper.GetString(configs.ConfigFields.RequestUserID)).(int)
	if !ok {
		customerror.PostError(w, r, hh.log, errors.New("Unauthorized"), clientError.Unauthorizied)
		return
	}

	hotel := hotelmodel.Hotel{}
	r.ParseMultipartForm(100 * configs.MB)

	r.Body = http.MaxBytesReader(w, r.Body, 100*configs.MB)

	err := easyjson.Unmarshal([]byte(r.FormValue("jsonData")), &hotel)
	if err != nil {
		customerror.PostError(w, r, hh.log, err, clientError.BadRequest)
		return
	}

	mainImage, _, err := r.FormFile("mainImage")
	if err != nil {
		customerror.PostError(w, r, hh.log, err, clientError.BadRequest)
		return
	}

	fileType, err := hh.DetectFileContentType(mainImage)
	if err != nil {
		customerror.PostError(w, r, hh.log, err, nil)
		return
	}

	err = hh.HotelUseCase.UploadPhoto(&hotel, mainImage, fileType, true, 0)
	if err != nil {
		customerror.PostError(w, r, hh.log, err, nil)
		return
	}

	photos := r.MultipartForm.File["photos"]
	for iterator, photo := range photos {
		file, err := photo.Open()
		if err != nil {
			customerror.PostError(w, r, hh.log, err, clientError.BadRequest)
			return
		}
		defer file.Close()
		fileType, err := hh.DetectFileContentType(file)
		if err != nil {
			customerror.PostError(w, r, hh.log, err, nil)
			return
		}
		err = hh.HotelUseCase.UploadPhoto(&hotel, mainImage, fileType, false, iterator)
		if err != nil {
			customerror.PostError(w, r, hh.log, err, nil)
			return
		}
	}

	err = hh.HotelUseCase.AddHotel(hotel, userID)
	if err != nil {
		customerror.PostError(w, r, hh.log, err, nil)
		return
	}

	responses.SendOkResponse(w)
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

	userID, ok := r.Context().Value(viper.GetString(configs.ConfigFields.RequestUserID)).(int)
	if !ok {
		responses.SendDataResponse(w, data)
		return
	}

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

	orderData := hotelmodel.HotelFiltering{rateStart, rateEnd, commStart,
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
