//go:generate mockgen -source repository.go -destination mocks/comment_repository_mock.go -package comment_mock
package comment

import commModel "github.com/go-park-mail-ru/2020_2_JMickhs/internal/app/comment/models"

type Repository interface {
	GetComments(hotelID string, limit int, offset string,user_id int) ([]commModel.FullCommentInfo, error)
	AddComment(comment commModel.Comment) (commModel.Comment, error)
	DeleteComment(ID int) error
	UpdateComment(comment *commModel.Comment) error
	UpdateHotelRating(hotelID int, NewRate float64) error
	GetCurrentRating(hotelID int) (commModel.RateInfo, error)
	CheckUser(comment *commModel.Comment) (int, error)
	CheckRateExistForComments(hotelID int, userID int) (bool,error)
	GetCommentsCount(hotelID int) (int, error)
}
