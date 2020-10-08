package userUsecase

import (
	"errors"
	"github.com/go-park-mail-ru/2020_2_JMickhs/configs"
	"io"
	"mime/multipart"
	"os"

	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/user"
	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/user/models"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type userUseCase struct {
	userRepo user.Repository
}

func NewUserUsecase(r user.Repository) *userUseCase {
	return &userUseCase{
		userRepo: r,
	}
}

func (u *userUseCase) GetByUserName(name string) (models.User, error) {
	user, err := u.userRepo.GetByUserName(name)
	return user, err
}

func (u *userUseCase) Add(user models.User) (models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
	user.Password = string(hashedPassword)
	id, err := u.userRepo.Add(user)
	return id, err
}

func (u *userUseCase) GetUserByID(ID int) (models.User, error) {
	user, err := u.userRepo.GetUserByID(ID)
	if u.CheckEmpty(user){
		return user, errors.New("User doesn't exist")
	}
	return user,err
}

func (u *userUseCase) SetDefaultAvatar(user *models.User) error {
	user.Avatar = configs.BaseAvatarPath
	return nil
}

func (u *userUseCase) UpdateUser(user models.User) error {
	return u.userRepo.UpdateUser(user)
}

func (u *userUseCase) UpdatePassword(user models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
	user.Password = string(hashedPassword)
	err = u.userRepo.UpdatePassword(user)
	return err
}

func (u *userUseCase) UpdateAvatar(user models.User) error {

	return u.userRepo.UpdateAvatar(user)
}

func (u *userUseCase) UploadAvatar(file multipart.File, fileType string ,user *models.User) error {
	filename := uuid.NewV4().String()
	user.Avatar = configs.StaticPath + "/" +  filename + "." + fileType
	f, err := os.OpenFile("../" + user.Avatar, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer f.Close()
	io.Copy(f, file)
	return nil
}

func (u *userUseCase) ComparePassword(passIn string, passDest string) error {
	return bcrypt.CompareHashAndPassword([]byte(passDest), []byte(passIn))
}

func (u *userUseCase) CheckEmpty(usr models.User) bool {
	empty := models.User{}
	return usr == empty
}
