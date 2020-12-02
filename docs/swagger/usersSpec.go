package swagger

import (
	"mime/multipart"
)

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

type UpdateAvatar struct {
	Avatar multipart.File `json:"avatar"`
}

type Avatar struct {
	Avatar string `json:"avatar"`
}

type UpdateUser struct {
	Username string `json:"username" db:"username"`
}

type UpdateEmail struct {
	Email string `json:"email" db:"email"`
}

type UserAuth struct {
	Username string `json:"username" db:"username" validate:"required" `
	Password string `json:"password" db:"password" validate:"required"`
}

type UserRegistation struct {
	Username string `json:"username" db:"username" validate:"required" `
	Password string `json:"password" db:"password" validate:"required"`
	Email    string `json:"email" db:"email"`
}

// swagger:parameters AddSessions
type userAuthRequestWrapper struct {
	// in: body
	Body UserAuth
}

// swagger:parameters signup
type userRegistrationRequestWrapper struct {
	// in: body
	Body UserRegistation
}

// swagger:response signup
type userRegistrationResponseWrapper struct {
	// in: body
	Body SafeUser
}

// swagger:parameters password
type userUpPasswordRequestWrapper struct {
	// in: body
	Body UpdatePassword
}

// swagger:parameters credentials
type userUpUserRequestWrapper struct {
	// in: body
	Body UpdateUser
}

// swagger:parameters avatar
type userUpAvatarRequestWrapper struct {
	// avatar in *.jpg *.jpeg *.png format
	//	in: body
	Body UpdateAvatar
}

// swagger:response safeUser
type SafeUserResponse struct {
	//in:body
	Body SafeUser
}

// swagger:parameters userById
type UserIdParameter struct {
	//in:path
	ID int `json:"id"`
}

// swagger:response avatar
type AvatarUserResponse struct {
	//in:body
	Body Avatar
}

// swagger:parameters email
type EmailParameteter struct {
	//in:body
	Body UpdateEmail
}
