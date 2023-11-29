package gormuser

import "github.com/mrbryside/rbh/domain/user/domain/user"

func (r Repository) Authenticate(email string, password string) (bool, user.Aggregate) {
	userAgg, err := r.GetByEmail(email)
	if err != nil {
		return false, user.Aggregate{}
	}
	return userAgg.Person().Password.ComparePassword(password), userAgg
}
