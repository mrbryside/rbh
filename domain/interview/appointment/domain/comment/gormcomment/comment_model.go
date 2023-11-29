package gormcomment

import (
	"github.com/mrbryside/rbh/domain/interview/appointment/domain/comment"
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Message       string
	AppointmentID uint
	UserID        uint
	Appointment   Appointment `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	User          User        `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type User struct {
	gorm.Model
}

type Appointment struct {
	gorm.Model
}

func toAggregate(c Comment) comment.Aggregate {
	agg := comment.New(c.Message)
	return agg.SetAppointmentId(c.AppointmentID).
		SetCreatorId(c.UserID).
		SetTimestamps(c.CreatedAt.String()).
		SetCommentId(c.ID)
}

func toAggregates(cs []Comment) []comment.Aggregate {
	var results []comment.Aggregate
	for _, c := range cs {
		results = append(results, toAggregate(c))
	}
	return results
}

func toGormCommentModel(agg comment.Aggregate) Comment {
	return Comment{
		Message:       agg.Comment().Message,
		AppointmentID: agg.Appointment().Id,
		UserID:        agg.Comment().Creator.Id,
	}
}
