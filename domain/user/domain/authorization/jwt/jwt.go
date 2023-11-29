package jwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/mrbryside/rbh/domain/user/domain/authorization"
	"github.com/mrbryside/rbh/pkg/claim"
)

type Jwt struct {
	SecretKey []byte
}

func NewJwt(secretKey string) authorization.JwtAuthorizer {
	return &Jwt{SecretKey: []byte(secretKey)}
}

func (j *Jwt) GenerateToken(userID uint, role string) (string, error) {
	claims := &claim.CustomClaims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(j.SecretKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (j *Jwt) ParseToken(tokenString string) (*claim.CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &claim.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SecretKey, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*claim.CustomClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}
