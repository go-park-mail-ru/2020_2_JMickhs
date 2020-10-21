package models

// swagger:response user
type User struct {
	ID       int    `json:"id" db:"id"`
	Username string `json:"username" db:"username" validate:"required,min=3,max=15"`
	Email    string `json:"email" db:"email" validate:"required,email"`
	Password string `json:"password" db:"password" validate:"required,min=5,max=30"`
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

type UpdatePassword struct {
	OldPassword string `json:"oldpassword" db:"password" validate:"required"`
	NewPassword string `json:"newpassword" db:"password" validate:"required,min=5,max=30"`
}
