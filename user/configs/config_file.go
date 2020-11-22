package configs

import (
	"os"
	"time"
)

type postgresConfig struct {
	User     string
	Password string
	DBName   string
	Port     string
}

var BdConfig postgresConfig
var PrefixPath string

const StaticPathForAvatars = "static/avatars/"
const CookieLifeTime = time.Hour * 24 * 30
const BaseAvatarPath = "static/avatars/defaultAvatar.png"
const RequestUser = "User"
const BucketName = "hostelscan"
const S3Url = "https://hostelscan.hb.bizmrg.com/"
const SessionID = "SessionID"
const S3Region = "ru-msk"
const S3EndPoint = "https://hb.bizmrg.com"
const SessionGrpcServicePort = ":8079"
const UserGrpcServicePort = ":8081"
const UserHttpServicePort = ":8082"
const (
	MB = 1 << 20
)

func Init() {
	BdConfig = postgresConfig{
		User:     os.Getenv("PostgresUser"),
		Password: os.Getenv("PostgresPassword"),
		DBName:   os.Getenv("UserPostgresDBName"),
		Port:     os.Getenv("UserPostgresHost"),
	}
}
