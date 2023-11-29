package gormcomment

import (
	"github.com/mrbryside/rbh/domain/interview/appointment/domain/comment"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) comment.Repository {
	db.AutoMigrate(&Comment{})
	return &repository{db: db}
}
