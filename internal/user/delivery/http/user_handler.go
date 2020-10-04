package userDelivery

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-park-mail-ru/2020_2_JMickhs/configs"
	"net/http"
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

	r.HandleFunc("/api/v1/signup", permissions.SetCSRF(handler.Registration)).Methods("POST")
	r.HandleFunc("/api/v1/signin", permissions.SetCSRF(handler.Auth)).Methods("POST")
	r.HandleFunc("/api/v1/get_current_user", permissions.SetCSRF(handler.GetCurrentUser)).Methods("GET")
	r.HandleFunc("/api/v1/signout", handler.SignOut).Methods("POST")
	r.HandleFunc("/api/v1/updateUser", permissions.CheckCSRF(handler.UpdateUser)).Methods("PUT")
	r.HandleFunc("/api/v1/updateAvatar", permissions.CheckCSRF(handler.UpdateAvatar)).Methods("PUT")
	r.HandleFunc("/api/v1/updatePassword", permissions.CheckCSRF(handler.updatePassword)).Methods("PUT")
	r.HandleFunc("/api/v1/getAccInfo", permissions.SetCSRF(handler.getAccInfo)).Methods("GET")
}

// swagger:route GET /api/v1/getAccInfo getAccInfo get
// Get info abous user by his nickname
// responses:
//  200: safeUser
func (u *UserHandler) getAccInfo(w http.ResponseWriter, r *http.Request) {
	var name models.UserName

	err := json.NewDecoder(r.Body).Decode(&name)
	if err != nil {
		u.log.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := u.UserUseCase.GetByUserName(name.Username)
	if err != nil {
		u.log.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	safeUser := models.SafeUser{ID: user.ID, Username: user.Username, Avatar: user.Avatar, Email: user.Email}
	err = json.NewEncoder(w).Encode(safeUser)
	if err != nil {
		u.log.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}

// swagger:route PUT /api/v1/updateAvatar updateAvatar update
// Update Avatar
func (u *UserHandler) UpdateAvatar(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(32 << 20)
	file, info, err := r.FormFile("avatar")
	if err != nil {
		u.log.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fileType := strings.Split(info.Header.Get("Content-Type"),"/")

	defer file.Close()
	usr := r.Context().Value("User").(models.User)
	extension := fileType[1]
	if (extension != "jpg" && extension != "png" && extension != "jpeg"){
		u.log.Error(errors.New("bad type of file to add in static"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	u.UserUseCase.UploadAvatar(file, fileType[1] , &usr)

	err = u.UserUseCase.UpdateAvatar(usr)
	if err != nil {
		u.log.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// swagger:route PUT /api/v1/updatePassword updatePassword update
// update password
func (u *UserHandler) updatePassword(w http.ResponseWriter, r *http.Request) {

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		u.log.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(user.ID)
	fmt.Println(user.Password)
	err = u.UserUseCase.UpdatePassword(user)
	if err != nil {
		u.log.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// swagger:route PUT /api/v1/updateUser updateUser update
// Get data from form  which need to change and change user data
func (u *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		u.log.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = u.UserUseCase.UpdateUser(user)
	if err != nil {
		u.log.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// swagger:route POST /api/v1/signup signup signup
// Creates a new User
// responses:
//  200: user
func (u *UserHandler) Registration(w http.ResponseWriter, r *http.Request) {

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	u.UserUseCase.SetDefaultAvatar(&user)
	if err != nil {
		u.log.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	usr, err := u.UserUseCase.Add(user)
	if err != nil {
		u.log.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	sessionID, err := u.SessionsUseCase.AddToken(usr.ID)
	if err != nil {
		u.log.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
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

	err = json.NewEncoder(w).Encode(safeUser)
}

// swagger:route POST /api/v1/signin signin auth
// user auth with coockie
// responses:
//  200: user
func (u *UserHandler) Auth(w http.ResponseWriter, r *http.Request) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		u.log.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	usr, err := u.UserUseCase.GetByUserName(user.Username)
	if err != nil {
		u.log.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = u.UserUseCase.ComparePassword(user.Password,usr.Password)

	if err != nil {
		u.log.Error(err.Error())
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	sessionID, err := u.SessionsUseCase.AddToken(usr.ID)
	if err != nil {
		u.log.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
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

	err = json.NewEncoder(w).Encode(safeUser)
}

// swagger:route GET /api/v1/get_current_user GetCurrUser safeUser
// Get current safe user
// responses:
//  200: safeUser
func (u *UserHandler) GetCurrentUser(w http.ResponseWriter, r *http.Request) {
	usr := r.Context().Value("User").(models.User)
	if checkEmpty(usr) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	safeUser := models.SafeUser{ID: usr.ID, Username: usr.Username, Avatar: usr.Avatar, Email: usr.Email}
	err := json.NewEncoder(w).Encode(safeUser)
	if err != nil {
		u.log.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

// swagger:route GET /api/v1/signout SignOut out
// sign out current user and delete cookie session
func (u *UserHandler) SignOut(w http.ResponseWriter, r *http.Request) {

	c, _ := r.Cookie("session_token")
	fmt.Println(c.Value)
	if c != nil {
		u.SessionsUseCase.DeleteSession(c.Value)
		c.Expires = time.Now().AddDate(0, 0, -1)
		c.Path = "/"
		http.SetCookie(w, c)
	}
}

func checkEmpty(usr models.User) bool {
	empty := models.User{}
	return usr == empty
}
