package chatDelivery

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	chat_mock "github.com/go-park-mail-ru/2020_2_JMickhs/main/internal/app/chat/mocks"
	chat_model "github.com/go-park-mail-ru/2020_2_JMickhs/main/internal/app/chat/models"
	packageConfig "github.com/go-park-mail-ru/2020_2_JMickhs/package/configs"
	"github.com/go-park-mail-ru/2020_2_JMickhs/package/logger"
	"github.com/go-park-mail-ru/2020_2_JMickhs/package/responses"
	"github.com/golang/mock/gomock"
	"github.com/mitchellh/mapstructure"
	"github.com/stretchr/testify/assert"
)

func TestChatHandler_History(t *testing.T) {
	messagesTest := []chat_model.Message{
		{OwnerID: "2", Room: "3", Message: "wecsd", Moderator: false},
		{OwnerID: "2", Room: "3", Message: "hfgsdcx", Moderator: false},
	}

	t.Run("History", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockCUseCase := chat_mock.NewMockUsecase(ctrl)

		mockCUseCase.EXPECT().
			AddOrGetChat(gomock.Any(), 2).
			Return(messagesTest, nil)

		req, err := http.NewRequest("GET", "/api/v1/comments/?id=2&limit=1&offset=0", nil)
		assert.NoError(t, err)

		req = req.WithContext(context.WithValue(req.Context(), packageConfig.RequestUserID, 2))
		rec := httptest.NewRecorder()
		handler := ChatHandler{
			register:    make(chan subscription),
			unregister:  make(chan subscription),
			broadcast:   make(chan chat_model.Message),
			ChatUseCase: mockCUseCase,
			log:         logger.NewLogger(os.Stdout),
			bot:         nil,
			roomTokens:  map[string]string{},
			rooms:       map[string]map[*connection]bool{},
		}

		handler.History(rec, req)
		resp := rec.Result()
		var messages []chat_model.Message

		body, _ := ioutil.ReadAll(resp.Body)
		response := responses.HttpResponse{}

		err = json.Unmarshal(body, &response)
		assert.NoError(t, err)
		err = mapstructure.Decode(response.Data, &messages)
		assert.NoError(t, err)

		assert.Equal(t, messages, messagesTest)
		assert.Equal(t, http.StatusOK, response.Code)
	})

	t.Run("History", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockCUseCase := chat_mock.NewMockUsecase(ctrl)

		mockCUseCase.EXPECT().
			GetChatHistoryByID(gomock.Any()).
			Return(messagesTest, nil)

		req, err := http.NewRequest("GET", "/api/v1/comments/?chatID=2343&token=1", nil)
		assert.NoError(t, err)

		req = req.WithContext(context.WithValue(req.Context(), packageConfig.RequestUserID, 2))
		req = req.WithContext(context.WithValue(req.Context(), packageConfig.RequestUserRule, true))
		rec := httptest.NewRecorder()
		handler := ChatHandler{
			register:    make(chan subscription),
			unregister:  make(chan subscription),
			broadcast:   make(chan chat_model.Message),
			ChatUseCase: mockCUseCase,
			log:         logger.NewLogger(os.Stdout),
			bot:         nil,
			roomTokens:  map[string]string{},
			rooms:       map[string]map[*connection]bool{},
		}
		handler.roomTokens["2343"] = "1"

		handler.History(rec, req)
		resp := rec.Result()
		var messages []chat_model.Message

		body, _ := ioutil.ReadAll(resp.Body)
		response := responses.HttpResponse{}

		err = json.Unmarshal(body, &response)
		assert.NoError(t, err)
		err = mapstructure.Decode(response.Data, &messages)
		assert.NoError(t, err)

		assert.Equal(t, messages, messagesTest)
		assert.Equal(t, http.StatusOK, response.Code)
	})
}

func TestChatHandler_InitConnectionForModer(t *testing.T) {
	messagesTest := []chat_model.Message{
		{OwnerID: "2", Room: "3", Message: "wecsd", Moderator: false},
		{OwnerID: "2", Room: "3", Message: "hfgsdcx", Moderator: false},
	}

	t.Run("History", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockCUseCase := chat_mock.NewMockUsecase(ctrl)

		mockCUseCase.EXPECT().
			AddOrGetChat(gomock.Any(), 2).
			Return(messagesTest, nil)

		req, err := http.NewRequest("GET", "/api/v1/comments/?chatID=2343&token=1", nil)
		assert.NoError(t, err)

		req = req.WithContext(context.WithValue(req.Context(), packageConfig.RequestUserID, 2))
		req = req.WithContext(context.WithValue(req.Context(), packageConfig.RequestUserRule, true))
		rec := httptest.NewRecorder()
		handler := ChatHandler{
			register:    make(chan subscription),
			unregister:  make(chan subscription),
			broadcast:   make(chan chat_model.Message),
			ChatUseCase: mockCUseCase,
			log:         logger.NewLogger(os.Stdout),
			bot:         nil,
			roomTokens:  map[string]string{},
			rooms:       map[string]map[*connection]bool{},
		}

		handler.History(rec, req)
		resp := rec.Result()
		var messages []chat_model.Message

		body, _ := ioutil.ReadAll(resp.Body)
		response := responses.HttpResponse{}

		err = json.Unmarshal(body, &response)
		assert.NoError(t, err)
		err = mapstructure.Decode(response.Data, &messages)
		assert.NoError(t, err)

		assert.Equal(t, messages, messagesTest)
		assert.Equal(t, http.StatusOK, response.Code)
	})

}
