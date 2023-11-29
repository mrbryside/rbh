package gormappointment

import (
	"github.com/mrbryside/rbh/domain/interview/appointment/domain/appointment"
)

func (r Repository) GetAll(page, pageSize uint) (appointment.Aggregates, error) {
	results, err := r.customQuery.paginateAppointmentQueryWithEnabled(int(page), int(pageSize))
	if err != nil {
		return appointment.Aggregates{}, err
	}

	return toAggregates(results)
}
