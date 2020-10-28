package commModel

// easyjson -all comment.go

// easyjson:json
type Comment struct {
	UserID  int     `json:"user_id" db:"user_id"`
	HotelID int     `json:"hotel_id"`
	CommID  int     `json:"comm_id"`
	Message string  `json:"message"`
	Rate    float64 `json:"rate"`
	Time    string  `json:"time"`
}

// easyjson:json
//swagger:response commentInfo
type FullCommentInfo struct {
	UserID   int     `json:"user_id" db:"user_id"`
	CommID   int     `json:"comm_id" db:"comm_id"`
	HotelID  int     `json:"hotel_id" db:"hotel_id"`
	Message  string  `json:"message" db:"message"`
	Rating   float64 `json:"rating" db:"rating"`
	Avatar   string  `json:"avatar" db:"avatar"`
	Username string  `json:"username" db:"username"`
	Time     string  `json:"time" db:"time"`
}

// easyjson:json
type RateInfo struct {
	RatesCount int
	CurrRating float64
}

// easyjson:json
type NewRate struct {
	Rate    float64 `json:"new_rate"`
	Comment Comment `json:"comment"`
}

// easyjson:json
type PrevRate struct {
	Rate    int     `json:"prev_rate"`
	Comment Comment `json:"comment"`
}
