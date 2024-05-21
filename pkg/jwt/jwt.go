package jwt

import (
	"hackathon/pkg/config"
	"time"
)

type JWT struct {
	secret  string
	expTime time.Duration
}

func New(cfg *config.Config) *JWT {
	return &JWT{
		secret:  cfg.JWTsecret,
		expTime: cfg.ExpTime,
	}
}
