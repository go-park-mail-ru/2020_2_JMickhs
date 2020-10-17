package swagger

import (
	hotelmodel "github.com/go-park-mail-ru/2020_2_JMickhs/internal/hotels/models"
)

// swagger:parameters searchHotel
type SearchStringRequest struct {
	// in: query
	Pattern string `json:"pattern"`
	Prev    string `json:"prev"`
	Next    string `json:"next"`
	// required:true
	Limit int `json:"limit"`
}

type RateRequest struct {
	HotelID int `json:"hotel_id"`
	UserID  int `json:"user_id"`
	Rate    int `json:"rate"`
}

type SearchDataResponse struct {
	Hotels []hotelmodel.Hotel `json:"hotels"`
	Cursor hotelmodel.Cursor  `json:"cursor"`
}

// swagger:parameters hotel
type hotelIDParameterWrapper struct {
	// the id of hotel to get from database
	// in: path
	// required:true
	ID int `json:"id"`
}

// swagger:response rates
type newRateResponse struct {
	//in:body
	Body hotelmodel.NewRate
}

// swagger:parameters rates
type newRateRequest struct {
	//in:body
	Body RateRequest
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
	Body []hotelmodel.Hotel
}

// swagger:response searchHotel
type searchHotelListWrapper struct {
	//in: body
	Body SearchDataResponse
}
