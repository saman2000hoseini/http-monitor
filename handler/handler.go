package handler

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/saman2000hoseini/http-monitor/router/middleware"
	"github.com/saman2000hoseini/http-monitor/store"
)

type Handler struct {
	UserStore store.UserStore
	URLStore  store.URLStore
}

func NewHandler(db *gorm.DB) *Handler {
	return &Handler{store.NewUserStore(db), store.NewURLStore(db)}
}

func ReadToken(c echo.Context) uint {
	u := c.Get("user").(*jwt.Token)
	claims := u.Claims.(*middleware.JWTCustomClaims)
	return claims.ID
}
