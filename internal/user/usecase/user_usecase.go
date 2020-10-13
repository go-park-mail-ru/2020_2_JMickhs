package userUsecase

import (
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strings"

	customerror "github.com/go-park-mail-ru/2020_2_JMickhs/internal/error"

	"github.com/go-park-mail-ru/2020_2_JMickhs/configs"

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
	return user, err
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
	if err != nil {
		return customerror.NewCustomError(err.Error(), http.StatusBadRequest)
	}
	user.Password = string(hashedPassword)
	err = u.userRepo.UpdatePassword(user)
	if err != nil {
		return customerror.NewCustomError(err.Error(), customerror.ParseCode(err))
	}
	return nil
}

func (u *userUseCase) UpdateAvatar(user models.User) error {
	return u.userRepo.UpdateAvatar(user)
}

func (u *userUseCase) UploadAvatar(file multipart.File, header string, user *models.User) error {
	filename := uuid.NewV4().String()
	fileType := strings.Split(header, "/")
	user.Avatar = configs.StaticPath + "/" + filename + "." + fileType[1]
	f, err := os.OpenFile("../"+user.Avatar, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return customerror.NewCustomError(err.Error(), http.StatusInternalServerError)
	}
	defer f.Close()
	io.Copy(f, file)
	return nil
}

func (u *userUseCase) CheckAvatar(file multipart.File) (string, error) {
	fileHeader := make([]byte, 512)
	ContentType := ""
	if _, err := file.Read(fileHeader); err != nil {
		return ContentType, customerror.NewCustomError(err.Error(), http.StatusInternalServerError)
	}

	if _, err := file.Seek(0, 0); err != nil {
		return ContentType, customerror.NewCustomError(err.Error(), http.StatusInternalServerError)
	}

	length, _ := file.Seek(0, 2)
	if length > 5*configs.MB {
		return ContentType, customerror.NewCustomError("file bigger then 5 MB", http.StatusBadRequest)
	}

	if _, err := file.Seek(0, 0); err != nil {
		return ContentType, customerror.NewCustomError(err.Error(), http.StatusInternalServerError)
	}

	ContentType = http.DetectContentType(fileHeader)

	if ContentType != "image/jpg" && ContentType != "image/png" && ContentType != "image/jpeg" {
		return ContentType, customerror.NewCustomError("Wrong file type", http.StatusUnsupportedMediaType)
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
