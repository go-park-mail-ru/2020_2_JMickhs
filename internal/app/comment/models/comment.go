package models

type Comment struct {
	UserID  int    `json:"user_id"`
	HotelID int    `json:"hotel_id"`
	CommID  int    `json:"comm_id"`
	Message string `json:"message"`
	Rating  int    `json:"rating"`
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
}
