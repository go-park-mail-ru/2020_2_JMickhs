package userDelivery

import (
	"encoding/json"
	"errors"
	"github.com/go-park-mail-ru/2020_2_JMickhs/configs"
	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/responses"
	"net/http"
	"strconv"
	"strings"
	"time"

	permissions "github.com/go-park-mail-ru/2020_2_JMickhs/internal/permission"

	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/sessions"

	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/user"
	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/user/models"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type UserHandler struct {
	UserUseCase     user.Usecase
	SessionsUseCase sessions.Usecase
	log             *logrus.Logger
}

func NewUserHandler(r *mux.Router, su sessions.Usecase, us user.Usecase, lg *logrus.Logger) {
	handler := UserHandler{
		UserUseCase:     us,
		SessionsUseCase: su,
		log:             lg,
	}

	r.HandleFunc("/api/v1/user/signup", permissions.SetCSRF(handler.Registration)).Methods("POST")
	r.HandleFunc("/api/v1/user/signin", permissions.SetCSRF(handler.Auth)).Methods("POST")
	r.HandleFunc("/api/v1/user", permissions.SetCSRF(handler.UserHandler)).Methods("GET")
	r.HandleFunc("/api/v1/user/signout", permissions.SetCSRF(handler.SignOut)).Methods("POST")
	r.HandleFunc("/api/v1/user/credentials", permissions.CheckCSRF(handler.UpdateUser)).Methods("PUT")
	r.HandleFunc("/api/v1/user/avatar", permissions.CheckCSRF(handler.UpdateAvatar)).Methods("PUT")
	r.HandleFunc("/api/v1/user/password", permissions.CheckCSRF(handler.updatePassword)).Methods("PUT")
	r.HandleFunc("/api/v1/user/{id:[0-9]+}", permissions.SetCSRF(handler.getAccInfo)).Methods("GET")
}

// swagger:route GET /api/v1/user/{id}  User userById
// Get info abous user by his id
// responses:
//  200: safeUser
func (u *UserHandler) getAccInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		u.log.Error(err.Error())
		responses.SendErrorResponse(w,http.StatusBadRequest,err)
		return
	}

	user, err := u.UserUseCase.GetUserByID(id)

	if err != nil {
		u.log.Error(err.Error())
		responses.SendErrorResponse(w,http.StatusInternalServerError,err)
		return
	}

	safeUser := models.SafeUser{ID: user.ID, Username: user.Username, Avatar: user.Avatar, Email: user.Email}

	responses.SendOkResponse(w,safeUser)
}

// swagger:route PUT /api/v1/user/avatar User avatar
// Update Avatar
func (u *UserHandler) UpdateAvatar(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(32 << 20)
	file, info, err := r.FormFile("avatar")

	if err != nil {
		u.log.Error(err.Error())
		responses.SendErrorResponse(w,http.StatusBadRequest,err)
		return
	}

	fileType := strings.Split(info.Header.Get("Content-Type"),"/")

	defer file.Close()
	usr, ok := r.Context().Value("User").(models.User)
	if !ok {
		responses.SendErrorResponse(w,http.StatusUnauthorized,errors.New("User Unauthorized"))
		return
	}
	extension := fileType[1]

	if (extension != "jpg" && extension != "png" && extension != "jpeg"){
		err = errors.New("bad type of file to add in static")
		u.log.Error(err)
		responses.SendErrorResponse(w,http.StatusBadRequest, err)
		return
	}
	u.UserUseCase.UploadAvatar(file, fileType[1] , &usr)

	err = u.UserUseCase.UpdateAvatar(usr)
	if err != nil {
		u.log.Error(err.Error())
		responses.SendErrorResponse(w,http.StatusInternalServerError, err)
		return
	}
}

// swagger:route PUT /api/v1/user/password User password
// update password
// responses:
// 409: conflict
func (u *UserHandler) updatePassword(w http.ResponseWriter, r *http.Request) {

	var twoPass models.UpdatePassword
	err := json.NewDecoder(r.Body).Decode(&twoPass)
	if err != nil {
		u.log.Error(err.Error())
		responses.SendErrorResponse(w,http.StatusBadRequest,err)
		return
	}

	usr, ok := r.Context().Value("User").(models.User)
	if !ok {
		responses.SendErrorResponse(w,http.StatusUnauthorized,errors.New("User Unauthorized"))
		return
	}

	err = u.UserUseCase.ComparePassword(twoPass.OldPassword,usr.Password)
	if err != nil {
		u.log.Info(err.Error())
		responses.SendErrorResponse(w,http.StatusConflict, errors.New("Wrong Old Password"))
		return
	}
	usr.Password = twoPass.NewPassword
	err = u.UserUseCase.UpdatePassword(usr)
	if err != nil {
		u.log.Error(err.Error())
		responses.SendErrorResponse(w,http.StatusInternalServerError, err)
		return
	}
}

