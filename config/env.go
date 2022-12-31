package config

import (
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type EnvConf struct {
	BaseUrl    string `env:"BASE_URL"`
	DbHost     string `env:"DB_HOST,notEmpty"`
	DbUser     string `env:"DB_USER,notEmpty"`
	DbPassword string `env:"DB_PASSWORD,notEmpty"`
	DbName     string `env:"DB_NAME,notEmpty"`
	DbPort     string `env:"DB_PORT,notEmpty"`
	JwtSecret  string `env:"JWT_SECRET,notEmpty"`
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
