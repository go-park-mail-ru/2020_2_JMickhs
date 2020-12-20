package chatRepository

const GetChatRequest = "SELECT chat_id from chats where user_id = $1"

const AddChatRequest = "INSERT INTO chats VALUES($1,$2)"
