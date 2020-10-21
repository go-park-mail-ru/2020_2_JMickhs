package swagger

import (
	hotelmodel "github.com/go-park-mail-ru/2020_2_JMickhs/internal/app/hotels/models"
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
	Hotels []hotelmodel.Hotel `json:"hotels"`
	Cursor hotelmodel.Cursor  `json:"cursor"`
}

type HotelData struct {
	Hotel    hotelmodel.Hotel `json:"hotel"`
	CurrRate int              `json:"rate"`
}

// swagger:parameters hotel
type hotelIDParameterWrapper struct {
	// the id of hotel to get from database
	// in: path
	// required:true
	ID int `json:"id"`
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
	Body []PreviewHotel
}

// swagger:response searchHotel
type searchHotelListWrapper struct {
	//in: body
	Body SearchDataResponse
}
