package config

import (
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type EnvConf struct {
	BaseUrl string `env:"BASE_URL"`
}

func LoadEnv() (*EnvConf, error) {
	_ = godotenv.Load(".env")
	var envConf EnvConf
	err := env.Parse(&envConf)
	if err != nil {
		return nil, err
	}

	return &envConf, nil
}
