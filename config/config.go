package config

import (
	"log"

	env "github.com/Netflix/go-env"
)

var cfg Config

type Config struct {
	AppConfig      AppConfig
	PostgresConfig PostgresConfig
	AuthConfig     AuthConfig
	RedisConfig    RedisConfig
}

type AppConfig struct {
	Port         string `env:"APP_PORT,required=true"`
	ReadTimeOut  uint32 `env:"APP_READTIMEOUT,required=true"`
	WriteTimeOut uint32 `env:"APP_WRIETIMEOUT,required=true"`
	IdleTimeout  uint32 `env:"APP_IDLETIMEOUT,required=true"`
}

type PostgresConfig struct {
	Host     string `env:"POSTGRES_HOST,required=true"`
	User     string `env:"POSTGRES_USER,required=true"`
	Password string `env:"POSTGRES_PASSWORD,required=true"`
	Db       string `env:"POSTGRES_DB,required=true"`
	Port     string `env:"POSTGRES_PORT,required=true"`
	SslMode  string `env:"POSTGRES_SSLMODE,required=true"`
	TimeZone string `env:"POSTGRES_TIMEZONE,required=true"`
}

type AuthConfig struct {
	PublicKey  string `env:"PUBLIC_KEY,required=true"`
	PrivateKey string `env:"PRIVATE_KEY,required=true"`
}

type RedisConfig struct {
	Address  string `env:"ADDRESS,required=true"`
	Password string `env:"PASSWORD,required=true"`
	Db       string `env:"DB,required=true"`
}

func init() {

	_, err := env.UnmarshalFromEnviron(&cfg)
	if err != nil {
		log.Panic(err)
	}
}

func GetConfig() Config {
	return cfg
}
