package configs

import (
	"os"

	"github.com/spf13/viper"
)

var ConfigFields = struct {
	StaticPathForAvatars   string
	CookieLifeTime         string
	BaseAvatarPath         string
	RequestUser            string
	BucketName             string
	S3Url                  string
	SessionID              string
	S3Region               string
	S3EndPoint             string
	SessionGrpcServicePort string
	UserGrpcServicePort    string
	UserHttpServicePort    string
}{
	StaticPathForAvatars:   "paths.StaticPathForAvatars",
	CookieLifeTime:         "cookie.LifeTime",
	BaseAvatarPath:         "paths.BaseAvatarPath",
	RequestUser:            "context.RequestUser",
	BucketName:             "s3.BucketName",
	S3Url:                  "s3.S3Url",
	SessionID:              "context.SessionID",
	S3Region:               "s3.S3Region",
	S3EndPoint:             "s3.S3EndPoint",
	SessionGrpcServicePort: "grpc.SessionGrpcServicePort",
	UserGrpcServicePort:    "grpc.UserGrpcServicePort",
	UserHttpServicePort:    "http.UserHttpServicePort",
}

type postgresConfig struct {
	User     string
	Password string
	DBName   string
	Port     string
}

var BdConfig postgresConfig
var PrefixPath string

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

func ExportConfig() error {
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.SetConfigName("config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return nil
}
