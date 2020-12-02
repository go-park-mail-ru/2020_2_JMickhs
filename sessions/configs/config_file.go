package configs

import (
	"os"

	"github.com/spf13/viper"
)

var ConfigFields = struct {
	CookieLifeTime         string
	SessionID              string
	CsrfExpire             string
	SessionGrpcServicePort string
}{
	CookieLifeTime:         "cookie.LifeTime",
	SessionID:              "context.SessionID",
	CsrfExpire:             "csrf.CsrfExpire",
	SessionGrpcServicePort: "grpc.SessionGrpcServicePort",
}

type redisConfig struct {
	Address  string
	Password string
	Bd       string
}

var SecretTokenKey string
var RedisConfig redisConfig
var PrefixPath string

func Init() {

	RedisConfig = redisConfig{
		Address:  os.Getenv("RedisAddress"),
		Password: os.Getenv("RedisPassword"),
		Bd:       os.Getenv("RedisBd"),
	}

	SecretTokenKey = os.Getenv("SecretTokenKey")
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
