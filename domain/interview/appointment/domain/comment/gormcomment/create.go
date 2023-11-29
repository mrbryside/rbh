package gormcomment

import (
	"github.com/mrbryside/rbh/domain/interview/appointment/domain/comment"
)

func (r *repository) Create(agg comment.Aggregate) (comment.Aggregate, error) {
	c := toGormCommentModel(agg)
	if result := r.db.Create(&c); result.Error != nil {
		return comment.Aggregate{}, result.Error
	}
	return toAggregate(c), nil
}
