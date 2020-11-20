package userHttpDelivery

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-park-mail-ru/2020_2_JMickhs/JMickhs_user/internal/user"

	"github.com/go-park-mail-ru/2020_2_JMickhs/JMickhs_user/configs"
	"github.com/go-park-mail-ru/2020_2_JMickhs/JMickhs_user/internal/user/models"

	customerror "github.com/go-park-mail-ru/2020_2_JMickhs/package/error"
	"github.com/go-park-mail-ru/2020_2_JMickhs/package/middlewareApi"
	sessionService "github.com/go-park-mail-ru/2020_2_JMickhs/package/proto/sessions"

	"github.com/go-park-mail-ru/2020_2_JMickhs/package/clientError"
	"github.com/go-park-mail-ru/2020_2_JMickhs/package/logger"
	"github.com/go-park-mail-ru/2020_2_JMickhs/package/responses"
	"github.com/go-park-mail-ru/2020_2_JMickhs/package/serverError"

	"github.com/gorilla/mux"
)

type UserHandler struct {
	UserUseCase      user.Usecase
	SessionsDelivery sessionService.AuthorizationServiceClient
	log              *logger.CustomLogger
}

func NewUserHandler(r *mux.Router, su sessionService.AuthorizationServiceClient, us user.Usecase, lg *logger.CustomLogger) {
	handler := UserHandler{
		UserUseCase:      us,
		SessionsDelivery: su,
		log:              lg,
	}

	r.HandleFunc("/api/v1/users", handler.Registration).Methods("POST")
	r.HandleFunc("/api/v1/users/sessions", handler.Auth).Methods("POST")
	r.HandleFunc("/api/v1/users", handler.UserHandler).Methods("GET")
	r.HandleFunc("/api/v1/users/sessions", handler.SignOut).Methods("DELETE")
	r.HandleFunc("/api/v1/users/credentials", middlewareApi.CheckCSRFOnHandler(handler.UpdateUser)).Methods("PUT")
	r.HandleFunc("/api/v1/users/avatar", middlewareApi.CheckCSRFOnHandler(handler.UpdateAvatar)).Methods("PUT")
	r.HandleFunc("/api/v1/users/password", middlewareApi.CheckCSRFOnHandler(handler.updatePassword)).Methods("PUT")
	r.HandleFunc("/api/v1/users/{id:[0-9]+}", handler.getAccInfo).Methods("GET")
	r.HandleFunc("/api/v1/csrf", handler.GetCsrf).Methods("GET")
}

// swagger:route GET /api/v1/users/{id}  Users userById
// Get info abous user by his id
// responses:
//  200: safeUser
//  400: badrequest
//  410:  gone
func (u *UserHandler) getAccInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	user, err := u.UserUseCase.GetUserByID(id)

	if err != nil {
		customerror.PostError(w, r, u.log, err, nil)
		return
	}

	safeUser := models.SafeUser{ID: user.ID, Username: user.Username, Avatar: user.Avatar, Email: user.Email}

	responses.SendDataResponse(w, safeUser)
}

// swagger:route PUT /api/v1/users/avatar Users avatar
// Update Avatar
// responses:
// 200: avatar
// 403: Forbidden
// 400: badrequest
// 401: unauthorizied
// 415: unsupport
func (u *UserHandler) UpdateAvatar(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(5 * configs.MB)

	r.Body = http.MaxBytesReader(w, r.Body, 5*configs.MB)

	file, _, err := r.FormFile("avatar")

	if err != nil {
		customerror.PostError(w, r, u.log, err, clientError.BadRequest)
		return
	}

	fileType, err := u.UserUseCase.CheckAvatar(file)
	if err != nil {
		customerror.PostError(w, r, u.log, err, clientError.BadRequest)
		return
	}

	usr, ok := r.Context().Value(configs.RequestUser).(models.User)
	if !ok {
		customerror.PostError(w, r, u.log, errors.New("Unauthorized"), clientError.Unauthorizied)
		return
	}

	path, err := u.UserUseCase.UploadAvatar(file, fileType, &usr)
	if err != nil {
		customerror.PostError(w, r, u.log, err, clientError.BadRequest)
		return
	}

	err = u.UserUseCase.UpdateAvatar(usr)
	if err != nil {
		customerror.PostError(w, r, u.log, err, clientError.BadRequest)
		return
	}

	responses.SendDataResponse(w, configs.S3Url+path)
}

// swagger:route PUT /api/v1/users/password Users password
// update password
// responses:
// 403: Forbidden
// 400: badrequest
// 401: unauthorizied
// 402: badCredentials
func (u *UserHandler) updatePassword(w http.ResponseWriter, r *http.Request) {

	var twoPass models.UpdatePassword
	err := json.NewDecoder(r.Body).Decode(&twoPass)
	if err != nil {
		customerror.PostError(w, r, u.log, err, clientError.BadRequest)
		return
	}

	usr, ok := r.Context().Value(configs.RequestUser).(models.User)
	if !ok {
		customerror.PostError(w, r, u.log, errors.New("Unauthorized"), clientError.Unauthorizied)
		return
	}

	err = u.UserUseCase.ComparePassword(twoPass.OldPassword, usr.Password)
	if err != nil {
		customerror.PostError(w, r, u.log, errors.New("wrong old password"), clientError.PaymentReq)
		return
	}
	usr.Password = twoPass.NewPassword
	err = u.UserUseCase.UpdatePassword(usr)
	if err != nil {
		customerror.PostError(w, r, u.log, err, nil)
		return
	}
	responses.SendOkResponse(w)
}

