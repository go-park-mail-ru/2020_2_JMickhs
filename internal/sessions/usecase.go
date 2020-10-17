//go:generate mockgen -source usecase.go -destination mocks/sessions_usecase_mock.go -package mocks
package sessions

type Usecase interface {
	AddToken(ID int) (string, error)
	GetIDByToken(token string) (int, error)
	DeleteSession(token string) error
}