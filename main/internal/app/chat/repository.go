package chat

import chat_model "github.com/go-park-mail-ru/2020_2_JMickhs/main/internal/app/chat/models"

type Repository interface {
	AddMessageInChat(chatID string, message chat_model.Message) error
	AddOrGetChat(chatID string, userID int) ([]chat_model.Message, error)
	GetChatID(userID int) (string, error)
}
