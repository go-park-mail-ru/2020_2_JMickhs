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

type redisConfig struct {
	Address  string
	Password string
	Bd       string
}

var ConfigFields = struct {
	StaticPathForAvatars     string
	CookieLifeTime           string
	BaseAvatarPath           string
	BucketName               string
	S3Url                    string
	S3Region                 string
	S3EndPoint               string
	SessionGrpcServicePort   string
	UserGrpcServicePort      string
	MainHttpServicePort      string
	StaticPathForHotels      string
	BasePageCount            string
	PreviewItemLimit         string
	BaseItemPerPage          string
	WishListIn               string
	WishListOut              string
	RecommendationCount      string
	UpdateRecommendationTick string
	PongWait                 string
	MaxMessageSize           string
	ChatID                   string
	StaticPathForComments    string
	CertPath                 string
	KeyPath                  string
	ChatTTL                  string
}{
	ChatTTL:                  "constants.ChatTTL",
	StaticPathForAvatars:     "paths.StaticPathForAvatars",
	StaticPathForComments:    "paths.StaticPathForComments",
	CertPath:                 "paths.CertPath",
	KeyPath:                  "paths.KeyPath",
	CookieLifeTime:           "cookie.LifeTime",
	BaseAvatarPath:           "paths.BaseAvatarPath",
	BucketName:               "s3.BucketName",
	S3Url:                    "s3.S3Url",
	S3Region:                 "s3.S3Region",
	S3EndPoint:               "s3.S3EndPoint",
	SessionGrpcServicePort:   "grpc.SessionGrpcServicePort",
	UserGrpcServicePort:      "grpc.UserGrpcServicePort",
	MainHttpServicePort:      "http.MainHttpServicePort",
	StaticPathForHotels:      "paths.StaticPathForHotels",
	BaseItemPerPage:          "constants.BaseItemPerPage",
	PreviewItemLimit:         "constants.PreviewItemLimit",
	BasePageCount:            "constants.BasePageCount",
	WishListIn:               "constants.WishListIn",
	WishListOut:              "constants.WishListOut",
	RecommendationCount:      "constants.RecommendationCount",
	UpdateRecommendationTick: "constants.UpdateRecommendationTick",
	PongWait:                 "constants.PongWait",
	MaxMessageSize:           "constants.MaxMessageSize",
	ChatID:                   "constants.ChatID",
}

var BdConfig postgresConfig
var PrefixPath string
var RedisConfig redisConfig

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

func ExportConfig() error {
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.SetConfigName("config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return nil
}
