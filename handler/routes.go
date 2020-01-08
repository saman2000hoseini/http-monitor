package handler

import "github.com/labstack/echo"

func (h *Handler) Setup(eg *echo.Group) {
	eg.POST("/register", h.Register)
	eg.POST("/login", h.Login)
	//	eg.GET("/user")
	//	eg.POST("/user")
	//	eg.PUT("/user")
}
