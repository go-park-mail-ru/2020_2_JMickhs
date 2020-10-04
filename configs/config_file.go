package configs

import (
	"os"
	"time"
)

type postgresConfig struct {
	User string
	Password string
	DBName string
}

type redisConfig struct{
	Address string
	Password string
	Bd string
}
var BdConfig postgresConfig
var RedisConfig redisConfig

const ApiVersion = "api/v1"
const StaticPath = "static/img"
const Port  = ":8080"
const CookieLifeTime = time.Hour*24*30
const BaseAvatarPath = "static/img/defaultAvatar.jpg"

func Init() {

	BdConfig = postgresConfig{
		User:     os.Getenv("PostgresUser"),
		Password: os.Getenv("PostgresPassword"),
		DBName:   os.Getenv("PostgresDBName"),
	}

	RedisConfig = redisConfig{
		Address:   os.Getenv("RedisAddress"),
		Password:  os.Getenv("RedisPassword"),
		Bd: os.Getenv("RedisBd"),
	}
}