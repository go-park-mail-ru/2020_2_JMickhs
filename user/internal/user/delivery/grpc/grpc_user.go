package userGrpcDelivery

import (
	"context"

	"github.com/go-park-mail-ru/2020_2_JMickhs/user/internal/user"

	userService "github.com/go-park-mail-ru/2020_2_JMickhs/package/proto/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserDelivery struct {
	UserUseCase user.Usecase
}

func NewUserDelivery(UserUseCase user.Usecase) *UserDelivery {
	return &UserDelivery{UserUseCase}
}

func (userDel *UserDelivery) GetUserByID(ctx context.Context, in *userService.UserID) (*userService.User, error) {
	user, err := userDel.UserUseCase.GetUserByID(int(in.UserID))
	if err != nil {
		return nil, status.Error(codes.Aborted, err.Error())
	}
	return &userService.User{UserID: int64(user.ID), Username: user.Username, Email: user.Email, Avatar: user.Avatar}, nil
}
