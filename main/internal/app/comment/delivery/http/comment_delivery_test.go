package commentDelivery

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	packageConfig "github.com/go-park-mail-ru/2020_2_JMickhs/package/configs"

	customerror "github.com/go-park-mail-ru/2020_2_JMickhs/package/error"
	userService "github.com/go-park-mail-ru/2020_2_JMickhs/package/proto/user"
	"github.com/stretchr/testify/assert"

	comment_mock "github.com/go-park-mail-ru/2020_2_JMickhs/main/internal/app/comment/mocks"
	commModel "github.com/go-park-mail-ru/2020_2_JMickhs/main/internal/app/comment/models"
	paginationModel "github.com/go-park-mail-ru/2020_2_JMickhs/main/internal/app/paginator/model"
	"github.com/go-park-mail-ru/2020_2_JMickhs/package/clientError"
	"github.com/go-park-mail-ru/2020_2_JMickhs/package/logger"
	"github.com/go-park-mail-ru/2020_2_JMickhs/package/responses"
	"github.com/go-park-mail-ru/2020_2_JMickhs/package/serverError"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/mitchellh/mapstructure"
)

func TestCommentHandler_ListComments(t *testing.T) {

	testComments := []commModel.FullCommentInfo{
		{UserID: 2, CommID: 1, HotelID: 2, Message: "kek", Rating: 2, Avatar: "src/kek.jpg", Username: "kostik", Time: "20-10-2000"},
		{UserID: 2, CommID: 2, HotelID: 2, Message: "kekw", Rating: 2, Avatar: "src/kek.jpg", Username: "kostik", Time: "20-10-2000"},
	}
	testUser := userService.User{UserID: 2, Username: "kostik", Email: "sdfs@mail.ru", Avatar: "kek/img.jpeg"}
	paginfo := paginationModel.PaginationInfo{ItemsCount: 56, NextLink: "api/v1/comments/?id=2&limit=1&offset=57",
		PrevLink: "api/v1/comments/?id=2&limit=1&offset=1"}
	commentsTest := commModel.Comments{Comments: testComments, Info: paginfo}

	t.Run("GetComments", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockCUseCase := comment_mock.NewMockUsecase(ctrl)

		mockCUseCase.EXPECT().
			GetComments("2", "1", "0", 2).
			Return(56, commentsTest, nil)

		req, err := http.NewRequest("GET", "/api/v1/comments/?id=2&limit=1&offset=0", nil)
		assert.NoError(t, err)

		req = req.WithContext(context.WithValue(req.Context(), packageConfig.RequestUserID, int(testUser.UserID)))
		rec := httptest.NewRecorder()
		handler := CommentHandler{
			CommentUseCase: mockCUseCase,
		}

		handler.ListComments(rec, req)
		resp := rec.Result()
		comments := []commModel.FullCommentInfo{}

		body, _ := ioutil.ReadAll(resp.Body)
		response := responses.HttpResponse{}

		err = json.Unmarshal(body, &response)
		assert.NoError(t, err)
		err = mapstructure.Decode(response.Data.(map[string]interface{})["comments"], &comments)
		assert.NoError(t, err)

		assert.Equal(t, comments, testComments)
		assert.Equal(t, http.StatusOK, response.Code)
	})

	t.Run("GetCommentsErr1", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockCUseCase := comment_mock.NewMockUsecase(ctrl)

		mockCUseCase.EXPECT().
			GetComments("2", "1", "0", 2).
			Return(56, commentsTest, customerror.NewCustomError(errors.New(""), serverError.ServerInternalError, 1))

		req, err := http.NewRequest("GET", "/api/v1/comments/?id=2&limit=1&offset=0", nil)
		assert.NoError(t, err)

		req = req.WithContext(context.WithValue(req.Context(), packageConfig.RequestUserID, int(testUser.UserID)))
		rec := httptest.NewRecorder()
		handler := CommentHandler{
			CommentUseCase: mockCUseCase,
			log:            logger.NewLogger(os.Stdout),
		}

		handler.ListComments(rec, req)
		resp := rec.Result()

		body, _ := ioutil.ReadAll(resp.Body)
		response := responses.HttpResponse{}

		err = json.Unmarshal(body, &response)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusInternalServerError, response.Code)
	})
	t.Run("GetCommentsErr2", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockCUseCase := comment_mock.NewMockUsecase(ctrl)

		mockCUseCase.EXPECT().
			GetComments("2", "1", "1", 2).
			Return(56, commentsTest, nil)

		req, err := http.NewRequest("GET", "/api/v1/comments/?id=2&limit=1&offset=1", nil)
		assert.NoError(t, err)

		req = req.WithContext(context.WithValue(req.Context(), packageConfig.RequestUserID, int(testUser.UserID)))
		rec := httptest.NewRecorder()
		handler := CommentHandler{
			CommentUseCase: mockCUseCase,
			log:            logger.NewLogger(os.Stdout),
		}

		handler.ListComments(rec, req)
		resp := rec.Result()

		body, _ := ioutil.ReadAll(resp.Body)
		response := responses.HttpResponse{}

		err = json.Unmarshal(body, &response)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusOK, response.Code)
	})
	t.Run("GetCommentsErr3", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockCUseCase := comment_mock.NewMockUsecase(ctrl)

		mockCUseCase.EXPECT().
			GetComments("2", "1", "56", 2).
			Return(56, commentsTest, nil)

		req, err := http.NewRequest("GET", "/api/v1/comments/?id=2&limit=1&offset=56", nil)
		assert.NoError(t, err)

		req = req.WithContext(context.WithValue(req.Context(), packageConfig.RequestUserID, int(testUser.UserID)))
		rec := httptest.NewRecorder()
		handler := CommentHandler{
			CommentUseCase: mockCUseCase,
			log:            logger.NewLogger(os.Stdout),
		}

		handler.ListComments(rec, req)
		resp := rec.Result()

		body, _ := ioutil.ReadAll(resp.Body)
		response := responses.HttpResponse{}

		err = json.Unmarshal(body, &response)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusOK, response.Code)
	})
}

