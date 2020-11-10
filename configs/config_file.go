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

var SecretTokenKey string

var BdConfig postgresConfig
var RedisConfig redisConfig
var PrefixPath string

const Domen = "https://hostelscan.ru"
const LocalOrigin = "http://127.0.0.1"
const StaticPath = "static/img/"
const Port = ":8080"
const CookieLifeTime = time.Hour * 24 * 30
const CsrfExpire = time.Minute * 15
const BaseAvatarPath = "static/img/defaultAvatar.png"
const RequestUser = "User"
const BasePageCount = 30
const BaseItemsPerPage = 1
const PreviewItemLimit = 6
const BucketName = "hostelscan"
const S3Url = "https://hostelscan.hb.bizmrg.com/"
const BaseItemPerPage = 28
const SessionID = "SessionID"
const CorrectToken = "CorrectToken"
const S3Region = "ru-msk"
const S3EndPoint = "https://hb.bizmrg.com"
const (
	MB = 1 << 20
)
var AllowedOrigins  = map[string]bool{
	Domen : true,
	LocalOrigin:true,
	Domen+":511" : true,
	Domen+":72"  : true,
	LocalOrigin+":511" :true,
	LocalOrigin+":72" :true,
	LocalOrigin+":443":true,
}

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

	SecretTokenKey = os.Getenv("SecretTokenKey")
}
