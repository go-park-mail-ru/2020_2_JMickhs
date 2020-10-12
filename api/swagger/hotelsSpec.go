package swagger

import "github.com/go-park-mail-ru/2020_2_JMickhs/internal/hotels/models"


// swagger:parameters searchHotel
type SearchStringRequest struct{
	// in: query
	// required:true
	Pattern string `json:"pattern"`
	// required:true
	From int `json:"from"`
	// required:true
	Limit int `json:"limit"`
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
type hotelListWrapper struct{
	//in: body
	Body []models.Hotel
}

// swagger:response searchHotel
type searchHotelListWrapper struct{
	//in: body
	Body []models.Hotel
}