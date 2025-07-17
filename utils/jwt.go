package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/pablon/backend-pablon/config"
)

func GenerateJWT(email string) (string, error) {
	cfg := config.GetConfig()
	claims := jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(cfg.JWTSecret))
}
