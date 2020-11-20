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
const (
	MB = 1 << 20
)

func Init() {
	BdConfig = postgresConfig{
		User:     os.Getenv("PostgresUser"),
		Password: os.Getenv("PostgresPassword"),
		DBName:   os.Getenv("UserPostgresDBName"),
	}
}