// swagger:route PUT /api/v1/users/credentials Users credentials
// Get data from form  which need to change and change user data
// responses:
// 403: Forbidden
// 400: badrequest
// 401: unauthorizied
// 409: conflict
func (u *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		customerror.PostError(w, r, u.log, err, clientError.BadRequest)
		return
	}

	usr, ok := r.Context().Value(configs.RequestUser).(models.User)
	if !ok {
		customerror.PostError(w, r, u.log, errors.New("Unauthorized"), clientError.Unauthorizied)
		return
	}
	user.ID = usr.ID

	err = u.UserUseCase.UpdateUser(user)
	if err != nil {
		customerror.PostError(w, r, u.log, err, nil)
		return
	}
	responses.SendOkResponse(w)
}

// swagger:route POST /api/v1/users Users signup
// Creates a new User
// responses:
//  200: safeUser
//  400: badrequest
//  409: conflict
func (u *UserHandler) Registration(w http.ResponseWriter, r *http.Request) {

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		customerror.PostError(w, r, u.log, err, clientError.BadRequest)
		return
	}

	usr, err := u.UserUseCase.Add(user)
	if err != nil {
		customerror.PostError(w, r, u.log, err, nil)
		return
	}

	sessionID, err := u.SessionsDelivery.CreateSession(context.Background(), &sessionService.UserID{UserID: int64(usr.ID)})
	if err != nil {
		customerror.PostError(w, r, u.log, err, serverError.ServerInternalError)
		return
	}

	safeUser := models.SafeUser{ID: usr.ID, Username: usr.Username, Avatar: usr.Avatar, Email: usr.Email}

	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    sessionID.SessionID,
		Path:     "/",
		HttpOnly: true,
		Expires:  time.Now().Add(configs.CookieLifeTime),
	})
	responses.SendDataResponse(w, safeUser)
}

// swagger:route POST /api/v1/users/sessions Sessions AddSessions
// user auth with coockie
// responses:
//  200: safeUser
//  400: badrequest
//  401: unauthorizied
func (u *UserHandler) Auth(w http.ResponseWriter, r *http.Request) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		customerror.PostError(w, r, u.log, err, clientError.BadRequest)
		return
	}

	usr, err := u.UserUseCase.GetByUserName(user.Username)
	if err != nil {
		customerror.PostError(w, r, u.log, err, nil)
		return
	}
	err = u.UserUseCase.ComparePassword(user.Password, usr.Password)

	if err != nil {
		customerror.PostError(w, r, u.log, err, clientError.Unauthorizied)
		return
	}

	sessionID, err := u.SessionsDelivery.CreateSession(r.Context(), &sessionService.UserID{UserID: int64(usr.ID)})
	fmt.Println("fd")
	if err != nil {
		customerror.PostError(w, r, u.log, err, nil)
		return
	}
	safeUser := models.SafeUser{ID: usr.ID, Username: usr.Username, Avatar: usr.Avatar, Email: usr.Email}

	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    sessionID.SessionID,
		Path:     "/",
		HttpOnly: true,
		Expires:  time.Now().Add(configs.CookieLifeTime),
	})
	responses.SendDataResponse(w, safeUser)
}

// swagger:route GET /api/v1/users Users user
// Get current safe user
// responses:
//  200: safeUser
//  401: unauthorizied
func (u *UserHandler) UserHandler(w http.ResponseWriter, r *http.Request) {
	usr, ok := r.Context().Value(configs.RequestUser).(models.User)
	if !ok {
		customerror.PostError(w, r, u.log, errors.New("user unothorizied"), clientError.Unauthorizied)
		return
	}
	if u.UserUseCase.CheckEmpty(usr) {
		customerror.PostError(w, r, u.log, errors.New("user unothorizied"), clientError.Unauthorizied)
		return
	}

	safeUser := models.SafeUser{ID: usr.ID, Username: usr.Username, Avatar: usr.Avatar, Email: usr.Email}
	responses.SendDataResponse(w, safeUser)
}

// swagger:route DELETE /api/v1/users/sessions Sessions DelSessions
// sign out current user and delete cookie session
func (u *UserHandler) SignOut(w http.ResponseWriter, r *http.Request) {

	c, _ := r.Cookie("session_token")
	if c != nil {
		_, err := u.SessionsDelivery.DeleteSession(r.Context(), &sessionService.SessionID{SessionID: c.Value})
		if err != nil {
			customerror.PostError(w, r, u.log, err, serverError.ServerInternalError)
			return
		}
		c.Expires = time.Now().AddDate(0, 0, -1)
		c.Path = "/"
		http.SetCookie(w, c)
		responses.SendOkResponse(w)
		return
	}
	responses.SendErrorResponse(w, clientError.BadRequest)
}

// swagger:route GET /api/v1/csrf Csrf Csrf
// get csrf token, token expire = 15 min
func (u *UserHandler) GetCsrf(w http.ResponseWriter, r *http.Request) {
	sId, ok := r.Context().Value(configs.SessionID).(string)
	if !ok {
		customerror.PostError(w, r, u.log, errors.New("Unauthorized"), clientError.Unauthorizied)
		return
	}
	token, err := u.SessionsDelivery.CreateCsrfToken(r.Context(), &sessionService.CsrfTokenInput{SessionID: sId, TimeStamp: time.Now().Unix()})
	if err != nil {
		customerror.PostError(w, r, u.log, err, serverError.ServerInternalError)
		return
	}
	w.Header().Set("Csrf", token.Token)
	responses.SendOkResponse(w)
}
