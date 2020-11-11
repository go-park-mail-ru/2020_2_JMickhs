package user

import (
	"mime/multipart"

	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/user/models"
)

type Usecase interface {
	GetByUserName(name string) (models.User, error)
	Add(user models.User) (models.User, error)
	GetUserByID(ID int) (models.User, error)
	SetDefaultAvatar(user *models.User) error
	UpdateUser(user models.User) error
	UpdateAvatar(user models.User) error
	UpdatePassword(user models.User) error
	UploadAvatar(file multipart.File, fileType string, user *models.User) error
	ComparePassword(passIn string, passDest string) error
	CheckEmpty(usr models.User) bool
}
