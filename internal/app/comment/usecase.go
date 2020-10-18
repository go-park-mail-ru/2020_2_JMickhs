//go:generate mockgen -source usecase.go -destination mocks/comment_usecase_mock.go -package mocks
package comment

import commModel "github.com/go-park-mail-ru/2020_2_JMickhs/internal/app/comment/models"

type Usecase interface {
	GetComments(hotelID int, StartID int) ([]commModel.FullCommentInfo, error)
	AddComment(comment commModel.Comment) (commModel.NewRate, error)
	DeleteComment(ID int) error
	UpdateComment(comment commModel.Comment) error
	UpdateRating(rating commModel.Rating) (int, error)
}
