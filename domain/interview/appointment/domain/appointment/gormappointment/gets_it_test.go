//go:build integration

package gormappointment_test

import (
	"testing"

	"github.com/mrbryside/rbh/domain/interview/appointment/domain/appointment/gormappointment"
	"github.com/mrbryside/rbh/pkg/db/mygorm"
	"github.com/stretchr/testify/assert"
)

func TestGetAllAppointmentWithPaginationSuccess(t *testing.T) {
	// Arrange
	db, close, err := mygorm.BasicConnection()
	if err != nil {
		t.Error(err)
	}
	defer close()
	db.Exec("DELETE FROM appointments")

	repo := gormappointment.NewRepository(db)

	createAppointment()

	// Act
	result, err := repo.GetAll(1, 1)
	if err != nil {
		t.Error(err.Error())
	}
	assert.Equal(t, 1, len(result.Aggregates))
}

func TestGetAllAppointmentWithPaginationWithNextFalseSuccess(t *testing.T) {
	// Arrange
	db, close, err := mygorm.BasicConnection()
	if err != nil {
		t.Error(err)
	}
	defer close()
	db.Exec("DELETE FROM appointments")

	repo := gormappointment.NewRepository(db)
	createAppointment()
	createAppointment()

	// Act
	result, err := repo.GetAll(2, 1)
	if err != nil {
		t.Error(err.Error())
	}
	assert.Equal(t, 1, len(result.Aggregates))
	assert.False(t, result.Next)
}

func TestGetAllAppointmentWithPaginationWithNextFalseWithFullSizeSuccess(t *testing.T) {
	// Arrange
	db, close, err := mygorm.BasicConnection()
	if err != nil {
		t.Error(err)
	}
	defer close()
	db.Exec("DELETE FROM appointments")

	repo := gormappointment.NewRepository(db)
	createAppointment()
	createAppointment()

	// Act
	result, err := repo.GetAll(1, 2)
	if err != nil {
		t.Error(err.Error())
	}
	assert.Equal(t, 2, len(result.Aggregates))
	assert.False(t, result.Next)
}
