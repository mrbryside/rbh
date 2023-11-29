package gormuser

import (
	"github.com/mrbryside/rbh/domain/user/domain/user"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Email    string `gorm:"unique;not null"`
	Password string
	Role     string
}

func toGormUserModel(aggregate user.Aggregate) User {
	return User{
		Name:     aggregate.Person().Name,
		Email:    aggregate.Person().Email.Value,
		Password: aggregate.Person().Password.Hash,
		Role:     aggregate.Person().Role.Value,
	}
}

func toAggregate(u User) (user.Aggregate, error) {
	userAgg, err := user.New(u.Name, u.Email, u.Password)
	if err != nil {
		return user.Aggregate{}, err
	}
	userAggWithRole, err := userAgg.SetRole(u.Role)
	if err != nil {
		return user.Aggregate{}, err
	}

	return userAggWithRole.SetId(u.ID), nil
}
