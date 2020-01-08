package router

import (
	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9"
)

func Router() *echo.Echo {
	e := echo.New()
	e.Validator = newValidator()
	return e
}

func newValidator() *Validator {
	return &Validator{
		validator: validator.New(),
	}
}

type Validator struct {
	validator *validator.Validate
}

func (v *Validator) Validate(i interface{}) error {
	return v.validator.Struct(i)
}
