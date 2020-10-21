package commModel

type Comment struct {
	UserID  int     `json:"user_id"`
	HotelID int     `json:"hotel_id"`
	CommID  int     `json:"comm_id"`
	Message string  `json:"message"`
	Rate    float64 `json:"rate"`
	Time    string  `json:"time"`
}

//swagger:response commentInfo
type FullCommentInfo struct {
	UserID   int     `json:"user_id"`
	CommID   int     `json:"comm_id"`
	HotelID  int     `json:"hotel_id"`
	Message  string  `json:"message"`
	Rating   float64 `json:"rating"`
	Avatar   string  `json:"avatar"`
	Username string  `json:"username"`
	Time     string  `json:"time"`
}

type RateInfo struct {
	RatesCount int
	CurrRating float64
}

type NewRate struct {
	Rate    float64 `json:"new_rate"`
	Comment Comment `json:"comment"`
}

type PrevRate struct {
	Rate    int     `json:"prev_rate"`
	Comment Comment `json:"comment"`
}
