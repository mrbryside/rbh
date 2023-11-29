package myrole

import "errors"

const (
	PeopleTeam = "PeopleTeam"
	Guest      = "Guest"
)

type Type struct {
	Value string
}

func NewType(rn string) (Type, error) {
	if rn != PeopleTeam && rn != Guest {
		return Type{}, errors.New("invalid role name")
	}
	return Type{Value: rn}, nil
}
