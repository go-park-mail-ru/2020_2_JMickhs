package configs

import (
	"os"

	"github.com/spf13/viper"
)

type postgresConfig struct {
	User     string
	Password string
	DBName   string
}

var ConfigFields = struct {
	StaticPathForAvatars   string
	CookieLifeTime         string
	BaseAvatarPath         string
	BucketName             string
	S3Url                  string
	S3Region               string
	S3EndPoint             string
	SessionGrpcServicePort string
	UserGrpcServicePort    string
	MainHttpServicePort    string
	StaticPathForHotels    string
	BasePageCount          string
	PreviewItemLimit       string
	BaseItemPerPage        string
	WishListIn             string
	WishListOut            string
}{
	StaticPathForAvatars:   "paths.StaticPathForAvatars",
	CookieLifeTime:         "cookie.LifeTime",
	BaseAvatarPath:         "paths.BaseAvatarPath",
	BucketName:             "s3.BucketName",
	S3Url:                  "s3.S3Url",
	S3Region:               "s3.S3Region",
	S3EndPoint:             "s3.S3EndPoint",
	SessionGrpcServicePort: "grpc.SessionGrpcServicePort",
	UserGrpcServicePort:    "grpc.UserGrpcServicePort",
	MainHttpServicePort:    "http.MainHttpServicePort",
	StaticPathForHotels:    "paths.StaticPathForHotels",
	BaseItemPerPage:        "constants.BaseItemPerPage",
	PreviewItemLimit:       "constants.PreviewItemLimit",
	BasePageCount:          "constants.BasePageCount",
	WishListIn:             "constants.WishListIn",
	WishListOut:            "constants.WishListOut",
}

type sessionID string
type requestUserID string

const SessionID sessionID = "SessionsID"
const RequestUserID requestUserID = "UserID"

var BdConfig postgresConfig
var PrefixPath string

const (
	MB = 1 << 20
)

func Init() {
	BdConfig = postgresConfig{
		User:     os.Getenv("PostgresUser"),
		Password: os.Getenv("PostgresPassword"),
		DBName:   os.Getenv("PostgresDBName"),
	}
}

func ExportConfig() error {
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.SetConfigName("config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return nil
}
