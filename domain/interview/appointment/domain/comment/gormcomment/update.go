package gormcomment

import (
	"github.com/mrbryside/rbh/domain/interview/appointment/domain/comment"
)

func (r repository) UpdateById(agg comment.Aggregate) (comment.Aggregate, error) {
	result := r.db.Model(&Comment{}).Where("id = ?", agg.Comment().Id).Update("message", agg.Comment().Message)
	if result.Error != nil {
		return comment.Aggregate{}, result.Error
	}

	return r.GetById(agg.Comment().Id)
}
