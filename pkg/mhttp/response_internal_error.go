package mhttp

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

const (
	internalErrorCode = 5000
)

func InternalError(ctx echo.Context, message string, args ...interface{}) error {
	msg := fmt.Sprintf(message, args...)
	resp := response{
		Code:    internalErrorCode,
		Message: fmt.Sprintf("internal server error: %s", msg),
	}
	return ctx.JSON(http.StatusInternalServerError, resp)
}
