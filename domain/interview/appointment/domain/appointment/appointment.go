package appointment

import (
	"github.com/mrbryside/rbh/domain/interview/appointment/domain"
	"github.com/mrbryside/rbh/domain/interview/appointment/types/mystatus"
)

type Aggregates struct {
	Aggregates []Aggregate
	Next       bool
}

type Aggregate struct {
	appointment domain.Appointment   // root entity
	creator     domain.Creator       // entity
	comments    []domain.Comment     // entity
	histories   []domain.Appointment // entity
}

func News(next bool) Aggregates {
	return Aggregates{
		Next: next,
	}
}

func New(name, description string) Aggregate {
	status, _ := mystatus.NewType(mystatus.Todo)
	return Aggregate{
		appointment: domain.Appointment{
			Name:        name,
			Description: description,
			Status:      status,
			Enabled:     true,
		},
		creator:   domain.Creator{},
		comments:  []domain.Comment{},
		histories: []domain.Appointment{},
	}
}

func (agg Aggregate) Appointment() domain.Appointment {
	return agg.appointment
}

func (agg Aggregate) Creator() domain.Creator {
	return agg.creator
}

func (agg Aggregate) Comments() []domain.Comment {
	return agg.comments
}

func (agg Aggregate) Histories() []domain.Appointment {
	return agg.histories
}

func (agg Aggregate) SetId(id uint) Aggregate {
	return Aggregate{
		appointment: domain.Appointment{
			Id:          id,
			Name:        agg.appointment.Name,
			Description: agg.appointment.Description,
			Status:      agg.appointment.Status,
			Enabled:     agg.appointment.Enabled,
			CreatedAt:   agg.appointment.CreatedAt,
		},
		creator:   agg.creator,
		comments:  agg.comments,
		histories: agg.histories,
	}
}

func (agg Aggregate) SetCreatorId(id uint) Aggregate {
	return Aggregate{
		appointment: agg.appointment,
		creator: domain.Creator{
			Id:    id,
			Name:  agg.creator.Name,
			Email: agg.creator.Email,
		},
		comments:  agg.comments,
		histories: agg.histories,
	}
}

func (agg Aggregate) SetCreator(c domain.Creator) Aggregate {
	return Aggregate{
		appointment: agg.appointment,
		creator:     c,
		comments:    agg.comments,
		histories:   agg.histories,
	}
}

func (agg Aggregate) SetEnabled(flag bool) Aggregate {
	return Aggregate{
		appointment: domain.Appointment{
			Id:          agg.appointment.Id,
			Name:        agg.appointment.Name,
			Description: agg.appointment.Description,
			Status:      agg.appointment.Status,
			CreatedAt:   agg.appointment.CreatedAt,
			Enabled:     flag,
		},
		creator:   agg.creator,
		comments:  agg.comments,
		histories: agg.histories,
	}
}

func (agg Aggregate) SetStatus(status string) (Aggregate, error) {
	stat, err := mystatus.NewType(status)
	if err != nil {
		return Aggregate{}, err
	}
	return Aggregate{
		appointment: domain.Appointment{
			Id:          agg.appointment.Id,
			Name:        agg.appointment.Name,
			Description: agg.appointment.Description,
			Status:      stat,
			CreatedAt:   agg.appointment.CreatedAt,
			Enabled:     agg.appointment.Enabled,
		},
		creator:   agg.creator,
		comments:  agg.comments,
		histories: agg.histories,
	}, nil
}

func (agg Aggregate) SetTimeStamps(created_at string) Aggregate {
	return Aggregate{
		appointment: domain.Appointment{
			Id:          agg.appointment.Id,
			Name:        agg.appointment.Name,
			Description: agg.appointment.Description,
			Status:      agg.appointment.Status,
			CreatedAt:   created_at,
			Enabled:     agg.appointment.Enabled,
		},
		creator:   agg.creator,
		comments:  agg.comments,
		histories: agg.histories,
	}
}

func (agg Aggregate) AddComment(c domain.Comment) Aggregate {
	comments := agg.comments
	return Aggregate{
		appointment: agg.appointment,
		creator:     agg.creator,
		comments:    append(comments, c),
		histories:   agg.histories,
	}
}

func (agg Aggregate) AddHistory(h domain.Appointment) Aggregate {
	histories := agg.histories
	return Aggregate{
		appointment: agg.appointment,
		creator:     agg.creator,
		comments:    agg.comments,
		histories:   append(histories, h),
	}
}
