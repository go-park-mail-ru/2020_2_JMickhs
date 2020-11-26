//go:generate mockgen -source usecase.go -destination mocks/comment_usecase_mock.go -package comment_mock
package comment

import (
	commModel "github.com/go-park-mail-ru/2020_2_JMickhs/main/internal/app/comment/models"
)

type Usecase interface {
	GetComments(hotelID string, limit string, offset string, user_id int) (int, commModel.Comments, error)
	AddComment(comment commModel.Comment) (commModel.NewRate, error)
	DeleteComment(ID int) error
	UpdateComment(comment commModel.Comment) (commModel.NewRate, error)
	UpdateRating(prevRate commModel.PrevRate) (float64, error)
	AddRating(comment commModel.Comment) (float64, error)
}
