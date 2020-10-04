package sessionsRepository

import (
	"context"
	"github.com/go-park-mail-ru/2020_2_JMickhs/configs"
	"github.com/go-redis/redis/v8"
	"strconv"
)

type sessionsRepository struct {
	sessStore *redis.Client
}

func NewSessionsUserRepository(sessStore *redis.Client) sessionsRepository {
	sessStorage := sessionsRepository{sessStore}
	return sessStorage
}

func (p *sessionsRepository) AddToken(token string, ID int) (string, error) {
	err := p.sessStore.Set(context.Background(), token, ID, configs.CookieLifeTime).Err()

	return token, err
}

func (p *sessionsRepository) GetIDByToken(token string) (int, error) {
	response, err := p.sessStore.Get(context.Background(), token).Result()
	if err != nil {
		return 0, err
	}
	res, err := strconv.Atoi(response)
	if err != nil {
		return 0, err
	}
	return res, nil
}

func (p *sessionsRepository) DeleteSession(token string) error {
	_, err := p.sessStore.Del(context.Background(), token).Result()
	return err
}