func TestCommentHandler_AddComment(t *testing.T) {

	testComment := commModel.Comment{
		UserID: 2, HotelID: 1, CommID: 2, Message: "kek", Rate: 2, Time: "20-10-2000",
	}

	testUser := userService.User{UserID: 2, Username: "kostik", Email: "sdfs@mail.ru", Avatar: "kek/img.jpeg"}
	newRate := commModel.NewRate{Comment: testComment, Rate: 3}
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

		req = req.WithContext(context.WithValue(req.Context(), packageConfig.RequestUserID, int(testUser.UserID)))
		rec := httptest.NewRecorder()
		handler := CommentHandler{
			CommentUseCase: mockCUseCase,
			log:            logger.NewLogger(os.Stdout),
		}

		handler.AddComment(rec, req)
		resp := rec.Result()
		comment := commModel.NewRate{}

		body, _ := ioutil.ReadAll(resp.Body)
		response := responses.HttpResponse{}

		err = json.Unmarshal(body, &response)
		assert.NoError(t, err)
		err = mapstructure.Decode(response.Data.(map[string]interface{}), &comment)
		assert.NoError(t, err)

		assert.Equal(t, newRate, comment)
		assert.Equal(t, http.StatusOK, response.Code)
	})

	t.Run("AddCommentErr1", func(t *testing.T) {
		testComments := []commModel.FullCommentInfo{
			{UserID: 2, CommID: 1, HotelID: 2, Message: "kek", Rating: 2, Avatar: "src/kek.jpg", Username: "kostik", Time: "20-10-2000"},
			{UserID: 2, CommID: 2, HotelID: 2, Message: "kekw", Rating: 2, Avatar: "src/kek.jpg", Username: "kostik", Time: "20-10-2000"},
		}
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockCUseCase := comment_mock.NewMockUsecase(ctrl)

		bodys, _ := json.Marshal(testComments)
		req, err := http.NewRequest("POST", "/api/v1/comments", bytes.NewBuffer(bodys))
		assert.NoError(t, err)

		req = req.WithContext(context.WithValue(req.Context(), packageConfig.RequestUserID, int(testUser.UserID)))
		rec := httptest.NewRecorder()
		handler := CommentHandler{
			CommentUseCase: mockCUseCase,
			log:            logger.NewLogger(os.Stdout),
		}

		handler.AddComment(rec, req)
		resp := rec.Result()

		body, _ := ioutil.ReadAll(resp.Body)
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
			log:            logger.NewLogger(os.Stdout),
		}

		handler.AddComment(rec, req)
		resp := rec.Result()

		body, _ := ioutil.ReadAll(resp.Body)
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
			Return(newRate, customerror.NewCustomError(errors.New(""), serverError.ServerInternalError, 1))

		bodys, _ := json.Marshal(testComment)
		req, err := http.NewRequest("POST", "/api/v1/comments", bytes.NewBuffer(bodys))

		assert.NoError(t, err)
		req = req.WithContext(context.WithValue(req.Context(), packageConfig.RequestUserID, int(testUser.UserID)))
		rec := httptest.NewRecorder()
		handler := CommentHandler{
			CommentUseCase: mockCUseCase,
			log:            logger.NewLogger(os.Stdout),
		}

		handler.AddComment(rec, req)
		resp := rec.Result()

		body, _ := ioutil.ReadAll(resp.Body)
		response := responses.HttpResponse{}

		err = json.Unmarshal(body, &response)
		assert.NoError(t, err)

		assert.Equal(t, serverError.ServerInternalError, response.Code)
	})

}

