package server

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type CommonValidator struct {
	validator *validator.Validate
}

func NewCommonValidator(validate *validator.Validate) *CommonValidator {
	return &CommonValidator{
		validator: validate,
	}
}

func (c CommonValidator) Validate(i interface{}) error {
	if err := c.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]string{"remarks": err.Error()})
	}

	return nil
}
