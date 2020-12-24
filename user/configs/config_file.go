package configs

import (
	"os"

	"github.com/spf13/viper"
)

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
	UserHttpServicePort    string
	CertPath               string
	KeyPath                string
}{
	StaticPathForAvatars:   "paths.StaticPathForAvatars",
	CertPath:               "paths.CertPath",
	KeyPath:                "paths.KeyPath",
	CookieLifeTime:         "cookie.LifeTime",
	BaseAvatarPath:         "paths.BaseAvatarPath",
	BucketName:             "s3.BucketName",
	S3Url:                  "s3.S3Url",
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
