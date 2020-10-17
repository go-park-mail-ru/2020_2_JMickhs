package commentDelivery

import (
	"encoding/json"
	"net/http"
	"strconv"

	customerror "github.com/go-park-mail-ru/2020_2_JMickhs/internal/pkg/error"

	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/pkg/logger"

	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/app/comment"
	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/app/comment/models"
	permissions "github.com/go-park-mail-ru/2020_2_JMickhs/internal/pkg/permission"
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

	r.HandleFunc("/api/v1/comments", permissions.CheckCSRF(handler.AddComment)).Methods("POST")
	r.HandleFunc("/api/v1/comments", permissions.CheckCSRF(handler.UpdateComment)).Methods("PUT")
	r.HandleFunc("/api/v1/comments/{id:[0-9]+}", permissions.CheckCSRF(handler.DeleteComment)).Methods("DELETE")
	r.Path("/api/v1/comments").Queries("id", "{id:[0-9]+}", "from", "{from:[0-9]+}").
		HandlerFunc(permissions.SetCSRF(handler.ListComments)).Methods("GET")
}

// swagger:route GET /api/v1/comments comment comments
// GetList of comments
// responses:
//  200: comments
//  400: badrequest
func (ch *CommentHandler) ListComments(w http.ResponseWriter, r *http.Request) {

	from := r.FormValue("from")
	startId, err := strconv.Atoi(from)

	if err != nil {
		err = customerror.NewCustomError(err.Error(), http.StatusBadRequest)
		ch.log.LogError(r.Context(), err)
		responses.SendErrorResponse(w, customerror.ParseCode(err))
		return
	}

	id := r.FormValue("id")
	hotelId, err := strconv.Atoi(id)
	if err != nil {
		err = customerror.NewCustomError(err.Error(), http.StatusBadRequest)
		ch.log.LogError(r.Context(), err)
		responses.SendErrorResponse(w, customerror.ParseCode(err))
		return
	}

	comments, err := ch.CommentUseCase.GetComments(hotelId, startId)

	if err != nil {
		ch.log.LogError(r.Context(), err)
		responses.SendErrorResponse(w, customerror.ParseCode(err))
		return
	}

	responses.SendDataResponse(w, comments)
}

// swagger:route POST /api/v1/comments comment AddComment
// responses:
//  200: AddComment
//  403: Forbidden
//  400: badrequest
func (ch *CommentHandler) AddComment(w http.ResponseWriter, r *http.Request) {
	comment := models.Comment{}
	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		err = customerror.NewCustomError(err.Error(), http.StatusBadRequest)
		ch.log.LogError(r.Context(), err)
		responses.SendErrorResponse(w, customerror.ParseCode(err))
		return
	}

	comm, err := ch.CommentUseCase.AddComment(comment)

	if err != nil {
		ch.log.LogError(r.Context(), err)
		responses.SendErrorResponse(w, customerror.ParseCode(err))
		return
	}

	responses.SendDataResponse(w, comm)
}

// swagger:route PUT /api/v1/comments comment UpdateComment
// responses:
// 403: Forbidden
// 400: badrequest
func (ch *CommentHandler) UpdateComment(w http.ResponseWriter, r *http.Request) {
	comment := models.Comment{}
	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		err = customerror.NewCustomError(err.Error(), http.StatusBadRequest)
		ch.log.LogError(r.Context(), err)
		responses.SendErrorResponse(w, customerror.ParseCode(err))
		return
	}

	err = ch.CommentUseCase.UpdateComment(comment)

	if err != nil {
		ch.log.LogError(r.Context(), err)
		responses.SendErrorResponse(w, customerror.ParseCode(err))
		return
	}
	responses.SendOkResponse(w)
}

// swagger:route DELETE /api/v1/comments/{id} comment DeleteComment
// responses:
//  403: Forbidden
//  400: badrequest
func (ch *CommentHandler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		err = customerror.NewCustomError(err.Error(), http.StatusBadRequest)
		ch.log.LogError(r.Context(), err)
		responses.SendErrorResponse(w, customerror.ParseCode(err))
		return
	}

	err = ch.CommentUseCase.DeleteComment(id)

	if err != nil {
		ch.log.LogError(r.Context(), err)
		responses.SendErrorResponse(w, customerror.ParseCode(err))
		return
	}

	responses.SendOkResponse(w)
}