// swagger:route PUT /api/v1/user/information User credentials
// Get data from form  which need to change and change user data
func (u *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		u.log.Error(err.Error())
		responses.SendErrorResponse(w,http.StatusBadRequest,err)
		return
	}

	usr,ok := r.Context().Value("User").(models.User)
	if !ok {
		responses.SendErrorResponse(w,http.StatusUnauthorized,errors.New("User Unauthorized"))
		return
	}
	user.ID = usr.ID

	err = u.UserUseCase.UpdateUser(user)
	if err != nil {
		u.log.Error(err.Error())
		responses.SendErrorResponse(w,http.StatusInternalServerError, err)
		return
	}
}

// swagger:route POST /api/v1/user/signup User signup
// Creates a new User
// responses:
//  200: safeUser
func (u *UserHandler) Registration(w http.ResponseWriter, r *http.Request) {

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	u.UserUseCase.SetDefaultAvatar(&user)
	if err != nil {
		u.log.Error(err.Error())
		responses.SendErrorResponse(w,http.StatusBadRequest,err)
		return
	}

	usr, err := u.UserUseCase.Add(user)
	if err != nil {
		u.log.Error(err.Error())
		responses.SendErrorResponse(w,http.StatusInternalServerError, err)
		return
	}

	sessionID, err := u.SessionsUseCase.AddToken(usr.ID)
	if err != nil {
		u.log.Error(err.Error())
		responses.SendErrorResponse(w,http.StatusInternalServerError, err)
		return
	}

	safeUser := models.SafeUser{ID: usr.ID, Username: usr.Username, Avatar: usr.Avatar, Email: usr.Email}

	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    sessionID,
		Path:     "/",
		HttpOnly: true,
		Expires:  time.Now().Add( configs.CookieLifeTime),

	})
	responses.SendOkResponse(w, safeUser)
}

// swagger:route POST /api/v1/user/signin User signin
// user auth with coockie
// responses:
//  200: safeUser
func (u *UserHandler) Auth(w http.ResponseWriter, r *http.Request) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		u.log.Error(err.Error())
		responses.SendErrorResponse(w,http.StatusBadRequest,err)
		return
	}

	usr, err := u.UserUseCase.GetByUserName(user.Username)
	if err != nil {
		u.log.Error(err.Error())
		responses.SendErrorResponse(w,http.StatusInternalServerError, err)
		return
	}
	err = u.UserUseCase.ComparePassword(user.Password,usr.Password)

	if err != nil {
		u.log.Error(err.Error())
		responses.SendErrorResponse(w,http.StatusUnauthorized, err)
		return
	}

	sessionID, err := u.SessionsUseCase.AddToken(usr.ID)
	if err != nil {
		u.log.Error(err.Error())
		responses.SendErrorResponse(w,http.StatusInternalServerError, err)
		return
	}
	safeUser := models.SafeUser{ID: usr.ID, Username: usr.Username, Avatar: usr.Avatar, Email: usr.Email}

	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    sessionID,
		Path:     "/",
		HttpOnly: true,
		Expires:  time.Now().Add( configs.CookieLifeTime),

	})
	responses.SendOkResponse(w, safeUser)
}

// swagger:route GET /api/v1/user User user
// Get current safe user
// responses:
//  200: safeUser
func (u *UserHandler) UserHandler(w http.ResponseWriter, r *http.Request) {
	usr, ok := r.Context().Value("User").(models.User)
	if !ok {
		responses.SendErrorResponse(w,http.StatusUnauthorized,errors.New("User Unauthorized"))
		return
	}
	if u.UserUseCase.CheckEmpty(usr) {
		responses.SendErrorResponse(w,http.StatusUnauthorized,errors.New("User Unauthorized"))
		return
	}

	safeUser := models.SafeUser{ID: usr.ID, Username: usr.Username, Avatar: usr.Avatar, Email: usr.Email}
	responses.SendOkResponse(w, safeUser)
}

// swagger:route POST /api/v1/user/signout User signout
// sign out current user and delete cookie session
func (u *UserHandler) SignOut(w http.ResponseWriter, r *http.Request) {

	c, _ := r.Cookie("session_token")
	if c != nil {

		err := u.SessionsUseCase.DeleteSession(c.Value)
		if err != nil {
			u.log.Error(err.Error())
			responses.SendErrorResponse(w,http.StatusInternalServerError, err)
			return
		}
		c.Expires = time.Now().AddDate(0, 0, -1)
		c.Path = "/"
		http.SetCookie(w, c)
	}
}

