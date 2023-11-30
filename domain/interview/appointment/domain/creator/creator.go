package creator

import (
	"errors"

	"github.com/mrbryside/rbh/domain/interview/appointment/domain"
	"github.com/mrbryside/rbh/pkg/myint"
)

type Aggregate struct {
	creator domain.Creator // Root Entity
}

func New(id uint, name string, email string) (Aggregate, error) {
	if myint.IsZero(id) {
		return Aggregate{}, errors.New("id cannot be Empty")
	}
	return Aggregate{
		creator: domain.Creator{
			Id:    id,
			Name:  name,
			Email: email,
		}}, nil
}

func (c Aggregate) Creator() domain.Creator {
	return c.creator
}
