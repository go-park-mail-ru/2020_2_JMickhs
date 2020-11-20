package configs

import (
	"os"
)

type postgresConfig struct {
	User     string
	Password string
	DBName   string
}

var BdConfig postgresConfig
var PrefixPath string

const Domen = "https://hostelscan.ru"
const LocalOrigin = "http://127.0.0.1"
const StaticPathForHotels = "static/img/"
const Port = ":8080"
const RequestUser = "User"
const BasePageCount = 30
const PreviewItemLimit = 6
const BucketName = "hostelscan"
const S3Url = "https://hostelscan.hb.bizmrg.com/"
const BaseItemPerPage = 28
const SessionID = "SessionID"
const S3Region = "ru-msk"
const S3EndPoint = "https://hb.bizmrg.com"
const RequestUserID = "UserID"

func Init() {

	BdConfig = postgresConfig{
		User:     os.Getenv("PostgresUser"),
		Password: os.Getenv("PostgresPassword"),
		DBName:   os.Getenv("PostgresDBName"),
	}

}
