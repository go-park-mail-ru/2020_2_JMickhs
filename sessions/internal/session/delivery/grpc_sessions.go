package delivery

import (
	"context"

	"github.com/go-park-mail-ru/2020_2_JMickhs/JMickhs_sessions/internal/session"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type SessionDelivery struct {
	SessionUseCase session.Usecase
}

func NewSessionDelivery(useCase session.Usecase) *SessionDelivery {
	return &SessionDelivery{useCase}
}

func (sessDel *SessionDelivery) CreateSession(ctx context.Context, in *session_proto.UserID) (*session.SessionID, error) {
	sessionID, err := sessDel.SessionUseCase.AddToken(in.UserID)
	if err != nil {
		return nil, status.Error(codes.Aborted, err.Error())
	}
	return &session.SessionID{SessionID: sessionID}, nil
}

func (sessDel *SessionDelivery) GetIDBySession(ctx context.Context, in *session.SessionID) (*session.UserID, error) {
	userID, err := sessDel.SessionUseCase.GetIDByToken(in.SessionID)
	if err != nil {
		return nil, status.Error(codes.Aborted, err.Error())
	}
	return &session.UserID{UserID: userID}, nil
}

func (sessDel *SessionDelivery) DeleteSession(ctx context.Context, in *session.SessionID) (*session.Empty, error) {
	err := sessDel.SessionUseCase.DeleteSession(in.SessionID)
	if err != nil {
		return nil, status.Error(codes.Aborted, err.Error())
	}
	return &session.Empty{}, nil
}
