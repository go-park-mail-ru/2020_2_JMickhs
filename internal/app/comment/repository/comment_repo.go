package commentRepository

import (
	"errors"
	"strconv"

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

func (r *CommentRepository) GetComments(hotelID int, StartID int) ([]commModel.FullCommentInfo, error) {
	rows, err := r.conn.Query(sqlrequests.GetCommentsPostgreRequest, strconv.Itoa(hotelID), strconv.Itoa(StartID))
	comments := []commModel.FullCommentInfo{}
	if err != nil {
		return comments, customerror.NewCustomError(err, serverError.ServerInternalError, nil)
	}
	comment := commModel.FullCommentInfo{}
	for rows.Next() {
		err := rows.Scan(&comment.UserID, &comment.CommID, &comment.Message, &comment.Rating, &comment.Avatar, &comment.Username, &comment.HotelID, &comment.Time)
		if err != nil {
			return comments, customerror.NewCustomError(err, serverError.ServerInternalError, nil)
		}
		comments = append(comments, comment)
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

func (p *CommentRepository) GetCurrentRating(hotelID int) (commModel.RateInfo, error) {
	rateInfo := commModel.RateInfo{}

	err := p.conn.QueryRow(sqlrequests.GetRatingCountOnHotelPostgreRequest, hotelID).Scan(&rateInfo.RatesCount)
	if err != nil {
		return rateInfo, customerror.NewCustomError(err, serverError.ServerInternalError, nil)
	}

	err = p.conn.QueryRow(sqlrequests.GetCurrRatingPostgreRequest, hotelID).Scan(&rateInfo.CurrRating)
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
