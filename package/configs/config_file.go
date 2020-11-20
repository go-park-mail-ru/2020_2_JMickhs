package packageConfig

var PrefixPath string

const RequestUser = "User"

const SessionID = "SessionID"
const CorrectToken = "CorrectToken"

var AllowedOrigins = map[string]bool{
	Domen:                true,
	LocalOrigin:          true,
	Domen + ":511":       true,
	Domen + ":72":        true,
	LocalOrigin + ":511": true,
	LocalOrigin + ":72":  true,
	LocalOrigin + ":443": true,
}
