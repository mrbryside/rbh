package gormcomment

import "github.com/mrbryside/rbh/domain/interview/appointment/domain/comment"

func (r repository) GetById(id uint) (comment.Aggregate, error) {
	var c Comment
	if err := r.db.First(&c, id).Error; err != nil {
		return comment.Aggregate{}, err
	}
	return toAggregate(c), nil
}
