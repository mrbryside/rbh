package gormuser

import (
	"github.com/mrbryside/rbh/domain/user/domain/user"
)

func (r Repository) Create(uag user.Aggregate) (user.Aggregate, error) {
	u := toGormUserModel(uag)
	if result := r.db.Create(&u); result.Error != nil {
		return user.Aggregate{}, result.Error
	}

	return toAggregate(u)
}
