package sessions

type Usecase interface {
	AddToken(ID int) (string, error)
	GetIDByToken(token string) (int, error)
	DeleteSession(token string) error
}
