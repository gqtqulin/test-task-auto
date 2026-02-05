package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	ServerPort string `envconfig:"SERVER_PORT"`
	DBHost     string `envconfig:"DB_HOST" required:"true"`
	DBPort     string `envconfig:"DB_PORT" required:"true"`
	DBName     string `envconfig:"DB_NAME" required:"true"`
	DBUser     string `envconfig:"DB_USER" required:"true"`
	DBPassword string `envconfig:"DB_PASSWORD" required:"true"`
}

func InitConfig() (*Config, error) {
	godotenv.Load()

	var config Config
	return &config, envconfig.Process("", &config)
}
