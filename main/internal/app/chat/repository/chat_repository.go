package chatRepository

import (
	"github.com/jmoiron/sqlx"
)

type ChatRepository struct {
	conn *sqlx.DB
}

func NewChatRepository(conn *sqlx.DB) ChatRepository {
	return ChatRepository{conn}
}
