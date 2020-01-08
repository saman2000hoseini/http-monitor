package handler

import (
	"errors"
	"fmt"
	"github.com/labstack/echo"
	"github.com/saman2000hoseini/http-monitor/model"
	"net/http"
)

func (h *Handler) Register(c echo.Context) error {
	r := &userRegisterRequest{}
	user := &model.User{}
	err := r.bind(c, user)
	if err != nil {
		return err
	}
	err = h.UserStore.Create(user)
	if err != nil {
		return err
	}
	err = c.String(201, "User Created")
	return err
}

func (h *Handler) Login(c echo.Context) error {
	r := &userLoginRequest{}
	err := r.bind(c)
	if err != nil {
		return c.String(http.StatusUnprocessableEntity, err.Error())
	}
	user := &model.User{}
	user, err = h.UserStore.GetByUsername(r.User.Username)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	if user == nil {
		return c.JSON(http.StatusForbidden, errors.New("AccessForbidden"))
	}
	if !user.CheckPassword(r.User.Password) {
		return c.JSON(http.StatusForbidden, errors.New("AccessForbidden"))
	}
	return c.JSON(http.StatusOK, "login successful")
}
