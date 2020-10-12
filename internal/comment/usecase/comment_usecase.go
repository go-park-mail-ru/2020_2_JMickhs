package commentUsecase

import (
	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/comment"
	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/comment/models"
)

type CommentUseCase struct {
	commentRepo     comment.Repository
}

func NewCommentUsecase(r comment.Repository) *CommentUseCase {
	return &CommentUseCase{
		commentRepo :     r,
	}
}


func(u *CommentUseCase) GetComments(hotelID int,StartID int) ([]models.FullCommentInfo, error){
	return u.commentRepo.GetComments(hotelID,StartID)
}
func(u *CommentUseCase) AddComment(comment models.Comment) (models.Comment, error){
	return u.commentRepo.AddComment(comment)
}
func(u *CommentUseCase) DeleteComment(ID int) error{
	return u.commentRepo.DeleteComment(ID)
}
func(u *CommentUseCase) UpdateComment(comment models.Comment) (error){
	return u.commentRepo.UpdateComment(comment)
}