package handler

import (
	"errors"
	"fmt"
	"github.com/labstack/echo"
	"github.com/saman2000hoseini/http-monitor/model"
	"github.com/saman2000hoseini/http-monitor/utils"
	"net/http"
)

func (h *Handler) Register(c echo.Context) error {
	r := &userRegisterRequest{}
	user := &model.User{}
	err := r.bind(c, user)
	if err != nil {
		return c.String(http.StatusConflict, err.Error())
	}
	err = h.UserStore.Create(user)
	if err != nil {
		return err
	}
	var t string
	t, err = utils.GenerateJWT(user.ID)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, echo.Map{"token": t})
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
	var t string
	t, err = utils.GenerateJWT(user.ID)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{"token": t})
}

func (h *Handler) Update(c echo.Context) error {
	id := ReadToken(c)
	user := &model.User{}
	user, _ = h.UserStore.GetByID(id)
	user.Username = c.FormValue("username")
	err := user.HashPassword(c.FormValue("password"))
	if err != nil {
		return c.String(http.StatusConflict, err.Error())
	}
	fmt.Println(user.Username)
	fmt.Println(user.Password)
	err = h.UserStore.Update(user)
	if err != nil {
		return c.String(http.StatusConflict, err.Error())
	}
	return c.String(http.StatusOK, "Successfully updated")
}
