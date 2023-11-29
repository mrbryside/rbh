package mhttp

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

const (
	successCreateCode = 2001
	successCode       = 2000

	createdMessage = "created"
	successMessage = "success"
)

func SuccessCreatedWithBody(ctx echo.Context, body interface{}) error {
	resp := response{
		Code:    successCreateCode,
		Message: createdMessage,
		Data:    body,
	}
	return ctx.JSON(http.StatusCreated, resp)
}

func SuccessWithBody(ctx echo.Context, body interface{}) error {
	resp := response{
		Code:    successCode,
		Message: successMessage,
		Data:    body,
	}
	return ctx.JSON(http.StatusOK, resp)
}

func SuccessWithMessage(ctx echo.Context, message string) error {
	resp := response{
		Code:    successCode,
		Message: message,
	}
	return ctx.JSON(http.StatusOK, resp)
}
