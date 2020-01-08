package handler

import (
	"github.com/labstack/echo"
	"github.com/saman2000hoseini/http-monitor/model"
)

type userRegisterRequest struct {
	User struct {
		Username string `json:"username"`
		Password string `json:"password"`
	} `json:"user"`
}

func (r *userRegisterRequest) bind(c echo.Context, u *model.User) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	u.Username = r.User.Username
	if err := u.HashPassword(r.User.Password); err != nil {
		return err
	}
	return nil
}

type createURLRequest struct {
	Address string `json:"address"`
}

func (r *createURLRequest) bind(c echo.Context, u *model.URL) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	u.Address = r.Address
	return nil
}
