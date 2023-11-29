package service

import "github.com/mrbryside/rbh/domain/interview/appointment/domain/history"

//go:generate mockgen -source=history.go -destination=../pkg/generated/historymock/service.go -package=historymock
type HistoryServicer interface {
	Create(CreateHistoryDto) (history.Aggregate, error)
	GetAllByAppointmentId(appointmentId uint) ([]history.Aggregate, error)
}

type historyService struct {
	historyDomain history.Repository
}

func NewHistoryService(hr history.Repository) HistoryServicer {
	return historyService{
		historyDomain: hr,
	}
}

type CreateHistoryDto struct {
	Id          uint
	Name        string
	Description string
	Status      string
}

func (hs historyService) Create(cd CreateHistoryDto) (history.Aggregate, error) {
	agg, err := history.New(
		cd.Id,
		cd.Name,
		cd.Description,
		cd.Status,
	)
	if err != nil {
		return agg, err
	}
	return hs.historyDomain.Create(agg)
}

func (hs historyService) GetAllByAppointmentId(appointmentId uint) ([]history.Aggregate, error) {
	return hs.historyDomain.GetAllByAppointmentId(appointmentId)
}
