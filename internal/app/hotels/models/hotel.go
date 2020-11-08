//go:generate easyjson -all hotel.go
package hotelmodel

import (
	commModel "github.com/go-park-mail-ru/2020_2_JMickhs/internal/app/comment/models"
	paginationModel "github.com/go-park-mail-ru/2020_2_JMickhs/internal/app/paginator/model"
)

// easyjson:json
type Hotel struct {
	HotelID     int      `json:"hotel_id" db:"hotel_id" mapstructure:"hotel_id"`
	Name        string   `json:"name" db:"name"`
	Description string   `json:"description" db:"description"`
	Image       string   `json:"image" db:"concat"`
	Location    string   `json:"location" db:"location"`
	Rating      float64  `json:"rating" db:"curr_rating"`
	Photos      []string `json:"photos,omitempty" db:"photos" faker:"len=50"`
	CommCount   int      `json:"comm_count" db:"comm_count" mapstructure:"comm_count"`
}

// easyjson:json
type HotelPreview struct {
	HotelID  int    `json:"hotel_id" db:"hotel_id"`
	Name     string `json:"name" db:"name"`
	Image    string `json:"image" db:"concat"`
	Location string `json:"location" db:"location"`
}

// easyjson:json
type HotelData struct {
	Hotel    Hotel `json:"hotel"`
	Comment  *commModel.FullCommentInfo  `json:"comment,omitempty" mapstructure:"currate"`
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
type Hotels struct{
	Hotels []Hotel `json:"hotels"`
}

// easyjson:json
type HotelsPreview struct{
	Hotels []HotelPreview`json:"hotels_preview"`
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
