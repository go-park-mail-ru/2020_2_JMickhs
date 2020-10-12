package swagger

import "github.com/go-park-mail-ru/2020_2_JMickhs/internal/comment/models"

type AddCommentRequest struct {
	UserID  int    `json:"user_id"`
	HotelID int    `json:"hotel_id"`
	Message string `json:"message"`
	Rating  int    `json:"rating"`
}

type UpdateCommentRequest struct {
	CommID  int    `json:"comm_id"`
	Message string `json:"message"`
}

// swagger:parameters comments
type CommentsIDParameterWrapper struct {
	// the start ID to get comments
	// in: query
	// required:true
	From int `json:"from"`
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

//swagger:response comments
type Comments struct {
	//in: body
	Body []models.FullCommentInfo
}

// swagger:parameters UpdateComment
type updateCommentRequest struct {
	//in: body
	Body UpdateCommentRequest
}

// swagger:response AddComment
type commentResponse struct {
	//in: body
	Body models.Comment
}

// swagger:parameters AddComment
type addCommentRequest struct {
	//in: body
	Body AddCommentRequest
}
