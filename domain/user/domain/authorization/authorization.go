package authorization

import (
	"github.com/mrbryside/rbh/pkg/claim"
)

const (
	Jwt = "jwt"
)

//go:generate mockgen -source=./authorization.go -destination=../../pkg/generated/mockgen/authorization.go -package=mockgen
type JwtAuthorizer interface {
	GenerateToken(userID uint, role string) (string, error)
	ParseToken(tokenString string) (*claim.CustomClaims, error)
}
