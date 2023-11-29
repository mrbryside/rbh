package gormcomment

import "github.com/mrbryside/rbh/domain/interview/appointment/domain/comment"

func (r repository) GetAllByAppointmentId(appointmentId uint) ([]comment.Aggregate, error) {
	comments, err := r.fetchCommentsByAppointmentId(appointmentId)
	if err != nil {
		return []comment.Aggregate{}, err
	}
	return toAggregates(comments), nil
}

func (r repository) fetchCommentsByAppointmentId(appointmentId uint) ([]Comment, error) {
	var comments []Comment
	err := r.db.Where("appointment_id = ?", appointmentId).Find(&comments).Error
	if err != nil {
		return comments, err
	}
	return comments, nil
}
