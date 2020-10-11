package models

import "mime/multipart"

// swagger:response user
type User struct {
	ID       int    `json:"id" db:"id" validate:"required"`
	Username string `json:"username" db:"username" validate:"required" `
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password" validate:"required"`
	Avatar   string `json:"avatar" db:"avatar"`
}

// swagger:response safeUser
type SafeUser struct {
	ID       int    `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
	Email    string `json:"email" db:"email"`
	Avatar   string `json:"avatar" db:"avatar"`
}

type UserName struct {
	Username string `json:"username"`
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

type UpdateUser struct {
	Username string `json:"username" db:"username"`
	Email    string `json:"email" db:"email"`
}
type UpdatePassword struct{
	OldPassword string `json:"oldpassword" db:"password" validate:"required"`
	NewPassword string `json:"newpassword" db:"password" validate:"required"`
}

type UpdateAvatar struct{
	Avatar multipart.File `json:"avatar"`
}

// swagger:parameters Addsessions
type userAuthRequestWrapper struct {
	// in: body
	Body UserAuth
}

// swagger:parameters signup
type userRegistrationRequestWrapper struct {
	// in: body
	Body UserRegistation
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

//wrong old password
//swagger:response conflict
type conflict struct{
}

// swagger:parameters userById
type hotelIDParameterWrapper struct {
	// the id of user to get from database
	// in: path
	// required:true
	ID int `json:"id"`
}