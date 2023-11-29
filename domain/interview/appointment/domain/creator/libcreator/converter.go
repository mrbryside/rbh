package libcreator

import (
	"github.com/mrbryside/rbh/domain/interview/appointment/domain/creator"
	"github.com/mrbryside/rbh/domain/user/domain/user"
)

func toAggregate(result user.Aggregate) (creator.Aggregate, error) {
	creatorAgg, err := creator.New(
		result.Person().Id,
		result.Person().Name,
		result.Person().Email.Value,
	)
	if err != nil {
		return creator.Aggregate{}, err
	}
	return creatorAgg, nil
}
