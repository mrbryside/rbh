package gormappointment

import (
	"github.com/mrbryside/rbh/pkg/pagination"
	"gorm.io/gorm"
)

type CustomQuery struct {
	db *gorm.DB
}

func (cq CustomQuery) paginateAppointmentQueryWithEnabled(page int, pageSize int) (PaginatedAppointmentResponse, error) {
	totalRecords, err := cq.countAppointmentTotalRecordWithEnabled()
	if err != nil {
		return PaginatedAppointmentResponse{}, err
	}
	offset := pagination.CalculateOffset(page, pageSize)

	var appointments []Appointment
	if err := cq.db.Limit(pageSize).Offset(offset).Where("enabled IS TRUE").Find(&appointments).Error; err != nil {
		return PaginatedAppointmentResponse{}, err
	}

	return PaginatedAppointmentResponse{
		Appointments: appointments,
		Next:         pagination.IsNextPage(totalRecords, page, pageSize),
	}, nil

}

func (cq CustomQuery) countAppointmentTotalRecordWithEnabled() (int64, error) {
	var totalRecords int64
	var model Appointment

	if err := cq.db.Model(model).
		Where("enabled IS TRUE").
		Count(&totalRecords).Error; err != nil {
		return 0, err
	}
	return totalRecords, nil
}

type PaginatedAppointmentResponse struct {
	Appointments []Appointment `json:"appointments"`
	Next         bool          `json:"next"`
}
