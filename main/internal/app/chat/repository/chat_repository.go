package chatRepository

import (
	"context"

	"github.com/mailru/easyjson"

	chat_model "github.com/go-park-mail-ru/2020_2_JMickhs/main/internal/app/chat/models"
	customerror "github.com/go-park-mail-ru/2020_2_JMickhs/package/error"
	"github.com/go-park-mail-ru/2020_2_JMickhs/package/serverError"
	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
)

type ChatRepository struct {
	ChatStore *redis.Client
	Conn      *sqlx.DB
}

func NewChatRepository(conn *sqlx.DB, ChatStore *redis.Client) ChatRepository {
	return ChatRepository{Conn: conn, ChatStore: ChatStore}
}

func (u *ChatRepository) AddMessageInChat(chatID string, message chat_model.Message) error {
	msg, _ := message.MarshalJSON()
	err := u.ChatStore.Do(context.Background(), "LPUSH", chatID, msg).Err()
	if err != nil {
		return customerror.NewCustomError(err, serverError.ServerInternalError, 1)
	}
	return nil
}
func (u *ChatRepository) GetChatID(userID int) (string, error) {
	chatID := ""
	err := u.Conn.QueryRow(GetChatRequest, userID).Scan(&chatID)
	if err != nil {
		return chatID, customerror.NewCustomError(err, serverError.ServerInternalError, 1)
	}
	return chatID, nil
}

func (u *ChatRepository) AddOrGetChat(chatID string, userID int) ([]chat_model.Message, error) {
	var messages []chat_model.Message
	err := u.Conn.QueryRow(GetChatRequest, userID).Scan(&chatID)
	if err != nil {
		_, err = u.Conn.Exec(AddChatRequest, chatID, userID)
		if err != nil {
			return messages, customerror.NewCustomError(err, serverError.ServerInternalError, 1)
		}
		return messages, nil
	}
	count, err := u.ChatStore.Do(context.Background(), "LLEN", chatID).Int()
	if err != nil {
		return messages, customerror.NewCustomError(err, serverError.ServerInternalError, 1)
	}
	patternsInterface, err := u.ChatStore.Do(context.Background(), "LRANGE", chatID, 0, count).Result()
	if err != nil {
		return messages, customerror.NewCustomError(err, serverError.ServerInternalError, 1)
	}
	listOfInterfaces := patternsInterface.([]interface{})
	count = len(listOfInterfaces)
	if count == 0 {
		return messages, nil
	}
	messages = make([]chat_model.Message, count)

	for num := range listOfInterfaces {
		var msg chat_model.Message
		err := easyjson.Unmarshal([]byte(listOfInterfaces[num].(string)), &msg)
		if err != nil {
			return messages, customerror.NewCustomError(err, serverError.ServerInternalError, 1)
		}
		messages[num] = msg
	}
	return messages, nil
}
