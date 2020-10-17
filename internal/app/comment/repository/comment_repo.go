package commentRepository

import (
	"net/http"
	"strconv"

	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/app/sqlrequests"

	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/app/comment/models"
	customerror "github.com/go-park-mail-ru/2020_2_JMickhs/internal/pkg/error"
	"github.com/jmoiron/sqlx"
)

type CommentRepository struct {
	conn *sqlx.DB
}

func NewCommentRepository(conn *sqlx.DB) CommentRepository {
	return CommentRepository{conn: conn}
}

func (r *CommentRepository) GetComments(hotelID int, StartID int) ([]models.FullCommentInfo, error) {
	rows, err := r.conn.Query(sqlrequests.GetCommentsPostgreRequest, strconv.Itoa(hotelID), strconv.Itoa(StartID))
	defer rows.Close()
	comments := []models.FullCommentInfo{}
	if err != nil {
		return comments, customerror.NewCustomError(err.Error(), http.StatusInternalServerError)
	}
	comment := models.FullCommentInfo{}
	for rows.Next() {
		err := rows.Scan(&comment.UserID, &comment.CommID, &comment.Message, &comment.Rating, &comment.Avatar, &comment.Username, &comment.HotelID)
		if err != nil {
			return comments, customerror.NewCustomError(err.Error(), http.StatusInternalServerError)
		}
		comments = append(comments, comment)
	}
	return comments, nil
}

func (r *CommentRepository) AddComment(comment models.Comment) (models.Comment, error) {
	var id int
	err := r.conn.QueryRow(sqlrequests.AddCommentsPostgreRequest,
		comment.UserID, comment.HotelID, comment.Message, comment.Rating).Scan(&id)

	if err != nil {
		return comment, customerror.NewCustomError(err.Error(), http.StatusInternalServerError)
	}
	comment.CommID = id
	return comment, nil
}

func (r *CommentRepository) DeleteComment(ID int) error {
	_, err := r.conn.Query(sqlrequests.DeleteCommentsPostgreRequest, ID)
	if err != nil {
		return customerror.NewCustomError(err.Error(), http.StatusInternalServerError)
	}
	return nil
}

func (r *CommentRepository) UpdateComment(comment models.Comment) error {
	_, err := r.conn.Query(sqlrequests.UpdateCommentsPostgreRequest,
		comment.CommID, comment.Message)
	if err != nil {
		return customerror.NewCustomError(err.Error(), http.StatusInternalServerError)
	}
	return nil
}
