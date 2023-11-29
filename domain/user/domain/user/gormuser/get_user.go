package gormuser

import (
	"github.com/mrbryside/rbh/domain/user/domain/user"
)

func (r Repository) GetById(userId uint) (user.Aggregate, error) {
	var u User
	result := r.db.First(&u, userId)
	if result.Error != nil {
		return user.Aggregate{}, result.Error
	}

	return toAggregate(u)
}

func (r Repository) GetByEmail(email string) (user.Aggregate, error) {
	var u User
	result := r.db.Where("email = ?", email).First(&u)
	if result.Error != nil {
		return user.Aggregate{}, result.Error
	}

	return toAggregate(u)
}
