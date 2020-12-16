package chat

type Usecase interface {
	SendInviteToModeration(userID int) error
	ServeChat(userID int, chatID int) error
}
