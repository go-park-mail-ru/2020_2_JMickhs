package commentRepository

import (
	"net/http"
	"strconv"

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
		return comments, customerror.NewCustomError(err.Error(), http.StatusInternalServerError)
	}
	comment := commModel.FullCommentInfo{}
	for rows.Next() {
		err := rows.Scan(&comment.UserID, &comment.CommID, &comment.Message, &comment.Rating, &comment.Avatar, &comment.Username, &comment.HotelID, &comment.Time)
		if err != nil {
			return comments, customerror.NewCustomError(err.Error(), http.StatusInternalServerError)
		}
		comments = append(comments, comment)
	}
	return comments, nil
}

func (r *CommentRepository) AddComment(comment commModel.Comment) (commModel.Comment, error) {
	err := r.conn.QueryRow(sqlrequests.AddCommentsPostgreRequest,
		comment.UserID, comment.HotelID, comment.Message).Scan(&comment.CommID, &comment.Time)
	if err != nil {
		return comment, customerror.NewCustomError(err.Error(), http.StatusInternalServerError)
	}
	return comment, nil
}

func (r *CommentRepository) DeleteComment(ID int) error {
	_, err := r.conn.Query(sqlrequests.DeleteCommentsPostgreRequest, ID)
	if err != nil {
		return customerror.NewCustomError(err.Error(), http.StatusInternalServerError)
	}
	return nil
}

func (r *CommentRepository) UpdateComment(comment commModel.Comment) error {
	_, err := r.conn.Query(sqlrequests.UpdateCommentsPostgreRequest,
		comment.CommID, comment.Message)
	if err != nil {
		return customerror.NewCustomError(err.Error(), http.StatusInternalServerError)
	}
	return nil
}

func (p *CommentRepository) InsertRating(rating commModel.Rating) error {
	err := p.conn.QueryRow(sqlrequests.InsertRatingPostgreRequest, rating.HotelID, rating.UserID, rating.Rate).Err()
	if err != nil {
		return customerror.NewCustomError(err.Error(), http.StatusLocked)
	}
	return nil
}

func (p *CommentRepository) UpdateHotelRating(hotelID int, NewRate int) error {
	err := p.conn.QueryRow(sqlrequests.UpdateHotelRatingPostgreRequest, NewRate, hotelID).Err()
	if err != nil {
		return customerror.NewCustomError(err.Error(), http.StatusBadRequest)
	}
	return nil
}

func (p *CommentRepository) GetCurrentRating(hotelID int) (commModel.RateInfo, error) {
	rateInfo := commModel.RateInfo{}

	err := p.conn.QueryRow(sqlrequests.GetRatingCountOnHotelPostgreRequest, hotelID).Scan(&rateInfo.RatesCount)
	if err != nil {
		return rateInfo, customerror.NewCustomError(err.Error(), http.StatusInternalServerError)
	}

	err = p.conn.QueryRow(sqlrequests.GetCurrRatingPostgreRequest, hotelID).Scan(&rateInfo.CurrRating)
	if err != nil {
		return rateInfo, customerror.NewCustomError(err.Error(), http.StatusInternalServerError)
	}
	return rateInfo, nil
}
