package swagger

import (
	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/user/models"
	"mime/multipart"
)

type UpdateAvatar struct{
	Avatar multipart.File `json:"avatar"`
}

type UpdateUser struct {
	Username string `json:"username" db:"username"`
	Email    string `json:"email" db:"email"`
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
	Body models.SafeUser
}

// swagger:parameters password
type userUpPasswordRequestWrapper struct {
	// in: body
	Body models.UpdatePassword
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


// swagger:response safeUser
type SafeUserResponse struct{
	//in:body
	Body models.SafeUser
}

// swagger:parameters userById
type UserIdParameter struct {
	//in:path
	ID int `json:"id"`
}