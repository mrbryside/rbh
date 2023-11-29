package comment

import "github.com/mrbryside/rbh/domain/interview/appointment/domain"

type Aggregate struct {
	comment     domain.Comment     // root entity
	appointment domain.Appointment // entity
}

func New(message string) Aggregate {
	return Aggregate{comment: domain.Comment{
		Message: message,
		Creator: domain.Creator{},
	}}
}

func (a Aggregate) Appointment() domain.Appointment {
	return a.appointment
}

func (a Aggregate) Comment() domain.Comment {
	return a.comment
}

func (a Aggregate) SetCommentId(id uint) Aggregate {
	return Aggregate{
		comment: domain.Comment{
			Id:        id,
			Message:   a.Comment().Message,
			CreatedAt: a.Comment().CreatedAt,
			Creator:   a.Comment().Creator,
		},
		appointment: a.appointment,
	}
}

func (a Aggregate) SetCreator(p domain.Creator) Aggregate {
	return Aggregate{
		comment: domain.Comment{
			Id:        a.Comment().Id,
			Message:   a.Comment().Message,
			CreatedAt: a.Comment().CreatedAt,
			Creator:   p,
		},
		appointment: a.appointment,
	}
}

func (a Aggregate) SetCreatorId(id uint) Aggregate {
	return Aggregate{
		comment: domain.Comment{
			Id:        a.Comment().Id,
			Message:   a.Comment().Message,
			CreatedAt: a.Comment().CreatedAt,
			Creator: domain.Creator{
				Id:    id,
				Name:  a.Comment().Creator.Name,
				Email: a.Comment().Creator.Email,
			},
		},
		appointment: a.appointment,
	}
}

func (a Aggregate) SetTimestamps(createdAt string) Aggregate {
	return Aggregate{
		comment: domain.Comment{
			Id:        a.Comment().Id,
			Message:   a.Comment().Message,
			CreatedAt: createdAt,
			Creator:   a.Comment().Creator,
		},
		appointment: a.appointment,
	}
}

func (a Aggregate) SetAppointmentId(id uint) Aggregate {
	return Aggregate{
		comment: a.comment,
		appointment: domain.Appointment{
			Id:          id,
			Name:        a.Appointment().Name,
			Description: a.Appointment().Description,
			Status:      a.Appointment().Status,
			Enabled:     a.Appointment().Enabled,
			CreatedAt:   a.Appointment().CreatedAt,
		},
	}
}
