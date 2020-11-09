package commentDelivery

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/go-park-mail-ru/2020_2_JMickhs/configs"
	comment_mock "github.com/go-park-mail-ru/2020_2_JMickhs/internal/app/comment/mocks"
	commModel "github.com/go-park-mail-ru/2020_2_JMickhs/internal/app/comment/models"
	paginationModel "github.com/go-park-mail-ru/2020_2_JMickhs/internal/app/paginator/model"
	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/app/user/models"
	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/pkg/clientError"
	customerror "github.com/go-park-mail-ru/2020_2_JMickhs/internal/pkg/error"
	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/pkg/logger"
	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/pkg/responses"
	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/pkg/serverError"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/mitchellh/mapstructure"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestCommentHandler_ListComments(t *testing.T) {

	testComments := []commModel.FullCommentInfo{
		{2,1,2,"kek",2,"src/kek.jpg","kostik","20-10-2000"},
		{2,2,2,"kekw",2,"src/kek.jpg","kostik","20-10-2000"},
	}

	testUser := models.User{ID: 2, Username: "kostik", Email: "sdfs@mail.ru", Password: "12345", Avatar: "kek/img.jpeg"}
	paginfo := paginationModel.PaginationInfo{ItemsCount: 56, NextLink: "api/v1/comments/?id=3&limit=1&offset=57",
		PrevLink:"api/v1/comments/?id=3&limit=1&offset=1"}
	commentsTest := commModel.Comments{Comments: testComments,Info: paginfo}

	t.Run("GetComments", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockCUseCase := comment_mock.NewMockUsecase(ctrl)

		mockCUseCase.EXPECT().
			GetComments("2","1","0",2).
			Return(commentsTest, nil)


		req, err := http.NewRequest("GET", "/api/v1/comments/?id=2&limit=1&offset=0", nil)
		assert.NoError(t, err)


		req = req.WithContext(context.WithValue(req.Context(), configs.RequestUser, testUser))
		rec := httptest.NewRecorder()
		handler := CommentHandler{
			CommentUseCase: mockCUseCase,
		}

		handler.ListComments(rec, req)
		resp := rec.Result()
		comments := []commModel.FullCommentInfo{}

		body, err := ioutil.ReadAll(resp.Body)
		response := responses.HttpResponse{}

		err = json.Unmarshal(body, &response)
		assert.NoError(t, err)
		err = mapstructure.Decode(response.Data.(map[string]interface{})["comments"], &comments)
		assert.NoError(t, err)

		assert.Equal(t, comments,testComments)
		assert.Equal(t, http.StatusOK, response.Code)
	})

	t.Run("GetCommentsErr1", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockCUseCase := comment_mock.NewMockUsecase(ctrl)

		mockCUseCase.EXPECT().
			GetComments("2","1","0",2).
			Return(commentsTest, customerror.NewCustomError(errors.New(""),serverError.ServerInternalError,1))


		req, err := http.NewRequest("GET", "/api/v1/comments/?id=2&limit=1&offset=0", nil)
		assert.NoError(t, err)


		req = req.WithContext(context.WithValue(req.Context(), configs.RequestUser, testUser))
		rec := httptest.NewRecorder()
		handler := CommentHandler{
			CommentUseCase: mockCUseCase,
			log: logger.NewLogger(os.Stdout),
		}

		handler.ListComments(rec, req)
		resp := rec.Result()

		body, err := ioutil.ReadAll(resp.Body)
		response := responses.HttpResponse{}

		err = json.Unmarshal(body, &response)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusInternalServerError, response.Code)
	})
}

