//go:generate mockgen -source usecase.go -destination mocks/comment_usecase_mock.go -package mocks
package comment

import "github.com/go-park-mail-ru/2020_2_JMickhs/internal/app/comment/models"

type Usecase interface {
	GetComments(hotelID int, StartID int) ([]models.FullCommentInfo, error)
	AddComment(comment models.Comment) (models.Comment, error)
	DeleteComment(ID int) error
	UpdateComment(comment models.Comment) error
}
