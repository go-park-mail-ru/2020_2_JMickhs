package csrfRepository

import (
	"context"
	"errors"

	customerror "github.com/go-park-mail-ru/2020_2_JMickhs/package/error"

	"github.com/go-park-mail-ru/2020_2_JMickhs/JMickhs_main/configs"

	"github.com/go-park-mail-ru/2020_2_JMickhs/package/serverError"
	"github.com/go-redis/redis/v8"
)

type csrfRepository struct {
	csrfStore *redis.Client
}

func NewCsrfRepository(sessStore *redis.Client) csrfRepository {
	return csrfRepository{sessStore}
}

func (r *csrfRepository) Add(token string) error {
	err := r.csrfStore.Set(context.Background(), token, 1, configs.CsrfExpire).Err()
	if err != nil {
		return customerror.NewCustomError(err, serverError.ServerInternalError, 1)
	}
	return nil
}

func (r *csrfRepository) Check(token string) error {
	_, err := r.csrfStore.Get(context.Background(), token).Result()
	if err != nil {
		if err == redis.Nil {
			return nil
		}
		return err
	}
	return errors.New("token not valid")
}
