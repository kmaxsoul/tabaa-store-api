package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort  string
	DatabaseURL string
}

func LoadConfig() (*Config, error) {
	var err error = godotenv.Load()

	if err != nil {
		log.Println(".env file is not found using environment variables instead!")
	}

	var config *Config = &Config{
		DatabaseURL: os.Getenv("DB_URL"),
		ServerPort:  os.Getenv("PORT"),
	}

	return config, nil
}
