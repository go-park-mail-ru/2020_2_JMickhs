package packageConfig

var PrefixPath string

type requestUser string
type sessionsId string
type correctToken string
type domen string
type localOrigin string
type requestUserID string
type requestUserRule string

const RequestUserID requestUserID = "UserID"
const RequestUserRule requestUserRule = "UserRule"
const RequestUser requestUser = "User"
const SessionID sessionsId = "SessionsID"
const CorrectToken correctToken = "CorrectToken"
const Domen domen = "https://hostelscan.ru"
const LocalOrigin localOrigin = "http://127.0.0.1"

var AllowedOrigins = map[string]bool{
	string(Domen):                true,
	string(LocalOrigin):          true,
	string(Domen + ":511"):       true,
	string(Domen + ":322"):       true,
	string(Domen + ":228"):       true,
	string(Domen + ":72"):        true,
	string(LocalOrigin + ":511"): true,
	string(LocalOrigin + ":72"):  true,
	string(LocalOrigin + ":443"): true,
}
