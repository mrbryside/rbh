package gormappointment

import (
	"github.com/mrbryside/rbh/domain/interview/appointment/domain/appointment"
	"gorm.io/gorm"
)

type Appointment struct {
	gorm.Model
	Name        string
	Description string
	Status      string
	Enabled     bool
	UserID      uint
	User        User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type User struct {
	gorm.Model
}

func toGormAppointmentModel(agg appointment.Aggregate) Appointment {
	return Appointment{
		Name:        agg.Appointment().Name,
		Description: agg.Appointment().Description,
		Status:      agg.Appointment().Status.Value,
		Enabled:     agg.Appointment().Enabled,
		UserID:      agg.Creator().Id,
	}
}

func toAggregate(a Appointment) (appointment.Aggregate, error) {
	appointmentAgg := appointment.New(a.Name, a.Description)
	appointmentAggWithStatus, err := appointmentAgg.SetStatus(a.Status)
	if err != nil {
		return appointment.Aggregate{}, err
	}
	return appointmentAggWithStatus.
		SetEnabled(a.Enabled).
		SetTimeStamps(a.CreatedAt.String()).
		SetId(a.ID).
		SetCreatorId(a.UserID), nil
}

func toAggregates(pa PaginatedAppointmentResponse) (appointment.Aggregates, error) {
	results := appointment.News(pa.Next)
	for _, a := range pa.Appointments {
		agg, error := toAggregate(a)
		if error != nil {
			return appointment.Aggregates{}, error
		}
		results.Aggregates = append(results.Aggregates, agg)
	}
	return results, nil
}
