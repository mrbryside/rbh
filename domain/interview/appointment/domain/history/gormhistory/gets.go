package gormhistory

import "github.com/mrbryside/rbh/domain/interview/appointment/domain/history"

func (r repository) GetAllByAppointmentId(appointmentId uint) ([]history.Aggregate, error) {
	var histories []History
	err := r.db.Where("appointment_id = ?", appointmentId).Find(&histories).Error
	if err != nil {
		return []history.Aggregate{}, err
	}

	return toAggregates(histories)
}
