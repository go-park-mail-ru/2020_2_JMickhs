//go:generate  easyjson -all comment.go
package commModel

import paginationModel "github.com/go-park-mail-ru/2020_2_JMickhs/main/internal/app/paginator/model"

// easyjson:json
type Comment struct {
	UserID  int     `json:"user_id" db:"user_id" mapstructure:"user_id"`
	HotelID int     `json:"hotel_id" mapstructure:"hotel_id"`
	CommID  int     `json:"comm_id" mapstructure:"comm_id"`
	Message string  `json:"message" mapstructure:"message"`
	Rate    float64 `json:"rating" mapstructure:"rating"`
	Time    string  `json:"time" mapstructure:"time"`
}

// easyjson:json
//swagger:response commentInfo
type FullCommentInfo struct {
	UserID   int     `json:"user_id" db:"user_id" mapstructure:"user_id"`
	CommID   int     `json:"comm_id" db:"comm_id" mapstructure:"comm_id"`
	HotelID  int     `json:"hotel_id" db:"hotel_id" mapstructure:"hotel_id"`
	Message  string  `json:"message" db:"message"`
	Rating   float64 `json:"rating" db:"rating"`
	Avatar   string  `json:"avatar" db:"concat"`
	Username string  `json:"username" db:"username"`
	Time     string  `json:"time" db:"time"`
}

// easyjson:json
type RateInfo struct {
	RatesCount int
	CurrRating float64
}

// easyjson:json
type CommentWithAlbum struct {
	UserID  int      `json:"user_id" db:"user_id" mapstructure:"user_id"`
	HotelID int      `json:"hotel_id" mapstructure:"hotel_id"`
	CommID  int      `json:"comm_id" mapstructure:"comm_id"`
	Message string   `json:"message" mapstructure:"message"`
	Rate    float64  `json:"rating" mapstructure:"rating"`
	Time    string   `json:"time" mapstructure:"time"`
	Photos  []string `json:"photos,omitempty" db:"photos"`
}

// easyjson:json
type Comments struct {
	Comments []FullCommentInfo              `json:"comments"`
	Info     paginationModel.PaginationInfo `json:"pag_info"`
}

// easyjson:json
type NewRate struct {
	Rate    float64 `json:"new_rate" mapstructure:"new_rate"`
	Comment Comment `json:"comment" mapstructure:"comment"`
}

// easyjson:json
type PrevRate struct {
	Rate    int     `json:"prev_rate"`
	Comment Comment `json:"comment"`
}
