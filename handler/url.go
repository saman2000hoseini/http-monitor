package handler

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/saman2000hoseini/http-monitor/model"
	"net/http"
	"strconv"
)

func (h *Handler) AddURL(c echo.Context) error {
	id := ReadToken(c)
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
	id := ReadToken(c)
	url := &model.URL{}
	uid, _ := strconv.ParseUint(c.FormValue("id"), 10, 32)
	url, err := h.URLStore.GetByID(uint(uid))
	if gorm.IsRecordNotFoundError(err) {
		return err
	}
	if id != url.UserID {
		return c.String(http.StatusUnauthorized, "Access to selected url is denied")
	}
	url.Address = c.FormValue("url")
	u64, err := strconv.ParseUint(c.FormValue("threshold"), 10, 32)
	url.Threshold = uint(u64)
	url.Alert = nil
	url.FailedCall = 0
	url.SuccessCall = 0
	err = h.URLStore.Update(url)
	if err != nil {
		return c.String(http.StatusConflict, err.Error())
	}
	return c.String(http.StatusOK, "Successfully updated")
}

func (h *Handler) GetURLs(c echo.Context) error {
	id := ReadToken(c)
	user := &model.User{}
	user, _ = h.UserStore.GetByID(id)
	jsonUser, err := json.Marshal(user)
	if err != nil {
		return c.String(http.StatusConflict, err.Error())
	}
	return c.String(http.StatusOK, string(jsonUser))
}

func (h *Handler) GetURL(c echo.Context) error {
	id := ReadToken(c)
	uid, _ := strconv.ParseUint(c.FormValue("id"), 10, 32)
	url, err := h.URLStore.GetByID(uint(uid))
	if err != nil {
		return c.String(http.StatusConflict, err.Error())
	}
	if id != url.UserID {
		return c.String(http.StatusUnauthorized, "Access to selected url is denied")
	}
	jsonURL, err := json.Marshal(url)
	if err != nil {
		return c.String(http.StatusConflict, err.Error())
	}
	strJson := string(jsonURL)
	url.Alert, _ = h.URLStore.GetAlert(url.ID)
	if url.Alert != nil {
		jsonAlert, _ := json.Marshal(url.Alert)
		strJson += string(jsonAlert)
		h.URLStore.PublishAlert(url)
	}
	return c.String(http.StatusOK, strJson)
}

func (h *Handler) GetURLStatistics(c echo.Context) error {
	id := ReadToken(c)
	uid, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	url, err := h.URLStore.GetByID(uint(uid))
	if err != nil {
		return c.String(http.StatusConflict, err.Error())
	}
	if id != url.UserID {
		return c.String(http.StatusUnauthorized, "Access to selected url is denied")
	}
	url.FailedCall, err = h.URLStore.GetTotalFailure(url)
	if err != nil {
		return c.String(http.StatusConflict, err.Error())
	}
	jsonURL, err := json.Marshal(url)
	if err != nil {
		return c.String(http.StatusConflict, err.Error())
	}
	strJson := string(jsonURL)
	return c.String(http.StatusOK, strJson)
}
