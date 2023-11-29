package gormappointment

import (
	"github.com/mrbryside/rbh/domain/interview/appointment/domain/appointment"
)

func (r Repository) UpdateById(agg appointment.Aggregate) (appointment.Aggregate, error) {
	a := toGormAppointmentModel(agg)
	ignoredFields := []string{"user_id"}
	result := r.db.Model(&Appointment{}).
		Where("id = ?", agg.Appointment().Id).
		Omit(ignoredFields...).
		Updates(a).
		Update("enabled", a.Enabled)

	if result.Error != nil {
		return appointment.Aggregate{}, result.Error
	}

	return r.GetById(agg.Appointment().Id)
}
