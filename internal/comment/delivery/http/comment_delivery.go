package commentDelivery

import (
	"encoding/json"
	"net/http"
	"strconv"

	customerror "github.com/go-park-mail-ru/2020_2_JMickhs/internal/error"

	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/logger"

	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/comment"
	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/comment/models"
	permissions "github.com/go-park-mail-ru/2020_2_JMickhs/internal/permission"
	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/responses"
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
func (ch *CommentHandler) ListComments(w http.ResponseWriter, r *http.Request) {

	from := r.FormValue("from")
	startId, err := strconv.Atoi(from)

	if err != nil {
		err = customerror.NewCustomError(err.Error())
		ch.log.LogError(r.Context(), err)
		responses.SendErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	id := r.FormValue("id")
	hotelId, err := strconv.Atoi(id)
	if err != nil {
		err = customerror.NewCustomError(err.Error())
		ch.log.LogError(r.Context(), err)
		responses.SendErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	comments, err := ch.CommentUseCase.GetComments(hotelId, startId)

	if err != nil {
		ch.log.LogError(r.Context(), err)
		responses.SendErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	responses.SendOkResponse(w, comments)
}

// swagger:route POST /api/v1/comments comment AddComment
// responses:
//  200: AddComment
func (ch *CommentHandler) AddComment(w http.ResponseWriter, r *http.Request) {
	comment := models.Comment{}
	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		err = customerror.NewCustomError(err.Error())
		ch.log.LogError(r.Context(), err)
		responses.SendErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	comm, err := ch.CommentUseCase.AddComment(comment)

	if err != nil {
		ch.log.LogError(r.Context(), err)
		responses.SendErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	responses.SendOkResponse(w, comm)
}

// swagger:route PUT /api/v1/comments comment UpdateComment
func (ch *CommentHandler) UpdateComment(w http.ResponseWriter, r *http.Request) {
	comment := models.Comment{}
	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		err = customerror.NewCustomError(err.Error())
		ch.log.LogError(r.Context(), err)
		responses.SendErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	err = ch.CommentUseCase.UpdateComment(comment)

	if err != nil {
		ch.log.LogError(r.Context(), err)
		responses.SendErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
}

// swagger:route DELETE /api/v1/comments/{id} comment DeleteComment
func (ch *CommentHandler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		err = customerror.NewCustomError(err.Error())
		ch.log.LogError(r.Context(), err)
		responses.SendErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	err = ch.CommentUseCase.DeleteComment(id)

	if err != nil {
		ch.log.LogError(r.Context(), err)
		responses.SendErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

}
