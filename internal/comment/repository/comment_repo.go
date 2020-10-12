package commentRepository

import (
	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/comment/models"
	"github.com/jmoiron/sqlx"
	"strconv"
)

type CommentRepository struct {
	conn *sqlx.DB
}

func NewCommentRepository(conn *sqlx.DB) CommentRepository{
	return CommentRepository{conn: conn}
}

func (r *CommentRepository) GetComments(hotelID int, StartID int) ([]models.FullCommentInfo, error){
	rows, err := r.conn.Query("SELECT us.user_id,comm_id,message,rating,avatar,username,hotel_id FROM comments as comm " +
		"INNER JOIN users as us on us.user_id = comm.user_id where hotel_id = $1 LIMIT 4 OFFSET $2",  strconv.Itoa(hotelID) , strconv.Itoa(StartID))
	defer rows.Close()
	comments := []models.FullCommentInfo{}
	if err != nil {
		return comments, err
	}
	comment := models.FullCommentInfo{}
	for rows.Next() {
		err := rows.Scan(&comment.UserID,&comment.CommID,&comment.Message,&comment.Rating,&comment.Avatar,&comment.Username,&comment.HotelID)
		if err != nil {
			return comments, err
		}
		comments = append(comments, comment)
	}
	return comments, nil
}

func (r *CommentRepository) AddComment(comment models.Comment) (models.Comment,error) {
	var id int
	err := r.conn.QueryRow("INSERT INTO comments VALUES (default, $1, $2,$3,$4) RETURNING comm_id",
		comment.UserID,comment.HotelID,comment.Message,comment.Rating).Scan(&id)
	comment.CommID = id
	return comment, err
}

func (r *CommentRepository) DeleteComment(ID int) error{
	_, err := r.conn.Query("DELETE FROM comments WHERE comm_id=$1", ID)
	return err
}

func (r *CommentRepository) UpdateComment(comment models.Comment) (error){
	_, err := r.conn.Query("UPDATE comments SET message=$2 WHERE comm_id=$1",
		comment.CommID,comment.Message)
	return err
}