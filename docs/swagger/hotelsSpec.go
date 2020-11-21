package swagger

import (
	commModel "github.com/go-park-mail-ru/2020_2_JMickhs/main/internal/app/comment/models"
	hotelmodel "github.com/go-park-mail-ru/2020_2_JMickhs/main/internal/app/hotels/models"
	paginationModel "github.com/go-park-mail-ru/2020_2_JMickhs/main/internal/app/paginator/model"
)

// swagger:parameters searchHotel
type SearchStringRequest struct {
	// required:true
	Pattern string `json:"pattern"`
	// page num start from 0
	// required:true
	Page string `json:"page"`
	// в метрах
	// in: query
	Radius string `json:"radius"`
	// широта
	// in: query
	Latitude string `json:"latitude"`
	// долгота
	// in: query
	Longitude string `json:"longitude"`
	//in:query
	RateStart int `json:"rateStart"`
	//in:query
	CommentCountStart int `json:"commentStart"`
	// какие оценки выбираем через запятую пример "3,4,5"
	//in:query
	CommCountConstraint int `json:"commCount"`
	//процент выбранных оценок от общего кол-ва
	//in:query
	CommCountPercent int `json:"commPercent"`
}

// swagger:parameters hotelPreview
type SearchPreviewStringRequest struct {
	// page num start from 0
	// required:true
	Pattern string `json:"pattern"`
}

type PreviewHotel struct {
	HotelID     int    `json:"hotel_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Location    string `json:"location"`
	Rating      string `json:"rating"`
}
type RateRequest struct {
	HotelID int `json:"hotel_id"`
	UserID  int `json:"user_id"`
	Rate    int `json:"rate"`
}

type SearchDataResponse struct {
	Hotels  []hotelmodel.Hotel             `json:"hotels"`
	PagInfo paginationModel.PaginationInfo `json:"Pag_info"`
}

type HotelData struct {
	Hotel   hotelmodel.Hotel          `json:"hotel"`
	Comment commModel.FullCommentInfo `json:"comment,omitempty"`
}

// swagger:parameters hotel hotelsByRadius
type hotelByRadiusParameterWrapper struct {
	// в метрах
	// in: query
	// required:true
	Radius string `json:"radius"`
	// широта
	// in: query
	// required:true
	Latitude string `json:"latitude"`
	// долгота
	// in: query
	// required:true
	Longitude string `json:"longitude"`
}

// swagger:parameters hotel
type hotelIDParameterWrapper struct {
	// the id of hotel to get from database
	// in: path
	// required:true
	ID int `json:"id"`
}

// swagger:parameters hotelsPreview
type hotelsPreviewRequest struct {
	// in: query
	// required:true
	ID int `json:"pattern"`
}

// swagger:response HotelData
type hotelResponseWrapper struct {
	//in: body
	Body HotelData
}

// swagger:parameters hotels
type hotelsIDParameterWrapper struct {
	// the start ID to get hotels
	// in: query
	// required:true
	ID int `json:"from"`
}

// swagger:response hotels
type hotelListWrapper struct {
	//in: body
	Body hotelmodel.Hotels
}

// swagger:response hotelsPreview
type hotelsPreviewResponse struct {
	//in: body
	Body hotelmodel.HotelsPreview
}

// swagger:response searchHotel
type searchHotelListWrapper struct {
	//in: body
	Body SearchDataResponse
}
