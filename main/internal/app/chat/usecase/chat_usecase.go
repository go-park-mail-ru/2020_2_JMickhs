package chatUsecase

import (
	"github.com/go-park-mail-ru/2020_2_JMickhs/main/internal/app/chat"
)

type ChatUseCase struct {
	chatRepo chat.Repository
}

func NewChatUseCase(r chat.Repository) *ChatUseCase {
	return &ChatUseCase{
		chatRepo: r,
	}
}

func (u *ChatUseCase) SendInviteToModeration(userID int) error {
	return nil
}

func (u *ChatUseCase) ServeChat(userID int, chatID int) error {
	return nil
}
