//go:generate mockgen -source repository.go -destination mocks/comment_repository_mock.go -package mocks
package comment

import commModel "github.com/go-park-mail-ru/2020_2_JMickhs/internal/app/comment/models"

type Repository interface {
	GetComments(hotelID int, offset int) ([]commModel.FullCommentInfo, error)
	AddComment(comment commModel.Comment) (commModel.Comment, error)
	DeleteComment(ID int) error
	UpdateComment(comment *commModel.Comment) error
	UpdateHotelRating(hotelID int, NewRate float64) error
	GetCurrentRating(hotelID int) (commModel.RateInfo, error)
	CheckUser(comment *commModel.Comment) (int, error)
	GetCommentsCount(hotelID int) (int, error)
}
