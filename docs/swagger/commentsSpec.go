package swagger

import "mime/multipart"

type NewRate struct {
	Rate    float64 `json:"new_rate" mapstructure:"new_rate"`
	Comment Comment `json:"comment" mapstructure:"comment"`
}

type PrevRate struct {
	Rate    int     `json:"prev_rate"`
	Comment Comment `json:"comment"`
}

type PaginationInfo struct {
	NextLink   string `json:"next,omitempty"`
	PrevLink   string `json:"prev,omitempty"`
	ItemsCount int    `json:"items_count,omitempty"`
}

type Comment struct {
	UserID  int      `json:"user_id" db:"user_id" mapstructure:"user_id"`
	HotelID int      `json:"hotel_id" mapstructure:"hotel_id"`
	CommID  int      `json:"comm_id" mapstructure:"comm_id"`
	Message string   `json:"message" mapstructure:"message"`
	Rate    float64  `json:"rating" mapstructure:"rating"`
	Time    string   `json:"time" mapstructure:"time"`
	Photos  []string `json:"photos"`
}

type FullCommentInfo struct {
	UserID   int      `json:"user_id" db:"user_id" mapstructure:"user_id"`
	CommID   int      `json:"comm_id" db:"comm_id" mapstructure:"comm_id"`
	HotelID  int      `json:"hotel_id" db:"hotel_id" mapstructure:"hotel_id"`
	Message  string   `json:"message" db:"message"`
	Rating   float64  `json:"rating" db:"rating"`
	Avatar   string   `json:"avatar" db:"concat"`
	Username string   `json:"username" db:"username"`
	Time     string   `json:"time" db:"time"`
	Photos   []string `json:"photos"`
}
type AddCommentFull struct {
	Add  AddComment       `json:"jsonData"`
	File []multipart.File `json:"photos"`
}

type UpdateCommentFull struct {
	Update UpdateComment    `json:"jsonData"`
	File   []multipart.File `json:"photos"`
}

type Photos struct {
	Photos []string `json:"photos"`
}

type AddComment struct {
	HotelID int    `json:"hotel_id"`
	Message string `json:"message"`
	Rating  int    `json:"rating"`
}

type UpdateComment struct {
	CommID  int    `json:"comm_id"`
	Message string `json:"message"`
	Rate    int    `json:"rating"`
}

type GetCommentsRequest struct {
	Comments []FullCommentInfo `json:"comments"`
	PagInfo  PaginationInfo    `json:"pag_info"`
}

// swagger:parameters comments
type CommentsIDParameterWrapper struct {
	// offset  start from 0
	// in: query
	// required:true
	Offset int `json:"offset"`
	// number of comments to show
	Limit int `json:"limit"`
	// the ID of hotel which comments you get
	// required:true
	ID int `json:"id"`
}

// swagger:parameters DeleteComment
type CommentsDeleteParameterWrapper struct {
	// the ID to delete comment
	// in: path
	// required:true
	ID int `json:"id"`
}

//swagger:response photos
type CommentsPhotos struct {
	//in: body
	Body Photos
}

//swagger:parameters Photos
type ParametersPhotos struct {
	//in: query
	ID int `json:"id"`
}

//swagger:response comments
type Comments struct {
	//in: body
	Body GetCommentsRequest
}

// swagger:response AddComment
type NewRateResponse struct {
	//in:body
	Body NewRate
}

// swagger:parameters UpdateComment
type UpdateCommentRequest struct {
	//in: body
	Body UpdateCommentFull
}

// swagger:response UpdateComment
type UpdateCommentResponse struct {
	//in: body
	Body NewRate
}

// swagger:parameters AddComment
type AddCommentRequest struct {
	//Это все в multipart
	//in: body
	Body AddCommentFull
}
