//go:generate mockgen -source repository.go -destination mocks/sessions_repository_mock.go -package sessions_mock
package session

type Repository interface {
	AddToken(token string, ID int64) (string, error)
	GetIDByToken(token string) (int64, error)
	DeleteSession(token string) error
}
