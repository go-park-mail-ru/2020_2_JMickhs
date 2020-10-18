package commModel

type Comment struct {
	UserID  int    `json:"user_id"`
	HotelID int    `json:"hotel_id"`
	CommID  int    `json:"comm_id"`
	Message string `json:"message"`
	Rate    int    `json:"rate"`
	Time    string `json:"time"`
}

//swagger:response commentInfo
type FullCommentInfo struct {
	UserID   int    `json:"user_id"`
	CommID   int    `json:"comm_id"`
	HotelID  int    `json:"hotel_id"`
	Message  string `json:"message"`
	Rating   int    `json:"rating"`
	Avatar   string `json:"avatar"`
	Username string `json:"username"`
	Time     string `json:"time"`
}

type Rating struct {
	RateID  int `json:"rate_id"`
	HotelID int `json:"hotel_id"`
	UserID  int `json:"user_id"`
	Rate    int `json:"rate"`
}

type RateInfo struct {
	RatesCount int
	CurrRating int
}
type NewRate struct {
	Rate    int     `json:"new_rate"`
	Comment Comment `json:"comment"`
}
