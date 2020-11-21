package packageConfig

var PrefixPath string

const RequestUser = "User"

const SessionID = "SessionID"
const CorrectToken = "CorrectToken"
const Domen = "https://hostelscan.ru"
const LocalOrigin = "http://127.0.0.1"
const RequestUserID = "UserID"

var AllowedOrigins = map[string]bool{
	Domen:                true,
	LocalOrigin:          true,
	Domen + ":511":       true,
	Domen + ":72":        true,
	LocalOrigin + ":511": true,
	LocalOrigin + ":72":  true,
	LocalOrigin + ":443": true,
}
