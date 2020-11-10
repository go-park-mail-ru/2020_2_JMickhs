package commentUsecase

import (
	"fmt"
	"math"
	"net/url"
	"strconv"

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

func (u *CommentUseCase) GetComments(hotelID string, limit string, offsets string, user_id int) (commModel.Comments, error) {
	pag := commModel.Comments{}
	hotId, _ := strconv.Atoi(hotelID)
	lim, _ := strconv.Atoi(limit)
	offset, _ := strconv.Atoi(offsets)

	if lim < 1 || lim > 30 {
		lim = 10
	}
	count, err := u.commentRepo.GetCommentsCount(hotId)
	if err != nil {
		return pag, err
	}
	if user_id != 0 {
		check, err := u.commentRepo.CheckRateExistForComments(hotId, user_id)
		if err != nil {
			return pag, err
		}
		if check {
			count--
		}
	}

	if offset > count {
		return pag, nil
	}

	data, err := u.commentRepo.GetComments(hotelID, lim, offsets, user_id)
	if err != nil {
		return pag, err
	}
	pag.Comments = data

	url := url.URL{
		Host:   "hostelscan.ru:8080",
		Scheme: "https",
		Path:   "/api/v1/comments/",
	}
	query := url.Query()
	query.Set("id", hotelID)
	query.Set("limit", limit)
	if offset+lim >= count {
		query.Set("offset", strconv.Itoa(offset-lim))
		url.RawQuery = query.Encode()
		pag.Info.PrevLink = url.Path + "?" + url.RawQuery

		query.Set("offset", "0")
		url.RawQuery = query.Encode()
		pag.Info.NextLink = url.Path + "?" + url.RawQuery
	}
	if offset-lim < 0 {
		query.Set("offset", strconv.Itoa(offset+lim))
		url.RawQuery = query.Encode()
		pag.Info.NextLink = url.Path + "?" + url.RawQuery

		query.Set("offset", strconv.Itoa(count-lim))
		url.RawQuery = query.Encode()
		pag.Info.PrevLink = url.Path + "?" + url.RawQuery
	}
	if offset+lim < count && offset-lim >= 0 {
		query.Set("offset", strconv.Itoa(offset+lim))
		url.RawQuery = query.Encode()
		pag.Info.NextLink = url.Path + "?" + url.RawQuery

		query.Set("offset", strconv.Itoa(offset-lim))
		url.RawQuery = query.Encode()
		pag.Info.PrevLink = url.Path + "?" + url.RawQuery
	}
	pag.Info.ItemsCount = count
	return pag, nil
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
	fmt.Println(hotelRate)
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
	nextRate = (summRate + comment.Rate) / float64(currRate.RatesCount)

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
