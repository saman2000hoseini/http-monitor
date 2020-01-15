package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"strings"
)

var JWTSecret = []byte("SECRET_TOKEN")

type JWTCustomClaims struct {
	ID uint `json:"id"`
	jwt.StandardClaims
}

//create func to skip authorization for public endpoints
func Skipper(c echo.Context) bool {
	if strings.Contains(c.Request().URL.Path, "/user/login") || strings.Contains(c.Request().URL.Path, "/user/register") {
		return true
	}
	return false
}

func JWT(key interface{}) middleware.JWTConfig {
	return middleware.JWTConfig{Skipper: Skipper, SigningKey: key, Claims: &JWTCustomClaims{}}
}
