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

const StaticPathForHotels = "static/img/"
const Port = ":8080"
const BasePageCount = 30
const PreviewItemLimit = 6
const BucketName = "hostelscan"
const S3Url = "https://hostelscan.hb.bizmrg.com/"
const BaseItemPerPage = 28
const SessionID = "SessionID"
const S3Region = "ru-msk"
const S3EndPoint = "https://hb.bizmrg.com"
const RequestUserID = "UserID"
const SessionGrpcServicePort = ":8079"
const UserGrpcServicePort = ":8081"
const MainHttpServicePort = ":8080"

func Init() {

	BdConfig = postgresConfig{
		User:     os.Getenv("PostgresUser"),
		Password: os.Getenv("PostgresPassword"),
		DBName:   os.Getenv("PostgresDBName"),
	}

}
