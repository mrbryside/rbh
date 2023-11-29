package gormhistory

import "github.com/mrbryside/rbh/domain/interview/appointment/domain/history"

func (r repository) Create(agg history.Aggregate) (history.Aggregate, error) {
	h := toGormHistoryModel(agg)
	if result := r.db.Create(&h); result.Error != nil {
		return history.Aggregate{}, result.Error
	}
	return toAggregate(h)
}
