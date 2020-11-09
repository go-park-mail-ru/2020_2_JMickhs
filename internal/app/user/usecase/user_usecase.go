package userUsecase

import (
	"errors"
	"mime/multipart"
	"net/http"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	uuid "github.com/satori/go.uuid"

	"github.com/aws/aws-sdk-go/service/s3"

	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/pkg/clientError"
	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/pkg/serverError"

	"github.com/go-park-mail-ru/2020_2_JMickhs/configs"
	customerror "github.com/go-park-mail-ru/2020_2_JMickhs/internal/pkg/error"
	"github.com/go-playground/validator/v10"

	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/app/user"
	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/app/user/models"
	"golang.org/x/crypto/bcrypt"
)

type userUseCase struct {
	userRepo   user.Repository
	validation *validator.Validate
	s3         *s3.S3
}

func NewUserUsecase(r user.Repository, validator *validator.Validate, s3 *s3.S3) *userUseCase {
	return &userUseCase{
		userRepo:   r,
		validation: validator,
		s3:         s3,
	}
}

func (u *userUseCase) GetByUserName(name string) (models.User, error) {
	return u.userRepo.GetByUserName(name)
}

func (u *userUseCase) Add(user models.User) (models.User, error) {
	err := u.validation.Struct(user)
	if err != nil {
		return user, customerror.NewCustomError(err, clientError.BadRequest, 1)
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)
	u.SetDefaultAvatar(&user)
	user, err = u.userRepo.Add(user)
	return user, err
}

func (u *userUseCase) GetUserByID(ID int) (models.User, error) {
	return u.userRepo.GetUserByID(ID)
}

func (u *userUseCase) SetDefaultAvatar(user *models.User) error {
	user.Avatar = configs.S3Url + configs.BaseAvatarPath
	return nil
}

func (u *userUseCase) UpdateUser(user models.User) error {
	return u.userRepo.UpdateUser(user)
}

func (u *userUseCase) UpdatePassword(user models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return customerror.NewCustomError(err, clientError.BadRequest, 1)
	}
	user.Password = string(hashedPassword)
	err = u.userRepo.UpdatePassword(user)
	if err != nil {
		return err
	}
	return nil
}

func (u *userUseCase) UpdateAvatar(user models.User) error {
	return u.userRepo.UpdateAvatar(user)
}

func (u *userUseCase) UploadAvatar(file multipart.File, header string, user *models.User) (string, error) {
	relDele := strings.Split(user.Avatar, "/")

	if user.Avatar != configs.S3Url+configs.BaseAvatarPath {
		_, err := u.s3.DeleteObject(&s3.DeleteObjectInput{
			Bucket: aws.String(configs.BucketName),
			Key:    aws.String(configs.StaticPath + relDele[len(relDele)-1]),
		})
		if err != nil {
			return "", customerror.NewCustomError(err, http.StatusInternalServerError, 1)
		}
	}
	filename := uuid.NewV4().String()
	fileType := strings.Split(header, "/")
	relPath := configs.StaticPath + filename + "." + fileType[1]
	user.Avatar = configs.S3Url + relPath

	_, err := u.s3.PutObject(&s3.PutObjectInput{
		Body:   file,
		Bucket: aws.String(configs.BucketName),
		Key:    aws.String(relPath),
		ACL:    aws.String(s3.BucketCannedACLPublicRead),
	})
	if err != nil {
		return "", customerror.NewCustomError(err, http.StatusInternalServerError, 1)
	}
	return user.Avatar, nil

}

func (u *userUseCase) CheckAvatar(file multipart.File) (string, error) {
	fileHeader := make([]byte, 512)
	ContentType := ""
	if _, err := file.Read(fileHeader); err != nil {
		return ContentType, customerror.NewCustomError(err, serverError.ServerInternalError, 1)
	}

	if _, err := file.Seek(0, 0); err != nil {
		return ContentType, customerror.NewCustomError(err, serverError.ServerInternalError, 1)
	}

	count,err := file.Seek(0,2)
	if err != nil{
		return ContentType, customerror.NewCustomError(err, serverError.ServerInternalError, 1)
	}
    if count > 5 * configs.MB {
		return ContentType, customerror.NewCustomError(errors.New("file bigger than 5 mb"), clientError.BadRequest, 1)
	}
	if _,err := file.Seek(0,0); err != nil{
		return ContentType, customerror.NewCustomError(err, serverError.ServerInternalError, 1)
	}
	ContentType = http.DetectContentType(fileHeader)

	if ContentType != "image/jpg" && ContentType != "image/png" && ContentType != "image/jpeg" {
		return ContentType, customerror.NewCustomError(errors.New("Wrong file type"), clientError.UnsupportedMediaType, 1)
	}

	return ContentType, nil
}

func (u *userUseCase) ComparePassword(passIn string, passDest string) error {
	return bcrypt.CompareHashAndPassword([]byte(passDest), []byte(passIn))
}

func (u *userUseCase) CheckEmpty(usr models.User) bool {
	empty := models.User{}
	return usr == empty
}
