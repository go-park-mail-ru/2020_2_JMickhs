//go:generate mockgen -source usecase.go -destination mocks/user_usecase_mock.go -package user_mock
package user

import (
	"mime/multipart"

	"github.com/go-park-mail-ru/2020_2_JMickhs/JMickhs_main/internal/app/user/models"
)

type Usecase interface {
	GetByUserName(name string) (models.User, error)
	Add(user models.User) (models.User, error)
	GetUserByID(ID int) (models.User, error)
	SetDefaultAvatar(user *models.User) error
	UpdateUser(user models.User) error
	UpdateAvatar(user models.User) error
	UpdatePassword(user models.User) error
	UploadAvatar(file multipart.File, fileType string, user *models.User) (string, error)
	ComparePassword(passIn string, passDest string) error
	CheckEmpty(usr models.User) bool
	CheckAvatar(file multipart.File) (string, error)
}
