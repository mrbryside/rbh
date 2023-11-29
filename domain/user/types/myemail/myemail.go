package myemail

import (
	"errors"
	"regexp"
)

type Type struct {
	Value string
}

func NewType(value string) (Type, error) {
	if !verifyFormat(value) {
		return Type{}, errors.New("invalid email format")
	}
	return Type{Value: value}, nil
}

func verifyFormat(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}
