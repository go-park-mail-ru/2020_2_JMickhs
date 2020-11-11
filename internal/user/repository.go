package user

import "github.com/go-park-mail-ru/2020_2_JMickhs/internal/user/models"

type Repository interface {
	GetByUserName(name string) (models.User, error)
	Add(user models.User) (models.User, error)
	GetUserByID(ID int) (models.User, error)
	UpdateUser(user models.User) error
	UpdateAvatar(user models.User) error
	UpdatePassword(user models.User) error
}
