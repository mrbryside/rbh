package gormhistory

import (
	"github.com/mrbryside/rbh/domain/interview/appointment/domain/history"
	"gorm.io/gorm"
)

type History struct {
	gorm.Model
	Name          string
	Description   string
	Status        string
	AppointmentID uint
	Appointment   Appointment `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Appointment struct {
	gorm.Model
}

func toAggregate(h History) (history.Aggregate, error) {
	result, err := history.New(h.ID, h.Name, h.Description, h.Status)
	if err != nil {
		return history.Aggregate{}, err
	}
	return result, nil
}
func toAggregates(hs []History) ([]history.Aggregate, error) {
	var results []history.Aggregate
	for _, h := range hs {
		result, err := toAggregate(h)
		if err != nil {
			return []history.Aggregate{}, err
		}
		results = append(results, result)
	}
	return results, nil
}

func toGormHistoryModel(agg history.Aggregate) History {
	return History{
		Name:          agg.Appointment().Name,
		Description:   agg.Appointment().Description,
		Status:        agg.Appointment().Status.Value,
		AppointmentID: agg.Appointment().Id,
	}
}
