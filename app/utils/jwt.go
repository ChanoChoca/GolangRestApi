package utils

import (
	"flashpage/app/config"
	"time"

	"github.com/golang-jwt/jwt"
)

var JwtSecret = []byte("super_secret")

func GenerateJWT(userID int) (string, error) {
	expiration := time.Now().Add(time.Duration(config.JwtExpirationHours) * time.Hour).Unix()
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     expiration,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JwtSecret)
}

func ParseJWT(tokenStr string) (*jwt.Token, error) {
	return jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return JwtSecret, nil
	})
}