package service

import (
	"github.com/mrbryside/rbh/domain/user/domain/authorization"
)

//go:generate mockgen -source=./jwt.go -destination=../pkg/generated/mockgen/jwt_service.go -package=mockgen
type JwtServicer interface {
	GenerateToken(userId uint, role string) (string, error)
}

type JwtService struct {
	jwtDomain authorization.JwtAuthorizer
}

func NewJwtService(ja authorization.JwtAuthorizer) JwtService {
	return JwtService{jwtDomain: ja}
}

func (js JwtService) GenerateToken(userId uint, role string) (string, error) {
	token, err := js.jwtDomain.GenerateToken(userId, role)
	if err != nil {
		return "", err
	}
	return token, nil
}
