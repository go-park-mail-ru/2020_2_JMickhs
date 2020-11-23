package commentDelivery

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strconv"

	"github.com/spf13/viper"

	customerror "github.com/go-park-mail-ru/2020_2_JMickhs/package/error"
	"github.com/go-park-mail-ru/2020_2_JMickhs/package/middlewareApi"

	"github.com/go-park-mail-ru/2020_2_JMickhs/JMickhs_main/configs"
	commModel "github.com/go-park-mail-ru/2020_2_JMickhs/JMickhs_main/internal/app/comment/models"
	"github.com/go-park-mail-ru/2020_2_JMickhs/package/clientError"
	"github.com/go-park-mail-ru/2020_2_JMickhs/package/logger"
	"github.com/go-park-mail-ru/2020_2_JMickhs/package/responses"

	"github.com/go-park-mail-ru/2020_2_JMickhs/JMickhs_main/internal/app/comment"

	"github.com/gorilla/mux"
)

type CommentHandler struct {
	CommentUseCase comment.Usecase
	log            *logger.CustomLogger
}

func NewCommentHandler(r *mux.Router, hs comment.Usecase, lg *logger.CustomLogger) {
	handler := CommentHandler{
		CommentUseCase: hs,
		log:            lg,
	}

	r.HandleFunc("/api/v1/comments", middlewareApi.CheckCSRFOnHandler(handler.AddComment)).Methods("POST")
	r.HandleFunc("/api/v1/comments", middlewareApi.CheckCSRFOnHandler(handler.UpdateComment)).Methods("PUT")
	r.HandleFunc("/api/v1/comments/{id:[0-9]+}", middlewareApi.CheckCSRFOnHandler(handler.DeleteComment)).Methods("DELETE")
	r.Path("/api/v1/comments").Queries("id", "{id:[0-9]+}", "limit", "{limit:[0-9]+}", "offset", "{from:[0-9]+}").
		HandlerFunc(handler.ListComments).Methods("GET")
}

// swagger:route GET /api/v1/comments comment comments
// GetList of comments
// responses:
//  200: comments
//  400: badrequest
func (ch *CommentHandler) ListComments(w http.ResponseWriter, r *http.Request) {

	limit := r.FormValue("limit")
	offsetVar := r.FormValue("offset")
	hotelID := r.FormValue("id")

	var user_id int

	userID, ok := r.Context().Value(viper.GetString(configs.ConfigFields.RequestUserID)).(int)
	if !ok {
		user_id = -1
	} else {
		user_id = userID
	}

	count, comments, err := ch.CommentUseCase.GetComments(hotelID, limit, offsetVar, user_id)
	if err != nil {

		customerror.PostError(w, r, ch.log, err, nil)
		return
	}
	ch.NextPrevUrl(count, limit, offsetVar, hotelID, &comments)
	responses.SendDataResponse(w, comments)
}

// swagger:route POST /api/v1/comments comment AddComment
// add comment and rate hotel
// responses:
//  200: AddComment
//  403: Forbidden
//  400: badrequest
//  423: locked
func (ch *CommentHandler) AddComment(w http.ResponseWriter, r *http.Request) {
	comment := commModel.Comment{}

	err := json.NewDecoder(r.Body).Decode(&comment)

	if err != nil {
		customerror.PostError(w, r, ch.log, err, clientError.BadRequest)
		return
	}
	userID, ok := r.Context().Value(viper.GetString(configs.ConfigFields.RequestUserID)).(int)
	if !ok {
		customerror.PostError(w, r, ch.log, errors.New("user unauthorized"), clientError.Unauthorizied)
		return
	}
	comment.UserID = userID
	comm, err := ch.CommentUseCase.AddComment(comment)

	if err != nil {
		customerror.PostError(w, r, ch.log, err, nil)
		return
	}

	responses.SendDataResponse(w, comm)
}

// swagger:route PUT /api/v1/comments comment UpdateComment
// responses:
// 200: AddComment
// 403: Forbidden
// 400: badrequest
// 423: locked
func (ch *CommentHandler) UpdateComment(w http.ResponseWriter, r *http.Request) {
	comment := commModel.Comment{}
	err := json.NewDecoder(r.Body).Decode(&comment)

	if err != nil {
		customerror.PostError(w, r, ch.log, err, clientError.BadRequest)
		return
	}

	userID, ok := r.Context().Value(viper.GetString(configs.ConfigFields.RequestUserID)).(int)
	if !ok {
		customerror.PostError(w, r, ch.log, errors.New("user unauthorized"), clientError.Unauthorizied)
		return
	}
	comment.UserID = userID

	upComm, err := ch.CommentUseCase.UpdateComment(comment)

	if err != nil {
		customerror.PostError(w, r, ch.log, err, nil)
		return
	}

	responses.SendDataResponse(w, upComm)
}

// swagger:route DELETE /api/v1/comments/{id} comment DeleteComment
// responses:
//  403: Forbidden
//  400: badrequest
func (ch *CommentHandler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		customerror.PostError(w, r, ch.log, err, clientError.BadRequest)
		return
	}

	_, ok := r.Context().Value(viper.GetString(configs.ConfigFields.RequestUserID)).(int)
	if !ok {
		customerror.PostError(w, r, ch.log, errors.New("user unauthorized"), clientError.Unauthorizied)
		return
	}

	err = ch.CommentUseCase.DeleteComment(id)

	if err != nil {
		customerror.PostError(w, r, ch.log, err, nil)
		return
	}

	responses.SendOkResponse(w)
}

func (ch *CommentHandler) NextPrevUrl(count int, limit string, offsetVar string, hotelID string, comments *commModel.Comments) {
	url := url.URL{
		Host:   "hostelscan.ru:8080",
		Scheme: "https",
		Path:   "/api/v1/comments/",
	}

	query := url.Query()
	query.Set("id", hotelID)
	query.Set("limit", limit)

	lim, _ := strconv.Atoi(limit)
	offset, _ := strconv.Atoi(offsetVar)

	if lim == count && offset == 0 {
		query.Set("offset", "0")
		url.RawQuery = query.Encode()
		comments.Info.NextLink = url.Path + "?" + url.RawQuery
		comments.Info.PrevLink = url.Path + "?" + url.RawQuery
		return
	}
	if offset+lim >= count {
		query.Set("offset", strconv.Itoa(offset-lim))
		url.RawQuery = query.Encode()
		comments.Info.PrevLink = url.Path + "?" + url.RawQuery

		query.Set("offset", "0")
		url.RawQuery = query.Encode()
		comments.Info.NextLink = url.Path + "?" + url.RawQuery
	}
	if offset-lim < 0 {
		query.Set("offset", strconv.Itoa(offset+lim))
		url.RawQuery = query.Encode()
		comments.Info.NextLink = url.Path + "?" + url.RawQuery

		query.Set("offset", strconv.Itoa(count-lim))
		url.RawQuery = query.Encode()
		comments.Info.PrevLink = url.Path + "?" + url.RawQuery
	}
	if offset+lim < count && offset-lim >= 0 {
		query.Set("offset", strconv.Itoa(offset+lim))
		url.RawQuery = query.Encode()
		comments.Info.NextLink = url.Path + "?" + url.RawQuery

		query.Set("offset", strconv.Itoa(offset-lim))
		url.RawQuery = query.Encode()
		comments.Info.PrevLink = url.Path + "?" + url.RawQuery
	}
}
