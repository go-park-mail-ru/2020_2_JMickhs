//go:generate easyjson -all chat_model.go
package chat_model

// easyjson:json
type Message struct {
	OwnerID   string
	Room      string
	Message   string
	Moderator bool
}
