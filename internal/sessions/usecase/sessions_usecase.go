package sessionsUseCase

import (
	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/sessions"
	uuid "github.com/satori/go.uuid"
)

type sessionsUseCase struct {
	sessionsRepo sessions.Repository
}

func NewSessionsUsecase(r sessions.Repository) *sessionsUseCase {
	return &sessionsUseCase{
		sessionsRepo: r,
	}
}

func (u *sessionsUseCase) AddToken(ID int) (string, error) {
	token := uuid.NewV4().String()
	return u.sessionsRepo.AddToken(token, ID)
}

func (u *sessionsUseCase) GetIDByToken(token string) (int, error) {
	return u.sessionsRepo.GetIDByToken(token)
}

func (u *sessionsUseCase) DeleteSession(token string) error {
	return u.sessionsRepo.DeleteSession(token)
}
