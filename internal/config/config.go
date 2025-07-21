package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	DBDriver string
	DBDSN    string
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	return &Config{
		DBDriver: os.Getenv("DB_DRIVER"),
		DBDSN:    os.Getenv("DB_DSN"),
	}
}
