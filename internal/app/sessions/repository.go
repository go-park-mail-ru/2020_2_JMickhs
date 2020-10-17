//go:generate mockgen -source repository.go -destination mocks/sessions_repository_mock.go -package mocks
package sessions

type Repository interface {
	AddToken(token string, ID int) (string, error)
	GetIDByToken(token string) (int, error)
	DeleteSession(token string) error
}
