package claim

import "github.com/golang-jwt/jwt/v5"

const (
	UserId = "user_id"
	Role   = "role"
)

type CustomClaims struct {
	UserID uint   `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}
