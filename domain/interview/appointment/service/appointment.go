package service

import (
	"github.com/labstack/gommon/log"
	"github.com/mrbryside/rbh/domain/interview/appointment/domain/appointment"
	"github.com/mrbryside/rbh/domain/interview/appointment/domain/comment"
	"github.com/mrbryside/rbh/domain/interview/appointment/domain/history"
)

//go:generate mockgen -source=appointment.go -destination=../pkg/generated/appointmentmock/service.go -package=appointmentmock
type AppointmentServicer interface {
	Create(CreateAppointmentDto) (appointment.Aggregate, error)
	UpdateById(UpdateAppointmentDto) (appointment.Aggregate, error)
	GetAll(GetAllAppointmentDto) (appointment.Aggregates, error)
	GetById(uint) (appointment.Aggregate, error)
}

type appointmentService struct {
	appointmentDomain appointment.Repository
	historyService    HistoryServicer
	creatorService    CreatorServicer
	commentService    CommentServicer
}

func NewAppointmentService(
	ar appointment.Repository,
	hs HistoryServicer,
	cs CreatorServicer,
	cos CommentServicer,
) AppointmentServicer {
	return appointmentService{
		appointmentDomain: ar,
		historyService:    hs,
		creatorService:    cs,
		commentService:    cos,
	}
}

// CreateAppointmentDto is the request struct for creating an appointment
type CreateAppointmentDto struct {
	Name        string
	Description string
	CreatorId   uint
}

func (as appointmentService) Create(dto CreateAppointmentDto) (appointment.Aggregate, error) {
	result, err := as.appointmentDomain.Create(
		appointment.
			New(dto.Name, dto.Description).
			SetCreatorId(dto.CreatorId),
	)
	if err != nil {
		return result, err
	}
	as.auditLog(result)

	return as.applyCreatorForAppointment(result)
}

// UpdateAppointmentDto is the request struct for updating an appointment
type UpdateAppointmentDto struct {
	Id          uint
	Name        string
	Description string
	Status      string
	Enabled     bool
}

func (as appointmentService) UpdateById(dto UpdateAppointmentDto) (appointment.Aggregate, error) {
	agg, err := appointment.
		New(dto.Name, dto.Description).
		SetId(dto.Id).
		SetEnabled(dto.Enabled).
		SetStatus(dto.Status)
	if err != nil {
		return agg, err
	}

	result, err := as.appointmentDomain.UpdateById(agg)
	if err != nil {
		return result, err
	}
	as.auditLog(result)

	return as.applyCreatorForAppointment(result)
}

// GetAllAppointmentDto is the request struct for getting all appointments
type GetAllAppointmentDto struct {
	Page     uint
	PageSize uint
}

func (as appointmentService) GetAll(dto GetAllAppointmentDto) (appointment.Aggregates, error) {
	aggs, err := as.appointmentDomain.GetAll(dto.Page, dto.PageSize)
	if err != nil {
		return aggs, err
	}

	return as.applyCreatorForAppointments(aggs)
}

func (as appointmentService) GetById(appointmentId uint) (appointment.Aggregate, error) {
	agg, err := as.appointmentDomain.GetById(appointmentId)
	if err != nil {
		return agg, err
	}
	aggWithCreator, err := as.applyCreatorForAppointment(agg)
	if err != nil {
		return aggWithCreator, err
	}
	aggWithComments, err := as.addComments(aggWithCreator)
	if err != nil {
		return aggWithComments, err
	}
	return as.addHistories(aggWithComments)

}

func (as appointmentService) applyCreatorForComment(cAgg comment.Aggregate) (comment.Aggregate, error) {
	creator, err := as.creatorService.GetById(cAgg.Comment().Creator.Id)
	if err != nil {
		return cAgg, err
	}
	return cAgg.SetCreator(creator.Creator()), nil
}

func (as appointmentService) addHistory(agg appointment.Aggregate, hAggs []history.Aggregate) (appointment.Aggregate, error) {
	for _, hAgg := range hAggs {
		agg = agg.AddHistory(hAgg.Appointment())
	}
	return agg, nil
}

func (as appointmentService) addHistories(agg appointment.Aggregate) (appointment.Aggregate, error) {
	hAggs, err := as.historyService.GetAllByAppointmentId(agg.Appointment().Id)
	if err != nil {
		return agg, err
	}
	return as.addHistory(agg, hAggs)
}

func (as appointmentService) addComment(agg appointment.Aggregate, cAggs []comment.Aggregate) (appointment.Aggregate, error) {
	for _, cAgg := range cAggs {
		cAggWithCreator, err := as.applyCreatorForComment(cAgg)
		if err != nil {
			return agg, err
		}
		agg = agg.AddComment(cAggWithCreator.Comment())
	}
	return agg, nil
}

func (as appointmentService) addComments(agg appointment.Aggregate) (appointment.Aggregate, error) {
	cAggs, err := as.commentService.GetAllByAppointmentId(agg.Appointment().Id)
	if err != nil {
		return agg, err
	}

	return as.addComment(agg, cAggs)
}

func (as appointmentService) applyCreatorForAppointment(agg appointment.Aggregate) (appointment.Aggregate, error) {
	creatorAgg, err := as.creatorService.GetById(agg.Creator().Id)
	if err != nil {
		return agg, err
	}
	return agg.SetCreator(creatorAgg.Creator()), nil
}

func (as appointmentService) applyCreatorForAppointments(aggs appointment.Aggregates) (appointment.Aggregates, error) {
	results := appointment.News(aggs.Next)
	for _, agg := range aggs.Aggregates {
		aggWithCreator, err := as.applyCreatorForAppointment(agg)
		if err != nil {
			return aggs, err
		}
		results.Aggregates = append(results.Aggregates, aggWithCreator)
	}
	return results, nil
}

func (as appointmentService) auditLog(agg appointment.Aggregate) {
	_, err := as.historyService.Create(CreateHistoryDto{
		Id:          agg.Appointment().Id,
		Name:        agg.Appointment().Name,
		Description: agg.Appointment().Description,
		Status:      agg.Appointment().Status.Value,
	})
	if err != nil {
		log.Printf("error creating history: %s", err.Error())
	}
}
