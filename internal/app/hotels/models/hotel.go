package hotelmodel

import paginationModel "github.com/go-park-mail-ru/2020_2_JMickhs/internal/app/paginator/model"

type Hotel struct {
	HotelID     int      `json:"hotel_id" db:"hotel_id"`
	Name        string   `json:"name" db:"name"`
	Description string   `json:"description" db:"description"`
	Image       string   `json:"image" db:"img"`
	Location    string   `json:"location" db:"location"`
	Rating      float64  `json:"rating" db:"round"`
	Photos      []string `json:"photos,omitempty" db:"photos"`
	CommCount   int      `json:"comm_count" db:"comm_count"`
}

type HotelPreview struct {
	HotelID  int    `json:"hotel_id"`
	Name     string `json:"name"`
	Image    string `json:"image"`
	Location string `json:"location"`
}

type HotelData struct {
	Hotel    Hotel `json:"hotel"`
	CurrRate int   `json:"rate,omitempty"`
}

type SearchString struct {
	Pattern string `json:"pattern"`
}

type SearchData struct {
	Hotels  []Hotel                        `json:"hotels"`
	PagInfo paginationModel.PaginationInfo `json:"pag_info"`
}

type Cursor struct {
	NextCursor string `json:"nextcursor"`
	PrevCursor string `json:"prevcursor"`
}

type FilterData struct {
	Rating float64
	ID     string
}
