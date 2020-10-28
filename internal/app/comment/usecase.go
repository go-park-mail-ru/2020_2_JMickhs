//go:generate mockgen -source usecase.go -destination mocks/comment_usecase_mock.go -package comment_mock
package comment

import (
	commModel "github.com/go-park-mail-ru/2020_2_JMickhs/internal/app/comment/models"
	paginationModel "github.com/go-park-mail-ru/2020_2_JMickhs/internal/app/paginator/model"
)

type Usecase interface {
	GetComments(hotelID int, page int) (paginationModel.PaginationModel, error)
	AddComment(comment commModel.Comment) (commModel.NewRate, error)
	DeleteComment(ID int) error
	UpdateComment(comment commModel.Comment) (commModel.NewRate, error)
	UpdateRating(prevRate commModel.PrevRate) (float64, error)
	AddRating(comment commModel.Comment) (float64, error)
}
