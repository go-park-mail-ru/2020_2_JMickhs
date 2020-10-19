package commentUsecase

import (
	"fmt"
	"math"

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
	comm, err := u.commentRepo.AddComment(comment)
	if err != nil {
		return newRate, err
	}
	newRate.Comment = comm
	hotelRate, err := u.AddRating(comment)
	if err != nil {
		return newRate, err
	}
	newRate.Rate = hotelRate
	return newRate, nil

}
func (u *CommentUseCase) DeleteComment(ID int) error {
	return u.commentRepo.DeleteComment(ID)
}
func (u *CommentUseCase) UpdateComment(comment commModel.Comment) (commModel.NewRate, error) {
	prevRate := commModel.PrevRate{}
	newRate := commModel.NewRate{}
	rate, err := u.commentRepo.CheckUser(&comment)
	if err != nil {
		return newRate, err
	}
	prevRate.Comment = comment
	prevRate.Rate = rate

	err = u.commentRepo.UpdateComment(&prevRate.Comment)
	if err != nil {
		return newRate, err
	}
	newRate.Comment = prevRate.Comment
	hotelRate, err := u.UpdateRating(prevRate)
	if err != nil {
		return newRate, err
	}
	newRate.Rate = hotelRate
	return newRate, nil
}

func (p *CommentUseCase) AddRating(comment commModel.Comment) (float64, error) {
	nextRate := float64(-1.0)

	currRate, err := p.commentRepo.GetCurrentRating(comment.HotelID)
	if err != nil {
		return nextRate, err
	}

	summRate := float64(currRate.RatesCount-1) * currRate.CurrRating
	fmt.Println(summRate)
	nextRate = (summRate + comment.Rate) / float64(currRate.RatesCount)
	fmt.Println(nextRate)

	err = p.commentRepo.UpdateHotelRating(comment.HotelID, nextRate)

	if err != nil {
		return nextRate, err
	}

	nextRate = math.Round(nextRate*10) / 10

	return nextRate, nil
}

func (p *CommentUseCase) UpdateRating(prevRate commModel.PrevRate) (float64, error) {
	nextRate := float64(-1.0)

	currRate, err := p.commentRepo.GetCurrentRating(prevRate.Comment.HotelID)
	if err != nil {
		return nextRate, err
	}

	summRate := float64(currRate.RatesCount) * currRate.CurrRating
	nextRate = (summRate - float64(prevRate.Rate) + prevRate.Comment.Rate) / float64(currRate.RatesCount)

	err = p.commentRepo.UpdateHotelRating(prevRate.Comment.HotelID, nextRate)
	if err != nil {
		return nextRate, err
	}
	nextRate = math.Round(nextRate*10) / 10
	return nextRate, nil
}
