package history

import (
	"errors"

	"github.com/mrbryside/rbh/domain/interview/appointment/domain"
	"github.com/mrbryside/rbh/domain/interview/appointment/types/mystatus"
	"github.com/mrbryside/rbh/pkg/myint"
	"github.com/mrbryside/rbh/pkg/mystr"
)

type Aggregate struct {
	appointment domain.Appointment // root entity
}

func New(id uint, name string, description string, status string) (Aggregate, error) {
	if myint.IsZero(id) && mystr.IsEmpty(name) && mystr.IsEmpty(description) && mystr.IsEmpty(status) {
		return Aggregate{}, errors.New("error creating aggregate field is empty")
	}
	stat, err := mystatus.NewType(status)
	if err != nil {
		return Aggregate{}, err
	}
	return Aggregate{
		appointment: domain.Appointment{
			Id:          id,
			Name:        name,
			Description: description,
			Status:      stat,
		},
	}, nil
}

func (a Aggregate) Appointment() domain.Appointment {
	return a.appointment
}
