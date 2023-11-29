package user

import (
	"errors"

	"github.com/mrbryside/rbh/domain/user/domain"
	"github.com/mrbryside/rbh/domain/user/types/myemail"
	"github.com/mrbryside/rbh/domain/user/types/mypass"
	"github.com/mrbryside/rbh/domain/user/types/myrole"
	"github.com/mrbryside/rbh/pkg/mystr"
)

type Aggregate struct {
	person domain.Person // root_entity
}

func New(name string, email string, password string) (Aggregate, error) {
	if mystr.IsEmpty(name) {
		return Aggregate{}, errors.New("name cannot be empty")
	}
	verifiedEmail, err := myemail.NewType(email)
	if err != nil {
		return Aggregate{}, err
	}
	verifiedPassword, err := mypass.NewType(password)
	if err != nil {
		return Aggregate{}, err
	}

	c := domain.Person{
		Name:     name,
		Email:    verifiedEmail,
		Password: verifiedPassword,
	}

	return Aggregate{person: c}, nil
}

func (a Aggregate) Person() domain.Person {
	return a.person
}

func (a Aggregate) SetId(id uint) Aggregate {
	c := domain.Person{
		Id:       id,
		Name:     a.person.Name,
		Email:    a.person.Email,
		Password: a.person.Password,
		Role:     a.person.Role,
	}
	return Aggregate{person: c}
}

func (a Aggregate) SetRole(roleName string) (Aggregate, error) {
	verifiedRole, err := myrole.NewType(roleName)
	if err != nil {
		return a, err
	}
	c := domain.Person{
		Id:       a.person.Id,
		Name:     a.person.Name,
		Email:    a.person.Email,
		Password: a.person.Password,
		Role:     verifiedRole}

	return Aggregate{person: c}, nil
}
