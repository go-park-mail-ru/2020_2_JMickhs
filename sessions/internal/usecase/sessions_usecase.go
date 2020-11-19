package sessionsUseCase

import (
	session "github.com/go-park-mail-ru/2020_2_JMickhs/JMickhs_sessions/internal"
	uuid "github.com/satori/go.uuid"
)

type sessionsUseCase struct {
	sessionsRepo session.Repository
}

func NewSessionsUsecase(r session.Repository) *sessionsUseCase {
	return &sessionsUseCase{
		sessionsRepo: r,
	}
}

func (u *sessionsUseCase) AddToken(ID int64) (string, error) {
	token := uuid.NewV4().String()
	return u.sessionsRepo.AddToken(token, ID)
}

func (u *sessionsUseCase) GetIDByToken(token string) (int64, error) {
	return u.sessionsRepo.GetIDByToken(token)
}

func (u *sessionsUseCase) DeleteSession(token string) error {
	return u.sessionsRepo.DeleteSession(token)
}
