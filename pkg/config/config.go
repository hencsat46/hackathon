package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Environment string
	JWTsecret   string
}

func New() *Config {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}

	return &Config{
		Environment: os.Getenv("ENV"),
		JWTsecret:   os.Getenv("JWT"),
	}
}
