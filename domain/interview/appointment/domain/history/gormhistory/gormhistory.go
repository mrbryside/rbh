package gormhistory

import (
	"github.com/mrbryside/rbh/domain/interview/appointment/domain/history"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) history.Repository {
	db.AutoMigrate(&History{})
	return repository{db}
}
