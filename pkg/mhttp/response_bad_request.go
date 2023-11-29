package mhttp

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

const (
	badRequestCode = 4000
)

func BadRequest(ctx echo.Context, message string, args ...interface{}) error {
	msg := fmt.Sprintf(message, args...)
	resp := response{
		Code:    badRequestCode,
		Message: fmt.Sprintf("bad request: %s", msg),
	}
	return ctx.JSON(http.StatusBadRequest, resp)
}
