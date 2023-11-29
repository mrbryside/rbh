package auth

import (
	"github.com/labstack/echo/v4"
	"github.com/mrbryside/rbh/pkg/mhttp"
)

func (ah AuthHandler) LoginHandler(c echo.Context) error {
	var payload loginPayload
	err := c.Bind(&payload)
	if err != nil {
		return mhttp.BadRequest(c, "error binding payload: %v", err.Error())
	}

	token, err := ah.authService.Login(payload.Email, payload.Password)
	if err != nil {
		return mhttp.InternalError(c, "error generate token: %v", err.Error())
	}
	return mhttp.SuccessWithBody(c, loginResponse{AccessToken: token, RefreshToken: "coming soon"})
}
