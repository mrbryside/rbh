package mhttp

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

const (
	unauthorizedCode = 4000
)

func ErrorWithUnauthorized(ctx echo.Context) error {
	resp := response{
		Code:    unauthorizedCode,
		Message: "unauthorized to access this resource",
	}
	return ctx.JSON(http.StatusUnauthorized, resp)
}
