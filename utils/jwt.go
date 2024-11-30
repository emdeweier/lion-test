package utils

import (
	"github.com/golang-jwt/jwt/v4"
	"lion-test/helper"
	"time"
)

func GetSecretKey() []byte {
	secretKey := []byte(helper.GoDotEnvVariable("SECRET_KEY"))
	return secretKey
}

func GenerateToken(username, role string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"role":     role,
		"exp":      time.Now().Add(time.Hour * 1).Unix(),
	})
	return token.SignedString(GetSecretKey())
}
