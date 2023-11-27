package config

import (
	"log"

	env "github.com/Netflix/go-env"
)

var cfg Config

type Config struct {
	AppConfig
}

type AppConfig struct {
	Port int `env:"APP_PORT,required=true"`
}

func init() {

	_, err := env.UnmarshalFromEnviron(&cfg)
	if err != nil {
		log.Println("=========================", err)
	}
}

func GetConfig() Config {
	return cfg
}
