package delivery

import (
	"context"

	"github.com/go-park-mail-ru/2020_2_JMickhs/sessions/internal/session"

	"github.com/go-park-mail-ru/2020_2_JMickhs/sessions/internal/csrf"

	sessionService "github.com/go-park-mail-ru/2020_2_JMickhs/package/proto/sessions"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type SessionDelivery struct {
	SessionUseCase session.Usecase
	CsrfUseCase    csrf.Usecase
}

func NewSessionDelivery(SessionUseCase session.Usecase, CsrfUseCase csrf.Usecase) *SessionDelivery {
	return &SessionDelivery{SessionUseCase, CsrfUseCase}
}

func (sessDel *SessionDelivery) CreateSession(ctx context.Context, in *sessionService.UserID) (*sessionService.SessionID, error) {

	sessionID, err := sessDel.SessionUseCase.AddToken(in.UserID)
	if err != nil {
		return nil, status.Error(codes.Aborted, err.Error())
	}
	return &sessionService.SessionID{SessionID: sessionID}, nil
}

func (sessDel *SessionDelivery) GetIDBySession(ctx context.Context, in *sessionService.SessionID) (*sessionService.UserID, error) {
	userID, err := sessDel.SessionUseCase.GetIDByToken(in.SessionID)
	if err != nil {
		return nil, status.Error(codes.Aborted, err.Error())
	}
	return &sessionService.UserID{UserID: userID}, nil
}

func (sessDel *SessionDelivery) DeleteSession(ctx context.Context, in *sessionService.SessionID) (*sessionService.Empty, error) {
	err := sessDel.SessionUseCase.DeleteSession(in.SessionID)
	if err != nil {
		return nil, status.Error(codes.Aborted, err.Error())
	}
	return &sessionService.Empty{}, nil
}

func (sessDel *SessionDelivery) CreateCsrfToken(ctx context.Context, in *sessionService.CsrfTokenInput) (*sessionService.CsrfToken, error) {
	token, err := sessDel.CsrfUseCase.CreateToken(in.SessionID, in.TimeStamp)
	if err != nil {
		return nil, status.Error(codes.Aborted, err.Error())
	}
	return &sessionService.CsrfToken{Token: token}, nil
}
func (sessDel *SessionDelivery) CheckCsrfToken(ctx context.Context, in *sessionService.CsrfTokenCheck) (*sessionService.CheckResult, error) {
	res, err := sessDel.CsrfUseCase.CheckToken(in.SessionID, in.Token)
	if err != nil {
		return nil, status.Error(codes.Aborted, err.Error())
	}
	return &sessionService.CheckResult{Result: res}, nil
}
