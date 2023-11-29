package mypass

import (
	"errors"
	"regexp"

	"github.com/mrbryside/rbh/pkg/mystr"
	"golang.org/x/crypto/bcrypt"
)

type Type struct {
	Hash    string
	NonHash string
}

func NewType(password string) (Type, error) {
	if mystr.IsEmpty(password) {
		return Type{}, errors.New("password cannot be empty")
	}

	if !isValidHash(password) {
		hashedPasswordBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			return Type{}, err
		}
		return Type{Hash: string(hashedPasswordBytes)}, nil
	}
	return Type{Hash: password, NonHash: ""}, nil
}

func (t Type) ComparePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(t.Hash), []byte(password))
	return err == nil
}

func isValidHash(hash string) bool {
	hashRegex := regexp.MustCompile(`^\$2[aby]\$[0-9]{2}\$[A-Za-z0-9./]{53}$`)
	return hashRegex.MatchString(hash)
}