func TestCommentHandler_UpdateComment(t *testing.T) {

	testComment := commModel.Comment{
		UserID: 2, HotelID: 1, CommID: 2, Message: "kek", Rate: 2, Time: "20-10-2000",
	}

	testUser := userService.User{UserID: 2, Username: "kostik", Email: "sdfs@mail.ru", Avatar: "kek/img.jpeg"}
	newRate := commModel.NewRate{Comment: testComment, Rate: 3}
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

		req = req.WithContext(context.WithValue(req.Context(), packageConfig.RequestUserID, int(testUser.UserID)))
		rec := httptest.NewRecorder()
		handler := CommentHandler{
			CommentUseCase: mockCUseCase,
			log:            logger.NewLogger(os.Stdout),
		}

		handler.UpdateComment(rec, req)
		resp := rec.Result()
		comment := commModel.NewRate{}

		body, _ := ioutil.ReadAll(resp.Body)
		response := responses.HttpResponse{}

		err = json.Unmarshal(body, &response)
		assert.NoError(t, err)
		err = mapstructure.Decode(response.Data.(map[string]interface{}), &comment)
		assert.NoError(t, err)

		assert.Equal(t, newRate, comment)
		assert.Equal(t, http.StatusOK, response.Code)
	})

	t.Run("UpdateCommentErr1", func(t *testing.T) {
		testComments := []commModel.FullCommentInfo{
			{UserID: 2, CommID: 1, HotelID: 2, Message: "kek", Rating: 2, Avatar: "src/kek.jpg", Username: "kostik", Time: "20-10-2000"},
			{UserID: 2, CommID: 2, HotelID: 2, Message: "kekw", Rating: 2, Avatar: "src/kek.jpg", Username: "kostik", Time: "20-10-2000"},
		}

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockCUseCase := comment_mock.NewMockUsecase(ctrl)

		bodys, _ := json.Marshal(testComments)
		req, err := http.NewRequest("POST", "/api/v1/comments", bytes.NewBuffer(bodys))
		assert.NoError(t, err)

		req = req.WithContext(context.WithValue(req.Context(), packageConfig.RequestUserID, int(testUser.UserID)))
		rec := httptest.NewRecorder()
		handler := CommentHandler{
			CommentUseCase: mockCUseCase,
			log:            logger.NewLogger(os.Stdout),
		}

		handler.UpdateComment(rec, req)
		resp := rec.Result()

		body, _ := ioutil.ReadAll(resp.Body)
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
			log:            logger.NewLogger(os.Stdout),
		}

		handler.UpdateComment(rec, req)
		resp := rec.Result()

		body, _ := ioutil.ReadAll(resp.Body)
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
			Return(newRate, customerror.NewCustomError(errors.New(""), serverError.ServerInternalError, 1))

		bodys, _ := json.Marshal(testComment)
		req, err := http.NewRequest("POST", "/api/v1/comments", bytes.NewBuffer(bodys))

		assert.NoError(t, err)
		req = req.WithContext(context.WithValue(req.Context(), packageConfig.RequestUserID, int(testUser.UserID)))
		rec := httptest.NewRecorder()
		handler := CommentHandler{
			CommentUseCase: mockCUseCase,
			log:            logger.NewLogger(os.Stdout),
		}

		handler.UpdateComment(rec, req)
		resp := rec.Result()

		body, _ := ioutil.ReadAll(resp.Body)
		response := responses.HttpResponse{}

		err = json.Unmarshal(body, &response)
		assert.NoError(t, err)

		assert.Equal(t, serverError.ServerInternalError, response.Code)
	})

}

