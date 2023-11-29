package gormappointment

import (
	"github.com/mrbryside/rbh/domain/interview/appointment/domain/appointment"
)

func (r Repository) GetById(appointmentId uint) (appointment.Aggregate, error) {
	var a Appointment
	result := r.db.First(&a, appointmentId)
	if result.Error != nil {
		return appointment.Aggregate{}, result.Error
	}

	return toAggregate(a)
}
