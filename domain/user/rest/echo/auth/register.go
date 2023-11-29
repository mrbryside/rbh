package auth

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/mrbryside/rbh/pkg/mhttp"
)

func (ah AuthHandler) RegisterHandler(c echo.Context) error {
	var payload registerPayload
	err := c.Bind(&payload)
	if err != nil {
		return mhttp.BadRequest(c, fmt.Sprintf("error binding payload: %v", err.Error()))
	}
	token, err := ah.authService.Register(payload.Name, payload.Email, payload.Password)
	if err != nil {
		return mhttp.InternalError(c, fmt.Sprintf("error register: %v", err.Error()))

	}
	return mhttp.SuccessWithBody(c, registerResponse{AccessToken: token, RefreshToken: "coming soon"})
}
