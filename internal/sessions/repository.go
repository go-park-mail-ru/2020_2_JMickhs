package sessions

type Repository interface {
	AddToken(token string, ID int) (string, error)
	GetIDByToken(token string) (int, error)
	DeleteSession(token string) error
}
