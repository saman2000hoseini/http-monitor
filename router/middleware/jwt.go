package middleware

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"strings"
)

func Skipper(c echo.Context) bool {
	if strings.Contains(c.Request().URL.Path, "/user/login") || strings.Contains(c.Request().URL.Path, "/user/register") {
		return true
	}
	return false
}

func JWT(key interface{}) middleware.JWTConfig {
	return middleware.JWTConfig{Skipper: Skipper, SigningKey: key}
}
