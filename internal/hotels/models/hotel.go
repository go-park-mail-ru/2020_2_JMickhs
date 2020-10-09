package models

// swagger:response Hotel
type Hotel struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description""`
	Image       string `json:"image""`
}

// swagger:parameters onehotel
type hotelIDParameterWrapper struct {
	// the id of hotel to get from database
	// in: path
	// required:true
	ID int `json:"id"`
}

// swagger:response listHotel
type hotelListWrapper struct{
	Body []Hotel
}
