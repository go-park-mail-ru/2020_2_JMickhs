package commentRepository

import (
	"errors"
	"mime/multipart"
	"strconv"

	"github.com/lib/pq"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/go-park-mail-ru/2020_2_JMickhs/main/configs"
	uuid "github.com/satori/go.uuid"
	"github.com/spf13/viper"

	"github.com/aws/aws-sdk-go/service/s3"

	commModel "github.com/go-park-mail-ru/2020_2_JMickhs/main/internal/app/comment/models"

	customerror "github.com/go-park-mail-ru/2020_2_JMickhs/package/error"

	"github.com/go-park-mail-ru/2020_2_JMickhs/package/clientError"
	"github.com/go-park-mail-ru/2020_2_JMickhs/package/serverError"

	"github.com/jmoiron/sqlx"
)

type CommentRepository struct {
	conn *sqlx.DB
	s3   *s3.S3
}

func NewCommentRepository(conn *sqlx.DB, s3 *s3.S3) CommentRepository {
	return CommentRepository{conn: conn, s3: s3}
}

func (r *CommentRepository) GetPhotos(hotelID string) (commModel.Photos, error) {
	photos := commModel.Photos{}

	err := r.conn.Select(&photos.Photos, GetCommentsPhotosPostgreRequest, hotelID, viper.GetString(configs.ConfigFields.S3Url))
	if err != nil {
		return photos, customerror.NewCustomError(err, serverError.ServerInternalError, 1)
	}
	return photos, nil
}

func (r *CommentRepository) DeletePhotos(comment commModel.Comment) error {
	for _, photo := range comment.Photos {
		var _, err = r.s3.DeleteObject(&s3.DeleteObjectInput{
			Bucket: aws.String(viper.GetString(configs.ConfigFields.BucketName)),
			Key:    aws.String(viper.GetString(configs.ConfigFields.StaticPathForComments) + photo),
		})
		return err
	}
	return nil
}

func (r *CommentRepository) UploadPhoto(file multipart.File, contentType string) (string, error) {
	newFilename := uuid.NewV4().String()
	relativePath := viper.GetString(configs.ConfigFields.StaticPathForComments) + newFilename + "." + contentType

	_, err := r.s3.PutObject(&s3.PutObjectInput{
		Body:   file,
		Bucket: aws.String(viper.GetString(configs.ConfigFields.BucketName)),
		Key:    aws.String(relativePath),
		ACL:    aws.String(s3.BucketCannedACLPublicRead),
	})

	return relativePath, err
}

func (r *CommentRepository) GetComments(hotelID string, limit int, offset string, user_id int) ([]commModel.FullCommentInfo, error) {
	var comments []commModel.FullCommentInfo
	err := r.conn.Select(&comments, GetCommentsPostgreRequest, offset, limit, hotelID, user_id)
	if err != nil {
		return comments, customerror.NewCustomError(err, clientError.BadRequest, 1)
	}

	for i := 0; i < len(comments); i++ {
		err = r.conn.Select(&comments[i].Photos, CheckPhotosExistPostgreRequest, hotelID, comments[i].UserID, viper.GetString(configs.ConfigFields.S3Url))
		if err != nil {
			return comments, customerror.NewCustomError(err, serverError.ServerInternalError, 1)
		}
	}
	return comments, nil
}

func (r *CommentRepository) AddComment(comment commModel.Comment) (commModel.Comment, error) {
	err := r.conn.QueryRow(AddCommentsPostgreRequest,
		comment.UserID, comment.HotelID, comment.Message, comment.Rate, pq.Array(comment.Photos)).Scan(&comment.CommID, &comment.Time)
	if err != nil {
		return comment, customerror.NewCustomError(err, clientError.Locked, 1)
	}
	return comment, nil
}

func (r *CommentRepository) DeleteComment(ID int) error {
	_, err := r.conn.Query(DeleteCommentsPostgreRequest, strconv.Itoa(ID))
	if err != nil {
		return customerror.NewCustomError(err, serverError.ServerInternalError, 1)
	}
	return nil
}

func (r *CommentRepository) UpdateComment(comment *commModel.Comment) error {
	err := r.conn.QueryRow(UpdateCommentsPostgreRequest,
		comment.CommID, comment.Message, comment.Rate, pq.Array(comment.Photos)).
		Scan(&comment.Time)
	if err != nil {
		return customerror.NewCustomError(err, serverError.ServerInternalError, 1)
	}
	comment.Photos = []string{}
	err = r.conn.Select(&comment.Photos, GetOneCommentPhotosPostgreRequest, comment.CommID, viper.GetString(configs.ConfigFields.S3Url))
	if err != nil {
		return customerror.NewCustomError(err, serverError.ServerInternalError, 1)
	}
	return nil
}

func (p *CommentRepository) UpdateHotelRating(hotelID int, NewRate float64) error {
	rate := strconv.FormatFloat(NewRate, 'f', 1, 64)
	err := p.conn.QueryRow(UpdateHotelRatingPostgreRequest, rate, strconv.Itoa(hotelID)).Err()
	if err != nil {
		return customerror.NewCustomError(err, clientError.BadRequest, 1)
	}
	return nil
}

func (p *CommentRepository) GetCommentsCount(hotelID int) (int, error) {
	count := -1
	err := p.conn.QueryRow(GetCommentsCountPostgreRequest, hotelID).Scan(&count)

	if err != nil {
		return count, customerror.NewCustomError(err, serverError.ServerInternalError, 1)
	}
	return count, nil
}

func (p *CommentRepository) CheckUserComment(comment commModel.Comment) (bool, error) {
	user_id := -1
	err := p.conn.QueryRow(CheckCommentForUpdateRequest, comment.CommID).Scan(&user_id)
	if err != nil {
		return false, customerror.NewCustomError(err, serverError.ServerInternalError, 1)
	}
	return user_id == comment.UserID, nil

}

func (p *CommentRepository) GetCurrentRating(hotelID int) (commModel.RateInfo, error) {
	rateInfo := commModel.RateInfo{}

	err := p.conn.QueryRow(GetCurrRatingPostgreRequest, hotelID).Scan(&rateInfo.CurrRating, &rateInfo.RatesCount)
	if err != nil {
		return rateInfo, customerror.NewCustomError(err, serverError.ServerInternalError, 1)
	}

	return rateInfo, nil
}

func (p *CommentRepository) CheckUser(comment *commModel.Comment) (int, error) {
	var destRate int
	var usr_id int
	err := p.conn.QueryRow(GetPrevRatingOnCommentPostgreRequest, strconv.Itoa(comment.CommID)).Scan(&destRate, &usr_id, &comment.HotelID)
	if err != nil {
		return destRate, customerror.NewCustomError(err, clientError.NotFound, 1)
	}
	if comment.UserID != usr_id {
		return destRate, customerror.NewCustomError(errors.New("user want update other comment"), clientError.Locked, 1)
	}
	return destRate, nil
}

func (p *CommentRepository) CheckRateExistForComments(hotelID int, userID int) (bool, error) {
	res, err := p.conn.Exec(CheckRateExistForCommentsRequest, hotelID, userID)
	if err != nil {
		return false, customerror.NewCustomError(err, serverError.ServerInternalError, 1)
	}
	count, _ := res.RowsAffected()
	if count == 0 {
		return false, nil
	}
	return true, nil
}
