package configs

import (
	"os"
	"time"
)

type redisConfig struct {
	Address  string
	Password string
	Bd       string
}

var SecretTokenKey string

var RedisConfig redisConfig
var PrefixPath string

const CookieLifeTime = time.Hour * 24 * 30
const CsrfExpire = time.Minute * 15
const BucketName = "hostelscan"
const SessionID = "SessionID"
const SessionGrpcServicePort = ":8079"

func Init() {

	RedisConfig = redisConfig{
		Address:  os.Getenv("RedisAddress"),
		Password: os.Getenv("RedisPassword"),
		Bd:       os.Getenv("RedisBd"),
	}

	SecretTokenKey = os.Getenv("SecretTokenKey")
}
