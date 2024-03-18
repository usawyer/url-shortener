package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"go.uber.org/zap"
)

type Config struct {
	Host       string `env-default:"localhost" env:"SERVER_HOST"`
	Port       string `env-default:"8080" env:"SERVER_PORT"`
	DbHost     string `env-default:"localhost" env:"DB_HOST"`
	DbPort     string `env-default:"5432" env:"DB_PORT"`
	DbName     string `env-default:"dbname" env:"POSTGRES_DB"`
	PgUser     string `env-default:"postgres" env:"POSTGRES_USER"`
	PgPassword string `env-default:"postgres" env:"POSTGRES_PASSWORD"`
	RdHost     string `env-default:"localhost" env:"REDIS_HOST"`
	RdPort     string `env-default:"6379" env:"REDIS_PORT"`
	RdPassword string `env-default:"redis" env:"REDIS_PASSWORD"`
}

func New(logger *zap.Logger) *Config {
	logger = logger.Named("Config")
	logger.Info("initializing config")
	var cfg Config
	err := cleanenv.ReadEnv(&cfg)
	//err := cleanenv.ReadConfig("config/.env", &cfg)
	if err != nil {
		logger.Fatal(err.Error())
	}

	return &cfg
}