func TestCommentHandler_AddComment(t *testing.T) {

	testComment := commModel.Comment{
		2,1,2,"kek",2,"20-10-2000",
	}

	testUser := models.User{ID: 2, Username: "kostik", Email: "sdfs@mail.ru", Password: "12345", Avatar: "kek/img.jpeg"}
	newRate := commModel.NewRate{Comment: testComment,Rate: 3}
	t.Run("AddComment", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockCUseCase := comment_mock.NewMockUsecase(ctrl)

		mockCUseCase.EXPECT().
			AddComment(testComment).
			Return(newRate, nil)

		bodys, _ := json.Marshal(testComment)
		req, err := http.NewRequest("POST", "/api/v1/comments", bytes.NewBuffer(bodys))
		assert.NoError(t, err)


		req = req.WithContext(context.WithValue(req.Context(), configs.RequestUser, testUser))
		rec := httptest.NewRecorder()
		handler := CommentHandler{
			CommentUseCase: mockCUseCase,
			log: logger.NewLogger(os.Stdout),
		}

		handler.AddComment(rec,req)
		resp := rec.Result()
		comment := commModel.NewRate{}

		body, err := ioutil.ReadAll(resp.Body)
		response := responses.HttpResponse{}

		err = json.Unmarshal(body, &response)
		assert.NoError(t, err)
		err = mapstructure.Decode(response.Data.(map[string]interface{}), &comment)
		assert.NoError(t, err)

		assert.Equal(t, newRate,comment)
		assert.Equal(t, http.StatusOK, response.Code)
	})

	t.Run("AddCommentErr1", func(t *testing.T) {
		testComments := []commModel.FullCommentInfo{
			{2,1,2,"kek",2,"src/kek.jpg","kostik","20-10-2000"},
			{2,2,2,"kekw",2,"src/kek.jpg","kostik","20-10-2000"},
		}

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockCUseCase := comment_mock.NewMockUsecase(ctrl)


		bodys, _ := json.Marshal(testComments)
		req, err := http.NewRequest("POST", "/api/v1/comments", bytes.NewBuffer(bodys))
		assert.NoError(t, err)


		req = req.WithContext(context.WithValue(req.Context(), configs.RequestUser, testUser))
		rec := httptest.NewRecorder()
		handler := CommentHandler{
			CommentUseCase: mockCUseCase,
			log: logger.NewLogger(os.Stdout),
		}

		handler.AddComment(rec,req)
		resp := rec.Result()


		body, err := ioutil.ReadAll(resp.Body)
		response := responses.HttpResponse{}

		err = json.Unmarshal(body, &response)
		assert.NoError(t, err)

		assert.Equal(t, clientError.BadRequest, response.Code)
	})

	t.Run("AddCommentErr2", func(t *testing.T) {

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockCUseCase := comment_mock.NewMockUsecase(ctrl)


		bodys, _ := json.Marshal(testComment)
		req, err := http.NewRequest("POST", "/api/v1/comments", bytes.NewBuffer(bodys))
		assert.NoError(t, err)

		rec := httptest.NewRecorder()
		handler := CommentHandler{
			CommentUseCase: mockCUseCase,
			log: logger.NewLogger(os.Stdout),
		}

		handler.AddComment(rec,req)
		resp := rec.Result()


		body, err := ioutil.ReadAll(resp.Body)
		response := responses.HttpResponse{}

		err = json.Unmarshal(body, &response)
		assert.NoError(t, err)

		assert.Equal(t, clientError.Unauthorizied, response.Code)
	})

	t.Run("AddCommentErr3", func(t *testing.T) {

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockCUseCase := comment_mock.NewMockUsecase(ctrl)

		mockCUseCase.EXPECT().
			AddComment(testComment).
			Return(newRate, customerror.NewCustomError(errors.New(""),serverError.ServerInternalError,1))

		bodys, _ := json.Marshal(testComment)
		req, err := http.NewRequest("POST", "/api/v1/comments", bytes.NewBuffer(bodys))

		assert.NoError(t, err)
		req = req.WithContext(context.WithValue(req.Context(), configs.RequestUser, testUser))
		rec := httptest.NewRecorder()
		handler := CommentHandler{
			CommentUseCase: mockCUseCase,
			log: logger.NewLogger(os.Stdout),
		}

		handler.AddComment(rec,req)
		resp := rec.Result()


		body, err := ioutil.ReadAll(resp.Body)
		response := responses.HttpResponse{}

		err = json.Unmarshal(body, &response)
		assert.NoError(t, err)

		assert.Equal(t, serverError.ServerInternalError, response.Code)
	})

}

