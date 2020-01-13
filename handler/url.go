package handler

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/saman2000hoseini/http-monitor/model"
	"github.com/saman2000hoseini/http-monitor/utils"
	"net/http"
	"strconv"
)

func (h *Handler) AddURL(c echo.Context) error {
	u := c.Get("user").(*jwt.Token)
	claims := u.Claims.(*utils.JWTCustomClaims)
	id := claims.ID
	user := &model.User{}
	user, _ = h.UserStore.GetByID(id)
	if len(user.URLs) >= model.URLLIMIT {
		return c.String(http.StatusForbidden, "You have reached the maximum limitation of urls")
	}
	Address := c.FormValue("url")
	u64, err := strconv.ParseUint(c.FormValue("threshold"), 10, 32)
	Threshold := uint(u64)
	url := model.NewURL(Address, Threshold, id)
	user.URLs = append(user.URLs, *url)
	err = h.UserStore.Update(user)
	if err != nil {
		return c.String(http.StatusConflict, err.Error())
	}
	return c.String(http.StatusOK, "url successfully added")
}

func (h *Handler) UpdateURL(c echo.Context) error {
	u := c.Get("user").(*jwt.Token)
	claims := u.Claims.(*utils.JWTCustomClaims)
	id := claims.ID
	url := &model.URL{}
	uid, _ := strconv.ParseUint(c.FormValue("id"), 10, 32)
	url, _ = h.URLStore.GetByID(uint(uid))
	if id != url.UserID {
		return c.String(http.StatusUnauthorized, "Access to selected url is denied")
	}
	url.Address = c.FormValue("url")
	u64, err := strconv.ParseUint(c.FormValue("threshold"), 10, 32)
	url.Threshold = uint(u64)
	err = h.URLStore.Update(url)
	if err != nil {
		return c.String(http.StatusConflict, err.Error())
	}
	return c.String(http.StatusOK, "Successfully updated")
}
