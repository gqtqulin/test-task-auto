package config

import (
	"errors"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	ServerPort string
	DBHost     string
	DBPort     string
	DBName     string
	DBUser     string
	DBPassword string
}

// TODO: можно viper и сразу сделать схему в структуре
func InitConfig() (*Config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}

	srvPort := os.Getenv("SERVER_PORT")
	if srvPort == "" {
		srvPort = "8080"
	}

	DBHost := os.Getenv("DB_HOST")
	if DBHost == "" {
		return nil, errors.New("DB_HOST environment variable not set")
	}

	DBPort := os.Getenv("DB_PORT")
	if DBPort == "" {
		return nil, errors.New("DB_PORT environment variable not set")
	}

	DBName := os.Getenv("DB_NAME")
	if DBName == "" {
		return nil, errors.New("DB_NAME environment variable not set")
	}

	DBUser := os.Getenv("DB_USER")
	if DBUser == "" {
		return nil, errors.New("DB_USER environment variable not set")
	}

	DBPassword := os.Getenv("DB_PASSWORD")
	if DBPassword == "" {
		return nil, errors.New("DB_PASSWORD environment variable not set")
	}

	return &Config{
		ServerPort: srvPort,
		DBHost:     DBHost,
		DBPort:     DBPort,
		DBName:     DBName,
		DBUser:     DBUser,
		DBPassword: DBPassword,
	}, nil
}
