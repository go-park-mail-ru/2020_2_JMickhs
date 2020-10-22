package swagger

import (
	commModel "github.com/go-park-mail-ru/2020_2_JMickhs/internal/app/comment/models"
	paginationModel "github.com/go-park-mail-ru/2020_2_JMickhs/internal/app/paginator/model"
)

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

type GetCommentsRequest struct {
	Comments []commModel.FullCommentInfo    `json:"comments"`
	PagInfo  paginationModel.PaginationInfo `json:"Pag_info"`
}

// swagger:parameters comments
type CommentsIDParameterWrapper struct {
	// page num start from 0
	// in: query
	// required:true
	Page int `json:"page"`
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
	Body GetCommentsRequest
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
