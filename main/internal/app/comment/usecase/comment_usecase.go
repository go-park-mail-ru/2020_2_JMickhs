package commentUsecase

import (
	"context"
	"math"
	"mime/multipart"
	"strconv"

	"github.com/go-park-mail-ru/2020_2_JMickhs/main/configs"
	"github.com/spf13/viper"

	"github.com/go-park-mail-ru/2020_2_JMickhs/main/internal/app/comment"
	commModel "github.com/go-park-mail-ru/2020_2_JMickhs/main/internal/app/comment/models"

	customerror "github.com/go-park-mail-ru/2020_2_JMickhs/package/error"
	userService "github.com/go-park-mail-ru/2020_2_JMickhs/package/proto/user"
	"github.com/go-park-mail-ru/2020_2_JMickhs/package/serverError"
)

type CommentUseCase struct {
	commentRepo comment.Repository
	userService userService.UserServiceClient
}

func NewCommentUsecase(r comment.Repository, userService userService.UserServiceClient) *CommentUseCase {
	return &CommentUseCase{
		commentRepo: r,
		userService: userService,
	}
}

func (u *CommentUseCase) CheckUserComment(comment commModel.Comment) (bool, error) {
	return u.commentRepo.CheckUserComment(comment)
}

func (u *CommentUseCase) GetPhotos(hotelID string) (commModel.Photos, error) {
	return u.commentRepo.GetPhotos(hotelID)
}

func (u *CommentUseCase) DeletePhotos(comment commModel.Comment) error {
	return u.commentRepo.DeletePhotos(comment)
}

func (u *CommentUseCase) UploadPhoto(comment *commModel.Comment, file multipart.File, contentType string) error {
	path, err := u.commentRepo.UploadPhoto(file, contentType)
	if err != nil {
		return err
	}
	comment.Photos = append(comment.Photos, path)
	return nil
}

func (u *CommentUseCase) GetComments(hotelID string, limit string, offsets string, user_id int) (int, commModel.Comments, error) {
	pag := commModel.Comments{}
	hotId, _ := strconv.Atoi(hotelID)
	lim, _ := strconv.Atoi(limit)
	offset, _ := strconv.Atoi(offsets)

	if lim < 1 || lim > 30 {
		lim = 10
	}
	count, err := u.commentRepo.GetCommentsCount(hotId)
	if err != nil {
		return 0, pag, err
	}
	if user_id != 0 {
		check, err := u.commentRepo.CheckRateExistForComments(hotId, user_id)
		if err != nil {
			return 0, pag, err
		}
		if check {
			count--
		}
	}
	if offset > count {

		return 0, pag, nil
	}

	comments, err := u.commentRepo.GetComments(hotelID, lim, offsets, user_id)
	if err != nil {
		return 0, pag, err
	}
	for pos, comm := range comments {
		user, err := u.userService.GetUserByID(context.Background(), &userService.UserID{UserID: int64(comm.UserID)})
		if err != nil {
			return count, pag, customerror.NewCustomError(err, serverError.ServerInternalError, 1)
		}
		comments[pos].Username = user.Username
		comments[pos].Avatar = user.Avatar
	}

	pag.Comments = comments

	pag.Info.ItemsCount = count
	return count, pag, nil
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
	for i := 0; i < len(newRate.Comment.Photos); i++ {
		newRate.Comment.Photos[i] = viper.GetString(configs.ConfigFields.S3Url) + newRate.Comment.Photos[i]
	}
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
