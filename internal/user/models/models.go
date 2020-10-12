package models

// swagger:response user
type User struct {
	ID       int    `json:"id" db:"id" validate:"required"`
	Username string `json:"username" db:"username" validate:"required" `
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password" validate:"required"`
	Avatar   string `json:"avatar" db:"avatar"`
}

type SafeUser struct {
	ID       int    `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
	Email    string `json:"email" db:"email"`
	Avatar   string `json:"avatar" db:"avatar"`
}

type UserName struct {
	Username string `json:"username"`
}

type UpdatePassword struct{
	OldPassword string `json:"oldpassword" db:"password" validate:"required"`
	NewPassword string `json:"newpassword" db:"password" validate:"required"`
}


