package jwt

import (
	"hackathon/pkg/config"
	e "hackathon/pkg/exceptions"
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
		"guid": guid,
		"exp":  time.Now().Add(j.expTime).Unix(),
	})

	stringToken, err := token.SignedString([]byte(j.secret))
	if err != nil {
		slog.Error(err.Error())
	}

	return stringToken
}

func (j *JWT) Validate(tokenString, guid string) bool {
	valid, uid := j.validateToken(tokenString)
	if !valid || uid != guid {
		return false
	}
	return true
}

func (j *JWT) validateToken(tokenString string) (bool, string) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, e.ErrInvalidSigningMethod
		}
		return []byte(j.secret), nil
	})

	// Check if it is valid.
	if err != nil || !token.Valid {
		slog.Error(err.Error())
		return false, ""
	}

	// Extract guid from claims.
	claims := token.Claims.(jwt.MapClaims)

	id := claims["guid"]
	guid, ok := id.(string)
	if !ok {
		return false, ""
	}

	return true, guid
}
