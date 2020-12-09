//go:generate mockgen -source usecase.go -destination mocks/comment_usecase_mock.go -package comment_mock
package comment

import (
	"mime/multipart"

	commModel "github.com/go-park-mail-ru/2020_2_JMickhs/main/internal/app/comment/models"
)

type Usecase interface {
	UploadPhoto(comment *commModel.CommentWithAlbum, file multipart.File, contentType string, mainImage bool, iterator int) error
	AddCommentWithAlbum(comment commModel.CommentWithAlbum, userID int) error
	GetComments(hotelID string, limit string, offset string, user_id int) (int, commModel.Comments, error)
	AddComment(comment commModel.Comment) (commModel.NewRate, error)
	DeleteComment(ID int) error
	UpdateComment(comment commModel.Comment) (commModel.NewRate, error)
	UpdateRating(prevRate commModel.PrevRate) (float64, error)
	AddRating(comment commModel.Comment) (float64, error)
}