func TestCommentHandler_UpdateComment(t *testing.T) {

	testComment := commModel.Comment{
		2,1,2,"kek",2,"20-10-2000",
	}

	testUser := models.User{ID: 2, Username: "kostik", Email: "sdfs@mail.ru", Password: "12345", Avatar: "kek/img.jpeg"}
	newRate := commModel.NewRate{Comment: testComment,Rate: 3}
	t.Run("UpdateComment", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockCUseCase := comment_mock.NewMockUsecase(ctrl)

		mockCUseCase.EXPECT().
			UpdateComment(testComment).
			Return(newRate, nil)

		bodys, _ := json.Marshal(testComment)
		req, err := http.NewRequest("POST", "/api/v1/comments", bytes.NewBuffer(bodys))
		assert.NoError(t, err)


		req = req.WithContext(context.WithValue(req.Context(), configs.RequestUser, testUser))
		rec := httptest.NewRecorder()
		handler := CommentHandler{
			CommentUseCase: mockCUseCase,
			log: logger.NewLogger(os.Stdout),
		}

		handler.UpdateComment(rec,req)
		resp := rec.Result()
		comment := commModel.NewRate{}

		body, err := ioutil.ReadAll(resp.Body)
		response := responses.HttpResponse{}

		err = json.Unmarshal(body, &response)
		assert.NoError(t, err)
		err = mapstructure.Decode(response.Data.(map[string]interface{}), &comment)
		assert.NoError(t, err)

		assert.Equal(t, newRate,comment)
		assert.Equal(t, http.StatusOK, response.Code)
	})

	t.Run("UpdateCommentErr1", func(t *testing.T) {
		testComments := []commModel.FullCommentInfo{
			{2,1,2,"kek",2,"src/kek.jpg","kostik","20-10-2000"},
			{2,2,2,"kekw",2,"src/kek.jpg","kostik","20-10-2000"},
		}

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockCUseCase := comment_mock.NewMockUsecase(ctrl)


		bodys, _ := json.Marshal(testComments)
		req, err := http.NewRequest("POST", "/api/v1/comments", bytes.NewBuffer(bodys))
		assert.NoError(t, err)


		req = req.WithContext(context.WithValue(req.Context(), configs.RequestUser, testUser))
		rec := httptest.NewRecorder()
		handler := CommentHandler{
			CommentUseCase: mockCUseCase,
			log: logger.NewLogger(os.Stdout),
		}

		handler.UpdateComment(rec,req)
		resp := rec.Result()


		body, err := ioutil.ReadAll(resp.Body)
		response := responses.HttpResponse{}

		err = json.Unmarshal(body, &response)
		assert.NoError(t, err)

		assert.Equal(t, clientError.BadRequest, response.Code)
	})

	t.Run("UpdateCommentErr2", func(t *testing.T) {

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockCUseCase := comment_mock.NewMockUsecase(ctrl)


		bodys, _ := json.Marshal(testComment)
		req, err := http.NewRequest("POST", "/api/v1/comments", bytes.NewBuffer(bodys))
		assert.NoError(t, err)

		rec := httptest.NewRecorder()
		handler := CommentHandler{
			CommentUseCase: mockCUseCase,
			log: logger.NewLogger(os.Stdout),
		}

		handler.UpdateComment(rec,req)
		resp := rec.Result()


		body, err := ioutil.ReadAll(resp.Body)
		response := responses.HttpResponse{}

		err = json.Unmarshal(body, &response)
		assert.NoError(t, err)

		assert.Equal(t, clientError.Unauthorizied, response.Code)
	})

	t.Run("UpdateCommentErr3", func(t *testing.T) {

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockCUseCase := comment_mock.NewMockUsecase(ctrl)

		mockCUseCase.EXPECT().
			UpdateComment(testComment).
			Return(newRate, customerror.NewCustomError(errors.New(""),serverError.ServerInternalError,1))

		bodys, _ := json.Marshal(testComment)
		req, err := http.NewRequest("POST", "/api/v1/comments", bytes.NewBuffer(bodys))

		assert.NoError(t, err)
		req = req.WithContext(context.WithValue(req.Context(), configs.RequestUser, testUser))
		rec := httptest.NewRecorder()
		handler := CommentHandler{
			CommentUseCase: mockCUseCase,
			log: logger.NewLogger(os.Stdout),
		}

		handler.UpdateComment(rec,req)
		resp := rec.Result()


		body, err := ioutil.ReadAll(resp.Body)
		response := responses.HttpResponse{}

		err = json.Unmarshal(body, &response)
		assert.NoError(t, err)

		assert.Equal(t, serverError.ServerInternalError, response.Code)
	})

}

