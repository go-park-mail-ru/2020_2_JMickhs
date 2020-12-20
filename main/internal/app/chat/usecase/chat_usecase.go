package chatUsecase

import (
	"github.com/go-park-mail-ru/2020_2_JMickhs/main/internal/app/chat"
	chat_model "github.com/go-park-mail-ru/2020_2_JMickhs/main/internal/app/chat/models"
)

type ChatUseCase struct {
	chatRepo chat.Repository
}

func NewChatUseCase(r chat.Repository) *ChatUseCase {
	return &ChatUseCase{
		chatRepo: r,
	}
}

func (u *ChatUseCase) AddMessageInChat(chatID string, message chat_model.Message) error {
	return u.chatRepo.AddMessageInChat(chatID, message)
}

func (u *ChatUseCase) AddOrGetChat(chatID string, userID int) ([]chat_model.Message, error) {
	return u.chatRepo.AddOrGetChat(chatID, userID)
}

func (u *ChatUseCase) GetChatID(userID int) (string, error) {
	return u.chatRepo.GetChatID(userID)
}
