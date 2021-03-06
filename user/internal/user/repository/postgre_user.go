package userRepository

import (
	"mime/multipart"
	"strconv"

	"github.com/go-park-mail-ru/2020_2_JMickhs/user/configs"
	"github.com/go-park-mail-ru/2020_2_JMickhs/user/internal/user/models"

	"github.com/spf13/viper"

	customerror "github.com/go-park-mail-ru/2020_2_JMickhs/package/error"

	"github.com/go-park-mail-ru/2020_2_JMickhs/package/clientError"
	"github.com/go-park-mail-ru/2020_2_JMickhs/package/serverError"

	uuid "github.com/satori/go.uuid"

	"golang.org/x/crypto/bcrypt"

	"github.com/aws/aws-sdk-go/aws"

	"github.com/aws/aws-sdk-go/service/s3"

	"github.com/jmoiron/sqlx"
)

type PostgresUserRepository struct {
	conn *sqlx.DB
	s3   *s3.S3
}

func NewPostgresUserRepository(conn *sqlx.DB, s3 *s3.S3) PostgresUserRepository {
	userStorage := PostgresUserRepository{conn, s3}
	return userStorage
}

func (p *PostgresUserRepository) Add(user models.User) (models.User, error) {
	var id int
	err := p.conn.QueryRow(AddUserPostgreRequest, user.Username, user.Email, user.Password, user.Avatar).Scan(&id)
	if err != nil {
		return user, customerror.NewCustomError(err, clientError.Conflict, 1)
	}
	user.ID = id
	user.Avatar = viper.GetString(configs.ConfigFields.S3Url) + user.Avatar
	return user, nil
}
func (p *PostgresUserRepository) DeleteAvatarInStore(user models.User, filename string) error {
	if user.Avatar != viper.GetString(configs.ConfigFields.S3Url)+viper.GetString(configs.ConfigFields.BaseAvatarPath) {
		var _, err = p.s3.DeleteObject(&s3.DeleteObjectInput{
			Bucket: aws.String(viper.GetString(configs.ConfigFields.BucketName)),
			Key:    aws.String(viper.GetString(configs.ConfigFields.StaticPathForAvatars) + filename),
		})
		return err
	}
	return nil
}

func (p *PostgresUserRepository) UpdateAvatarInStore(file multipart.File, user *models.User, fileType string) error {

	newFilename := uuid.NewV4().String()
	relativePath := viper.GetString(configs.ConfigFields.StaticPathForAvatars) + newFilename + "." + fileType

	_, err := p.s3.PutObject(&s3.PutObjectInput{
		Body:   file,
		Bucket: aws.String(viper.GetString(configs.ConfigFields.BucketName)),
		Key:    aws.String(relativePath),
		ACL:    aws.String(s3.BucketCannedACLPublicRead),
	})

	if err == nil {
		user.Avatar = relativePath
	}
	return err
}

func (u *PostgresUserRepository) GenerateHashFromPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}
func (u *PostgresUserRepository) CompareHashAndPassword(hashedPassword string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (p *PostgresUserRepository) GetByUserName(name string) (models.User, error) {
	user := models.User{}
	err := p.conn.QueryRow(GetUserByNamePostgreRequest, name, viper.GetString(configs.ConfigFields.S3Url)).
		Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Avatar, &user.ModRule)
	if err != nil {
		return user, customerror.NewCustomError(err, clientError.Unauthorizied, 1)
	}
	return user, nil
}

func (p *PostgresUserRepository) GetUserByID(ID int) (models.User, error) {
	row := p.conn.QueryRow(GetUserByIDPostgreRequest, strconv.Itoa(ID), viper.GetString(configs.ConfigFields.S3Url))
	user := models.User{}

	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Avatar, &user.ModRule)
	if err != nil {
		return user, customerror.NewCustomError(err, clientError.Gone, 1)
	}
	return user, nil
}

func (p *PostgresUserRepository) UpdateUser(user models.User) error {
	_, err := p.conn.Query(UpdateUserPostgreRequest,
		user.ID, user.Username, user.Email)
	if err != nil {
		return customerror.NewCustomError(err, clientError.Conflict, 1)
	}
	return nil

}

func (p *PostgresUserRepository) UpdateAvatar(user models.User) error {
	_, err := p.conn.Query(UpdateUserAvatarPostgreRequest,
		user.ID, user.Avatar)
	if err != nil {
		return customerror.NewCustomError(err, serverError.ServerInternalError, 1)
	}
	return nil

}

func (p *PostgresUserRepository) UpdatePassword(user models.User) error {
	_, err := p.conn.Query(UpdateUserPasswordPostgreRequest,
		user.ID, user.Password)
	if err != nil {
		return customerror.NewCustomError(err, serverError.ServerInternalError, 1)
	}
	return nil
}
