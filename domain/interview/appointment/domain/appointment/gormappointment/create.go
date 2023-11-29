package gormappointment

import (
	"github.com/mrbryside/rbh/domain/interview/appointment/domain/appointment"
)

func (r Repository) Create(agg appointment.Aggregate) (appointment.Aggregate, error) {
	a := toGormAppointmentModel(agg)
	if result := r.db.Create(&a); result.Error != nil {
		return appointment.Aggregate{}, result.Error
	}
	return toAggregate(a)
}
