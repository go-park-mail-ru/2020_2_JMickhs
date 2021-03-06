//go:generate easyjson -all hotel.go
package hotelmodel

import (
	commModel "github.com/go-park-mail-ru/2020_2_JMickhs/main/internal/app/comment/models"
	paginationModel "github.com/go-park-mail-ru/2020_2_JMickhs/main/internal/app/paginator/model"
)

// easyjson:json
type Hotel struct {
	HotelID       int      `json:"hotel_id" db:"hotel_id" mapstructure:"hotel_id"`
	Name          string   `json:"name" db:"name"`
	Description   string   `json:"description,omitempty" db:"description"`
	Image         string   `json:"image" db:"concat"`
	Location      string   `json:"location" db:"location"`
	Email         string   `json:"email,omitempty" db:"email"`
	Country       string   `json:"country,omitempty" db:"country"`
	City          string   `json:"city,omitempty"   db:"city"`
	Rating        float64  `json:"rating" db:"curr_rating"`
	Photos        []string `json:"photos,omitempty" db:"photos" faker:"len=50"`
	CommCount     int      `json:"comm_count,omitempty" db:"comm_count" mapstructure:"comm_count"`
	Latitude      float64  `json:"latitude,omitempty" db:"x"`
	Longitude     float64  `json:"longitude,omitempty" db:"y"`
	WishListExist string   `json:"wishlist_exist,omitempty"`
}

// easyjson:json
type HotelPreview struct {
	HotelID  int    `json:"hotel_id" db:"hotel_id" mapstructure:"hotel_id"`
	Name     string `json:"name" db:"name"`
	Image    string `json:"image" db:"concat"`
	Location string `json:"location" db:"location"`
}

type HotelFiltering struct {
	RatingFilterStartNumber   string
	RatingFilterEndNumber     string
	CommentsFilterStartNumber string
	Longitude                 string
	Latitude                  string
	Radius                    string
	CommCountConstraint       string
	CommCountPercent          string
}

// easyjson:json
type HotelData struct {
	Hotel   Hotel                      `json:"hotel"`
	Comment *commModel.FullCommentInfo `json:"comment,omitempty" mapstructure:"currate"`
}

// easyjson:json
type SearchString struct {
	Pattern string `json:"pattern"`
}

// easyjson:json
type SearchData struct {
	Hotels  []Hotel                        `json:"hotels"`
	PagInfo paginationModel.PaginationInfo `json:"pag_info"`
}

// easyjson:json
type Hotels struct {
	Hotels []Hotel `json:"hotels"`
}

// easyjson:json
type HotelsPreview struct {
	Hotels []HotelPreview `json:"hotels_preview"`
}

// easyjson:json
type Cursor struct {
	NextCursor string `json:"nextcursor"`
	PrevCursor string `json:"prevcursor"`
}

// easyjson:json
type FilterData struct {
	Rating float64
	ID     string
}

// easyjson:json
type MiniHotel struct {
	HotelID  int     `json:"hotel_id" db:"hotel_id"  mapstructure:"hotel_id"`
	Name     string  `json:"name" db:"name"`
	Image    string  `json:"image" db:"img"`
	Location string  `json:"location" db:"location"`
	Rating   float64 `json:"rating" db:"curr_rating"`
}