func TestCommentHandler_DeleteComment(t *testing.T) {
	testUser := userService.User{UserID: 2, Username: "kostik", Email: "sdfs@mail.ru", Avatar: "kek/img.jpeg"}

	t.Run("DeleteComment", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockCUseCase := comment_mock.NewMockUsecase(ctrl)

		mockCUseCase.EXPECT().
			DeleteComment(1).
			Return(nil)

		req, err := http.NewRequest("DELETE", "/api/v1/comments/1", nil)
		assert.NoError(t, err)

		req = req.WithContext(context.WithValue(req.Context(), packageConfig.RequestUserID, int(testUser.UserID)))
		req = mux.SetURLVars(req, map[string]string{
			"id": "1",
		})
		rec := httptest.NewRecorder()
		handler := CommentHandler{
			CommentUseCase: mockCUseCase,
			log:            logger.NewLogger(os.Stdout),
		}

		handler.DeleteComment(rec, req)
		resp := rec.Result()

		body, _ := ioutil.ReadAll(resp.Body)
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

		req = req.WithContext(context.WithValue(req.Context(), packageConfig.RequestUserID, int(testUser.UserID)))
		rec := httptest.NewRecorder()
		handler := CommentHandler{
			CommentUseCase: mockCUseCase,
			log:            logger.NewLogger(os.Stdout),
		}

		handler.DeleteComment(rec, req)
		resp := rec.Result()

		body, _ := ioutil.ReadAll(resp.Body)
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
			log:            logger.NewLogger(os.Stdout),
		}

		handler.DeleteComment(rec, req)
		resp := rec.Result()

		body, _ := ioutil.ReadAll(resp.Body)
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
			Return(customerror.NewCustomError(errors.New(""), serverError.ServerInternalError, 1))

		req, err := http.NewRequest("DELETE", "/api/v1/comments/1", nil)
		assert.NoError(t, err)

		req = req.WithContext(context.WithValue(req.Context(), packageConfig.RequestUserID, int(testUser.UserID)))
		req = mux.SetURLVars(req, map[string]string{
			"id": "1",
		})
		rec := httptest.NewRecorder()
		handler := CommentHandler{
			CommentUseCase: mockCUseCase,
			log:            logger.NewLogger(os.Stdout),
		}

		handler.DeleteComment(rec, req)
		resp := rec.Result()

		body, _ := ioutil.ReadAll(resp.Body)
		response := responses.HttpResponse{}

		err = json.Unmarshal(body, &response)
		assert.NoError(t, err)

		assert.Equal(t, serverError.ServerInternalError, response.Code)
	})

}
