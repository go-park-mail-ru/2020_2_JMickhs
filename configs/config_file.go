package configs

import (
	"os"
	"time"
)

type postgresConfig struct {
	User     string
	Password string
	DBName   string
}

type redisConfig struct {
	Address  string
	Password string
	Bd       string
}

var BdConfig postgresConfig
var RedisConfig redisConfig
var PrefixPath string

const ApiVersion = "api/v1"
const Domen = "http://www.hostelscan.ru"
const LocalOrigin = "http://127.0.0.1"
const StaticPath = "static/img/"
const Port = ":8080"
const CookieLifeTime = time.Hour * 24 * 30
const BaseAvatarPath = "static/img/defaultAvatar.png"
const RequestUser = "User"
const DeliveryError = "Error"
const BasePageCount = 30
const BaseItemsPerPage = 5
const PreviewItemLimit = 6
const BucketName = "hostelscan"
const S3Url = "http://s3.hostelscan.ru/"

const (
	MB = 1 << 20
)

func Init() {

	BdConfig = postgresConfig{
		User:     os.Getenv("PostgresUser"),
		Password: os.Getenv("PostgresPassword"),
		DBName:   os.Getenv("PostgresDBName"),
	}

	RedisConfig = redisConfig{
		Address:  os.Getenv("RedisAddress"),
		Password: os.Getenv("RedisPassword"),
		Bd:       os.Getenv("RedisBd"),
	}
}
