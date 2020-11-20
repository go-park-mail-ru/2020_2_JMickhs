//go:generate mockgen -source usecase.go -destination mocks/sessions_usecase_mock.go -package sessions_mock
package session

type Usecase interface {
	AddToken(ID int64) (string, error)
	GetIDByToken(token string) (int64, error)
	DeleteSession(token string) error
}
