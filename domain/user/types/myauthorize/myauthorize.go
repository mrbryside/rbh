package myauthorize

import "errors"

const (
	Jwt = "Jwt"
)

type Type struct {
	Value string
}

func NewType(authorizationType string) (Type, error) {
	if authorizationType != Jwt {
		return Type{}, errors.New("invalid authorization type")
	}
	return Type{Value: authorizationType}, nil
}
