package gormappointment

import (
	"github.com/mrbryside/rbh/domain/interview/appointment/domain/appointment"
	"gorm.io/gorm"
)

type Repository struct {
	db          *gorm.DB
	customQuery CustomQuery
}

func NewRepository(db *gorm.DB) appointment.Repository {
	db.AutoMigrate(&Appointment{})
	return Repository{
		db:          db,
		customQuery: CustomQuery{db: db},
	}
}
