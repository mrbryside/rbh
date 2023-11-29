package auth

import (
	"github.com/labstack/echo/v4"
	"github.com/mrbryside/rbh/domain/user/service"
)

type AuthHandler struct {
	authService service.AuthServicer
}

func NewAuthHandler(as service.AuthServicer) AuthHandler {
	return AuthHandler{
		authService: as,
	}
}

func (ah AuthHandler) RegisterRoutes(e *echo.Echo) {
	e.POST("/auth/register", ah.RegisterHandler)
	e.POST("/auth/login", ah.LoginHandler)
}
