package commentRepository

import (
	"errors"
	"strconv"

	"github.com/go-park-mail-ru/2020_2_JMickhs/configs"

	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/pkg/clientError"
	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/pkg/serverError"

	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/app/sqlrequests"

	commModel "github.com/go-park-mail-ru/2020_2_JMickhs/internal/app/comment/models"
	customerror "github.com/go-park-mail-ru/2020_2_JMickhs/internal/pkg/error"
	"github.com/jmoiron/sqlx"
)

type CommentRepository struct {
	conn *sqlx.DB
}

func NewCommentRepository(conn *sqlx.DB) CommentRepository {
	return CommentRepository{conn: conn}
}

func (r *CommentRepository) GetComments(hotelID int, offset int) ([]commModel.FullCommentInfo, error) {
	comments := []commModel.FullCommentInfo{}
	err := r.conn.Select(&comments, sqlrequests.GetCommentsPostgreRequest, strconv.Itoa(offset), configs.BaseItemsPerPage, strconv.Itoa(hotelID))
	if err != nil {
		return comments, customerror.NewCustomError(err, clientError.BadRequest, nil)
	}
	return comments, nil
}

func (r *CommentRepository) AddComment(comment commModel.Comment) (commModel.Comment, error) {
	err := r.conn.QueryRow(sqlrequests.AddCommentsPostgreRequest,
		comment.UserID, comment.HotelID, comment.Message, comment.Rate).Scan(&comment.CommID, &comment.Time)
	if err != nil {
		return comment, customerror.NewCustomError(err, clientError.Locked, nil)
	}
	return comment, nil
}

func (r *CommentRepository) DeleteComment(ID int) error {
	_, err := r.conn.Query(sqlrequests.DeleteCommentsPostgreRequest, ID)
	if err != nil {
		return customerror.NewCustomError(err, serverError.ServerInternalError, nil)
	}
	return nil
}

func (r *CommentRepository) UpdateComment(comment *commModel.Comment) error {
	err := r.conn.QueryRow(sqlrequests.UpdateCommentsPostgreRequest,
		comment.CommID, comment.Message, comment.Rate).Scan(&comment.Time)
	if err != nil {
		return customerror.NewCustomError(err, serverError.ServerInternalError, nil)
	}
	return nil
}

func (p *CommentRepository) UpdateHotelRating(hotelID int, NewRate float64) error {
	err := p.conn.QueryRow(sqlrequests.UpdateHotelRatingPostgreRequest, NewRate, hotelID).Err()
	if err != nil {
		return customerror.NewCustomError(err, clientError.BadRequest, nil)
	}
	return nil
}

func (p *CommentRepository) GetCommentsCount(hotelID int) (int, error) {
	count := -1
	err := p.conn.QueryRow(sqlrequests.GetCommentsCountPostgreRequest, hotelID).Scan(&count)
	if err != nil {
		return count, customerror.NewCustomError(err, serverError.ServerInternalError, nil)
	}
	return count, nil
}

func (p *CommentRepository) GetCurrentRating(hotelID int) (commModel.RateInfo, error) {
	rateInfo := commModel.RateInfo{}

	err := p.conn.QueryRow(sqlrequests.GetCurrRatingPostgreRequest, hotelID).Scan(&rateInfo.CurrRating, &rateInfo.RatesCount)
	if err != nil {
		return rateInfo, customerror.NewCustomError(err, serverError.ServerInternalError, nil)
	}

	return rateInfo, nil
}

func (p *CommentRepository) CheckUser(comment *commModel.Comment) (int, error) {
	var destRate int
	var usr_id int
	err := p.conn.QueryRow(sqlrequests.GetPrevRatingOnCommentPostgreRequest, comment.CommID).Scan(&destRate, &usr_id, &comment.HotelID)
	if err != nil {
		return destRate, customerror.NewCustomError(err, serverError.ServerInternalError, nil)
	}
	if comment.UserID != usr_id {
		return destRate, customerror.NewCustomError(errors.New("user want update other comment"), clientError.Locked, nil)
	}
	return destRate, nil
}
