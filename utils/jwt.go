package utils

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/saman2000hoseini/http-monitor/router/middleware"
	"time"
)

//generate jason web token with expire date and user_id
func GenerateJWT(id uint) (string, error) {
	claims := &middleware.JWTCustomClaims{ID: id, StandardClaims: jwt.StandardClaims{ExpiresAt: time.Now().Add(time.Hour * 72).Unix()}}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString(middleware.JWTSecret)
	return t, err
}
