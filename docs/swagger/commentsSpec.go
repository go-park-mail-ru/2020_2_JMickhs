package swagger

import commModel "github.com/go-park-mail-ru/2020_2_JMickhs/internal/app/comment/models"

type AddCommentRequest struct {
	HotelID int    `json:"hotel_id"`
	Message string `json:"message"`
	Rating  int    `json:"rate"`
}

type UpdateCommentRequest struct {
	CommID  int    `json:"comm_id"`
	Message string `json:"message"`
	Rate    int    `json:"rate"`
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
	Body []commModel.FullCommentInfo
}

// swagger:response AddComment
type newRateResponse struct {
	//in:body
	Body commModel.NewRate
}

// swagger:parameters UpdateComment
type updateCommentRequest struct {
	//in: body
	Body UpdateCommentRequest
}

// swagger:response UpdateComment
type updateCommentResponse struct {
	//in: body
	Body commModel.NewRate
}

// swagger:parameters AddComment
type addCommentRequest struct {
	//in: body
	Body AddCommentRequest
}
