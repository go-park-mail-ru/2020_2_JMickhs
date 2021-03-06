package swagger

import (
	"os"
)

type Hotel struct {
	HotelID     int      `json:"hotel_id" db:"hotel_id" mapstructure:"hotel_id"`
	Name        string   `json:"name" db:"name"`
	Description string   `json:"description" db:"description"`
	Image       string   `json:"image" db:"concat"`
	Location    string   `json:"location" db:"location"`
	Email       string   `json:"email" db:"email"`
	Country     string   `json:"country" db:"country"`
	City        string   `json:"city"   db:"city"`
	Rating      float64  `json:"rating" db:"curr_rating"`
	Photos      []string `json:"photos,omitempty" db:"photos" faker:"len=50"`
	CommCount   int      `json:"comm_count" db:"comm_count" mapstructure:"comm_count"`
	Latitude    float64  `json:"latitude,omitempty" db:"x"`
	Longitude   float64  `json:"longitude,omitempty" db:"y"`
	//enum "wishListIn" "wishListOut"
	WishListExist string `json:"wishlist_exist"`
}

type HotelPreview struct {
	HotelID  int    `json:"hotel_id" db:"hotel_id" mapstructure:"hotel_id"`
	Name     string `json:"name" db:"name"`
	Image    string `json:"image" db:"concat"`
	Location string `json:"location" db:"location"`
}

type Hotels struct {
	Hotels []Hotel `json:"hotels"`
}

type HotelsPreview struct {
	Hotels []HotelPreview `json:"hotels_preview"`
}

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
	RateEnd int `json:"rateEnd"`
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

type HotelRecommend struct {
	HotelID  int    `json:"hotel_id"`
	Name     string `json:"name"`
	Image    string `json:"image"`
	Location string `json:"location"`
	Rating   string `json:"rating"`
}

type SearchDataResponse struct {
	Hotels  []Hotel        `json:"hotels"`
	PagInfo PaginationInfo `json:"Pag_info"`
}

type RecommendationResponse struct {
	Hotels []HotelRecommend `json:"hotels"`
}

type HotelData struct {
	Hotel   Hotel           `json:"hotel"`
	Comment FullCommentInfo `json:"comment,omitempty"`
}

// swagger:response recommendations
type HotelsRecommendation struct {
	//in: body
	Hotels RecommendationResponse
}

// swagger:parameters hotel hotelsByRadius
type HotelByRadiusParameterWrapper struct {
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

type AddHotelStruct struct {
	JsonData  HotelStructToAdd `json:"jsonData"`
	MainImage os.File          `json:"mainImage"`
	Photos    []os.File        `json:"photos"`
}

type HotelStructToAdd struct {
	Name        string `json:"name"`
	County      string `json:"country"`
	City        string `json:"city"`
	Description string `json:"description"`
	Location    string `json:"location"`
}

// swagger:parameters AddHotel
type AddHotelResponse struct {
	//in: body
	Body AddHotelStruct
}

// swagger:parameters hotel
type HotelIDParameterWrapper struct {
	// the id of hotel to get from database
	// in: path
	// required:true
	ID int `json:"id"`
}

// swagger:parameters hotelsPreview
type HotelsPreviewRequest struct {
	// in: query
	// required:true
	ID int `json:"pattern"`
}

// swagger:response HotelData
type HotelResponseWrapper struct {
	//in: body
	Body HotelData
}

// swagger:parameters hotels
type HotelsIDParameterWrapper struct {
	// the start ID to get hotels
	// in: query
	// required:true
	ID int `json:"from"`
}

// swagger:response hotels
type HotelListWrapper struct {
	//in: body
	Body Hotels
}

// swagger:response hotelsPreview
type HotelsPreviewResponse struct {
	//in: body
	Body HotelsPreview
}

// swagger:response searchHotel
type SearchHotelListWrapper struct {
	//in: body
	Body SearchDataResponse
}
