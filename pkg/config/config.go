package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Environment string
}

func New() *Config {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}

	return &Config{
		Environment: os.Getenv("ENV"),
	}
}
