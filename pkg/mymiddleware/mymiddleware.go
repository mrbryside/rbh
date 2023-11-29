package mymiddleware

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/mrbryside/rbh/pkg/claim"
)

func ExtractFromClaims(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(*claim.CustomClaims)

		c.Set("user_id", claims.UserID)
		c.Set("role", claims.Role)
		return next(c)
	}
}
