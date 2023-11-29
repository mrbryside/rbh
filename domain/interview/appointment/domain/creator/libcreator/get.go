package libcreator

import "github.com/mrbryside/rbh/domain/interview/appointment/domain/creator"

func (r *repository) GetById(creatorId uint) (creator.Aggregate, error) {
	result, err := r.authService.GetById(creatorId)
	if err != nil {
		return creator.Aggregate{}, err
	}

	return toAggregate(result)
}
