package sessionsRepository

import (
	"context"
	"strconv"
	"time"

	"github.com/go-park-mail-ru/2020_2_JMickhs/sessions/configs"

	"github.com/spf13/viper"

	"github.com/go-park-mail-ru/2020_2_JMickhs/package/clientError"
	customerror "github.com/go-park-mail-ru/2020_2_JMickhs/package/error"

	"github.com/go-redis/redis/v8"
)

type sessionsRepository struct {
	sessStore *redis.Client
}

func NewSessionsUserRepository(sessStore *redis.Client) sessionsRepository {
	sessStorage := sessionsRepository{sessStore}
	return sessStorage
}

func (p *sessionsRepository) AddToken(token string, ID int64) (string, error) {
	err := p.sessStore.Set(context.Background(), token, ID, time.Duration(viper.GetInt(configs.ConfigFields.CookieLifeTime))*time.Minute).Err()

	if err != nil {
		return token, customerror.NewCustomError(err, clientError.BadRequest, 1)
	}
	return token, nil
}

func (p *sessionsRepository) GetIDByToken(token string) (int64, error) {
	response, err := p.sessStore.Get(context.Background(), token).Result()
	if err != nil {
		return 0, customerror.NewCustomError(err, clientError.BadRequest, 1)
	}
	res, err := strconv.Atoi(response)
	if err != nil {
		return 0, customerror.NewCustomError(err, clientError.BadRequest, 1)
	}
	return int64(res), nil
}

func (p *sessionsRepository) DeleteSession(token string) error {
	_, err := p.sessStore.Del(context.Background(), token).Result()
	if err != nil {
		return customerror.NewCustomError(err, clientError.BadRequest, 1)
	}
	return nil
}
