//go:generate mockgen -source repository.go -destination mocks/user_repository_mock.go -package user_mock
package user

import (
	"mime/multipart"

	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/app/user/models"
)

type Repository interface {
	GetByUserName(name string) (models.User, error)
	Add(user models.User) (models.User, error)
	GetUserByID(ID int) (models.User, error)
	UpdateUser(user models.User) error
	UpdateAvatar(user models.User) error
	UpdatePassword(user models.User) error
	GenerateHashFromPassword(password string) ([]byte, error)
	CompareHashAndPassword(hashedPassword string, password string) error
	DeleteAvatarInStore(user models.User, filename string) error
	UpdateAvatarInStore(file multipart.File, user *models.User, fileType string) error
}
