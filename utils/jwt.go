package utils

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

var JWTSecret = []byte("SECRET_TOKEN")

type JWTCustomClaims struct {
	ID uint `json:"id"`
	jwt.StandardClaims
}

func GenerateJWT(id uint) (string, error) {
	claims := &JWTCustomClaims{id, jwt.StandardClaims{ExpiresAt: time.Now().Add(time.Hour * 48).Unix()}}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString(JWTSecret)
	return t, err
}
