//go:generate mockgen -source repository.go -destination mocks/chat_repository_mock.go -package chat_mock
package chat

import chat_model "github.com/go-park-mail-ru/2020_2_JMickhs/main/internal/app/chat/models"

type Repository interface {
	AddMessageInChat(chatID string, message chat_model.Message) error
	AddOrGetChat(chatID string, userID int) ([]chat_model.Message, error)
	GetChatID(userID int) (string, error)
	GetChatHistoryByID(chatID string) ([]chat_model.Message, error)
}
