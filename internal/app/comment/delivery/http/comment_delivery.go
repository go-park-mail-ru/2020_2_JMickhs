package commentDelivery

import (
	"encoding/json"
	"errors"
	middlewareApi "github.com/go-park-mail-ru/2020_2_JMickhs/internal/app/middleware"
	"net/http"
	"strconv"

	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/pkg/clientError"

	commModel "github.com/go-park-mail-ru/2020_2_JMickhs/internal/app/comment/models"

	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/app/user/models"

	"github.com/go-park-mail-ru/2020_2_JMickhs/configs"

	customerror "github.com/go-park-mail-ru/2020_2_JMickhs/internal/pkg/error"

	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/pkg/logger"

	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/app/comment"
	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/pkg/responses"
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
	r.Path("/api/v1/comments").Queries("id", "{id:[0-9]+}", "limit", "{limit:[0-9]+}","offset", "{from:[0-9]+}").
		HandlerFunc(handler.ListComments).Methods("GET")
}

// swagger:route GET /api/v1/comments comment comments
// GetList of comments
// responses:
//  200: comments
//  400: badrequest
func (ch *CommentHandler) ListComments(w http.ResponseWriter, r *http.Request) {

	limit := r.FormValue("limit")
	offset := r.FormValue("offset")
	hotelID := r.FormValue("id")

	var user_id int
	user, ok := r.Context().Value(configs.RequestUser).(models.User)
	if !ok {
		user_id = -1
	}else {
		user_id = user.ID
	}

	comments, err := ch.CommentUseCase.GetComments(hotelID,limit,offset ,user_id)

	if err != nil {
		customerror.PostError(w, r, ch.log, err, nil)
		return
	}

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
	usr, ok := r.Context().Value(configs.RequestUser).(models.User)
	if !ok {
		customerror.PostError(w, r, ch.log, errors.New("user unauthorized"), clientError.Unauthorizied)
		return
	}
	comment.UserID = usr.ID
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

	usr, ok := r.Context().Value(configs.RequestUser).(models.User)
	if !ok {
		customerror.PostError(w, r, ch.log, errors.New("user unauthorized"), clientError.Unauthorizied)
		return
	}
	comment.UserID = usr.ID

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

	_, ok := r.Context().Value(configs.RequestUser).(models.User)
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
