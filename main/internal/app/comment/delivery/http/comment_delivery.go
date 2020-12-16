package commentDelivery

import (
	"errors"
	"mime/multipart"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	packageConfig "github.com/go-park-mail-ru/2020_2_JMickhs/package/configs"

	"github.com/mailru/easyjson"

	"github.com/go-park-mail-ru/2020_2_JMickhs/main/configs"
	"github.com/go-park-mail-ru/2020_2_JMickhs/main/internal/app/comment"
	commModel "github.com/go-park-mail-ru/2020_2_JMickhs/main/internal/app/comment/models"
	customerror "github.com/go-park-mail-ru/2020_2_JMickhs/package/error"
	"github.com/go-park-mail-ru/2020_2_JMickhs/package/middlewareApi"

	"github.com/go-park-mail-ru/2020_2_JMickhs/package/clientError"
	"github.com/go-park-mail-ru/2020_2_JMickhs/package/logger"
	"github.com/go-park-mail-ru/2020_2_JMickhs/package/responses"

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

	r.Path("/api/v1/comments/photos").Queries("id", "{id:[0-9]+}").HandlerFunc(handler.GetPhotos).Methods("GET")
	r.HandleFunc("/api/v1/comments", middlewareApi.CheckCSRFOnHandler(handler.AddComment)).Methods("POST")
	r.HandleFunc("/api/v1/comments", middlewareApi.CheckCSRFOnHandler(handler.UpdateComment)).Methods("PUT")
	r.HandleFunc("/api/v1/comments/{id:[0-9]+}", middlewareApi.CheckCSRFOnHandler(handler.DeleteComment)).Methods("DELETE")
	r.Path("/api/v1/comments").Queries("id", "{id:[0-9]+}", "limit", "{limit:[0-9]+}", "offset", "{from:[0-9]+}").
		HandlerFunc(handler.ListComments).Methods("GET")
}

func (hh *CommentHandler) DetectFileContentType(file multipart.File) (string, error) {
	fileHeader := make([]byte, 512)
	contentType := ""
	if _, err := file.Read(fileHeader); err != nil {
		return contentType, customerror.NewCustomError(err, clientError.BadRequest, 1)
	}

	if _, err := file.Seek(0, 0); err != nil {
		return contentType, customerror.NewCustomError(err, clientError.BadRequest, 1)
	}
	contentTypeStr := http.DetectContentType(fileHeader)
	contentType = strings.Split(contentTypeStr, "/")[1]
	if contentType != "jpg" && contentType != "png" && contentType != "jpeg" {
		return contentType, customerror.NewCustomError(errors.New("Wrong file type"), clientError.UnsupportedMediaType, 1)
	}
	return contentType, nil
}

// swagger:route GET /api/v1/comments/photos comment Photos
// GetList of photos
// responses:
//  200: photos
//  400: badrequest
func (ch *CommentHandler) GetPhotos(w http.ResponseWriter, r *http.Request) {

	hotelID := r.FormValue("id")

	photos, err := ch.CommentUseCase.GetPhotos(hotelID)
	if err != nil {
		customerror.PostError(w, r, ch.log, err, nil)
		return
	}
	responses.SendDataResponse(w, photos)
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

	userID, ok := r.Context().Value(packageConfig.RequestUserID).(int)
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
	userID, ok := r.Context().Value(packageConfig.RequestUserID).(int)
	if !ok {
		customerror.PostError(w, r, ch.log, errors.New("user unauthorized"), clientError.Unauthorizied)
		return
	}

	comment := commModel.Comment{}

	err := r.ParseMultipartForm(100 * configs.MB)
	if err != nil {
		customerror.PostError(w, r, ch.log, err, clientError.BadRequest)
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, 100*configs.MB)

	err = easyjson.Unmarshal([]byte(r.FormValue("jsonData")), &comment)
	if err != nil {
		customerror.PostError(w, r, ch.log, err, clientError.BadRequest)
		return
	}

	photos := r.MultipartForm.File["photos"]
	for _, photo := range photos {
		file, err := photo.Open()
		if err != nil {
			customerror.PostError(w, r, ch.log, err, clientError.BadRequest)
			return
		}
		fileType, err := ch.DetectFileContentType(file)
		if err != nil {
			customerror.PostError(w, r, ch.log, err, nil)
			return
		}

		err = ch.CommentUseCase.UploadPhoto(&comment, file, fileType)
		if err != nil {
			customerror.PostError(w, r, ch.log, err, nil)
			return
		}
		file.Close()
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

	userID, ok := r.Context().Value(packageConfig.RequestUserID).(int)
	if !ok {
		customerror.PostError(w, r, ch.log, errors.New("user unauthorized"), clientError.Unauthorizied)
		return
	}

	comment := commModel.Comment{}

	err := r.ParseMultipartForm(100 * configs.MB)
	if err != nil {
		customerror.PostError(w, r, ch.log, err, clientError.BadRequest)
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, 100*configs.MB)

	err = easyjson.Unmarshal([]byte(r.FormValue("jsonData")), &comment)
	if err != nil {
		customerror.PostError(w, r, ch.log, err, clientError.BadRequest)
		return
	}

	photos := r.MultipartForm.File["photos"]
	if len(photos) > 0 {
		err := ch.CommentUseCase.DeletePhotos(comment)
		if err != nil {
			customerror.PostError(w, r, ch.log, err, clientError.BadRequest)
			return
		}
	}
	for _, photo := range photos {
		file, err := photo.Open()
		if err != nil {
			customerror.PostError(w, r, ch.log, err, clientError.BadRequest)
			return
		}
		fileType, err := ch.DetectFileContentType(file)
		if err != nil {
			customerror.PostError(w, r, ch.log, err, nil)
			return
		}

		err = ch.CommentUseCase.UploadPhoto(&comment, file, fileType)
		if err != nil {
			customerror.PostError(w, r, ch.log, err, nil)
			return
		}
		file.Close()
	}
	comment.UserID = userID
	comm, err := ch.CommentUseCase.UpdateComment(comment)

	if err != nil {
		customerror.PostError(w, r, ch.log, err, nil)
		return
	}

	responses.SendDataResponse(w, comm)
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

	_, ok := r.Context().Value(packageConfig.RequestUserID).(int)
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
		Path:   "/api/v1/comments",
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
