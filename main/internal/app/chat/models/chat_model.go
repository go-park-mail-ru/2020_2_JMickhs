//go:generate easyjson -all chat_model.go
package chat_model

// easyjson:json
type Message struct {
	OwnerID   string `mapstructure:"ownerID"`
	Room      string `mapstructure:"room"`
	Message   string `mapstructure:"message"`
	Moderator bool   `mapstructure:"moderator"`
}
