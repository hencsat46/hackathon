package jwt

import (
	"hackathon/pkg/config"
	"log/slog"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWT struct {
	secret  string
	expTime time.Duration
}

func New(cfg *config.Config) *JWT {
	return &JWT{
		secret:  cfg.JWTsecret,
		expTime: time.Duration(cfg.ExpTime),
	}
}

func (j *JWT) CreateToken(guid string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"guid":      guid,
		"exp":       time.Now().Add(j.expTime).Unix(),
		"createdAt": time.Now().Unix(),
	})

	stringToken, err := token.SignedString([]byte(j.secret))
	if err != nil {
		slog.Error(err.Error())
	}

	return stringToken
}