func TestCommentHandler_DeleteComment(t *testing.T) {
	testUser := models.User{ID: 2, Username: "kostik", Email: "sdfs@mail.ru", Password: "12345", Avatar: "kek/img.jpeg"}

	t.Run("DeleteComment", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockCUseCase := comment_mock.NewMockUsecase(ctrl)

		mockCUseCase.EXPECT().
			DeleteComment(1).
			Return(nil)

		req, err := http.NewRequest("DELETE", "/api/v1/comments/1", nil)
		assert.NoError(t, err)


		req = req.WithContext(context.WithValue(req.Context(), configs.RequestUser, testUser))
		req = mux.SetURLVars(req, map[string]string{
			"id": "1",
		})
		rec := httptest.NewRecorder()
		handler := CommentHandler{
			CommentUseCase: mockCUseCase,
			log: logger.NewLogger(os.Stdout),
		}

		handler.DeleteComment(rec,req)
		resp := rec.Result()

		body, err := ioutil.ReadAll(resp.Body)
		response := responses.HttpResponse{}

		err = json.Unmarshal(body, &response)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusOK, response.Code)
	})

	t.Run("DeleteCommentErr1", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockCUseCase := comment_mock.NewMockUsecase(ctrl)

		req, err := http.NewRequest("DELETE", "/api/v1/comments/1", nil)
		assert.NoError(t, err)


		req = req.WithContext(context.WithValue(req.Context(), configs.RequestUser, testUser))
		rec := httptest.NewRecorder()
		handler := CommentHandler{
			CommentUseCase: mockCUseCase,
			log: logger.NewLogger(os.Stdout),
		}

		handler.DeleteComment(rec,req)
		resp := rec.Result()

		body, err := ioutil.ReadAll(resp.Body)
		response := responses.HttpResponse{}

		err = json.Unmarshal(body, &response)
		assert.NoError(t, err)

		assert.Equal(t, clientError.BadRequest, response.Code)
	})

	t.Run("DeleteCommentErr2", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockCUseCase := comment_mock.NewMockUsecase(ctrl)

		req, err := http.NewRequest("DELETE", "/api/v1/comments/1", nil)
		assert.NoError(t, err)

		req = mux.SetURLVars(req, map[string]string{
			"id": "1",
		})
		rec := httptest.NewRecorder()
		handler := CommentHandler{
			CommentUseCase: mockCUseCase,
			log: logger.NewLogger(os.Stdout),
		}

		handler.DeleteComment(rec,req)
		resp := rec.Result()

		body, err := ioutil.ReadAll(resp.Body)
		response := responses.HttpResponse{}

		err = json.Unmarshal(body, &response)
		assert.NoError(t, err)

		assert.Equal(t, clientError.Unauthorizied, response.Code)
	})

	t.Run("DeleteCommentErr3", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockCUseCase := comment_mock.NewMockUsecase(ctrl)

		mockCUseCase.EXPECT().
			DeleteComment(1).
			Return(customerror.NewCustomError(errors.New(""),serverError.ServerInternalError,1))

		req, err := http.NewRequest("DELETE", "/api/v1/comments/1", nil)
		assert.NoError(t, err)


		req = req.WithContext(context.WithValue(req.Context(), configs.RequestUser, testUser))
		req = mux.SetURLVars(req, map[string]string{
			"id": "1",
		})
		rec := httptest.NewRecorder()
		handler := CommentHandler{
			CommentUseCase: mockCUseCase,
			log: logger.NewLogger(os.Stdout),
		}

		handler.DeleteComment(rec,req)
		resp := rec.Result()

		body, err := ioutil.ReadAll(resp.Body)
		response := responses.HttpResponse{}

		err = json.Unmarshal(body, &response)
		assert.NoError(t, err)

		assert.Equal(t, serverError.ServerInternalError, response.Code)
	})

}