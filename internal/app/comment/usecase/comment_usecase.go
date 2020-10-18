package commentUsecase

import (
	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/app/comment"
	commModel "github.com/go-park-mail-ru/2020_2_JMickhs/internal/app/comment/models"
)

type CommentUseCase struct {
	commentRepo comment.Repository
}

func NewCommentUsecase(r comment.Repository) *CommentUseCase {
	return &CommentUseCase{
		commentRepo: r,
	}
}

func (u *CommentUseCase) GetComments(hotelID int, StartID int) ([]commModel.FullCommentInfo, error) {
	return u.commentRepo.GetComments(hotelID, StartID)
}
func (u *CommentUseCase) AddComment(comment commModel.Comment) (commModel.NewRate, error) {
	newRate := commModel.NewRate{}
	rate := commModel.Rating{comment.UserID, comment.HotelID, comment.UserID, comment.Rate}
	hotelRate, err := u.UpdateRating(rate)
	if err != nil {
		return newRate, err
	}
	newRate.Rate = hotelRate
	comm, err := u.commentRepo.AddComment(comment)
	if err != nil {
		return newRate, err
	}
	newRate.Comment = comm
	return newRate, nil

}
func (u *CommentUseCase) DeleteComment(ID int) error {
	return u.commentRepo.DeleteComment(ID)
}
func (u *CommentUseCase) UpdateComment(comment commModel.Comment) error {
	return u.commentRepo.UpdateComment(comment)
}

func (p *CommentUseCase) UpdateRating(rating commModel.Rating) (int, error) {
	err := p.commentRepo.InsertRating(rating)
	nextRate := -1
	if err != nil {
		return nextRate, err
	}

	currRate, err := p.commentRepo.GetCurrentRating(rating.HotelID)
	if err != nil {
		return nextRate, err
	}

	summRate := (currRate.RatesCount - 1) * currRate.CurrRating
	nextRate = (summRate + rating.Rate) / currRate.RatesCount

	err = p.commentRepo.UpdateHotelRating(rating.HotelID, nextRate)
	if err != nil {
		return nextRate, err
	}
	return nextRate, nil
}
