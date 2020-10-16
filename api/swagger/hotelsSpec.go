package swagger

import "github.com/go-park-mail-ru/2020_2_JMickhs/internal/hotels/models"

// swagger:parameters searchHotel
type SearchStringRequest struct {
	// in: query
	Pattern string `json:"pattern"`
	Prev    string `json:"prev"`
	Next    string `json:"next"`
	// required:true
	Limit int `json:"limit"`
}

type SearchDataResponse struct {
	Hotels []models.Hotel `json:"hotels"`
	Cursor models.Cursor  `json:"cursor"`
}

// swagger:parameters hotel
type hotelIDParameterWrapper struct {
	// the id of hotel to get from database
	// in: path
	// required:true
	ID int `json:"id"`
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
	Body []models.Hotel
}

// swagger:response searchHotel
type searchHotelListWrapper struct {
	//in: body
	Body SearchDataResponse
}
