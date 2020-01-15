package handler

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	middleware2 "github.com/saman2000hoseini/http-monitor/router/middleware"
)

func (h *Handler) Setup(eg *echo.Group) {
	config := middleware2.JWT(middleware2.JWTSecret)
	user := eg.Group("/user", middleware.JWTWithConfig(config))
	user.POST("/register", h.Register)
	user.POST("/login", h.Login)
	user.PUT("/update", h.Update)
	url := user.Group("/url", middleware.JWTWithConfig(config))
	url.PUT("", h.UpdateURL)
	url.POST("", h.AddURL)
	url.GET("/all", h.GetURLs)
	url.GET("", h.GetURL)
}
