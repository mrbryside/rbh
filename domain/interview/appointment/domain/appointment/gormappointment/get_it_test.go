//go:build integration

package gormappointment_test

import (
	"testing"

	"github.com/mrbryside/rbh/domain/interview/appointment/domain/appointment"
	"github.com/mrbryside/rbh/domain/interview/appointment/domain/appointment/gormappointment"
	"github.com/mrbryside/rbh/pkg/db/mygorm"
	"github.com/stretchr/testify/assert"
)

func TestGetAppointmentById(t *testing.T) {
	// Arrange
	db, close, err := mygorm.BasicConnection()
	if err != nil {
		t.Error(err)
	}
	defer close()
	db.Exec("DELETE FROM appointments")

	repo := gormappointment.NewRepository(db)
	agg := appointment.New("name", "description")
	aggWithCreator := agg.SetCreatorId(1)

	created_agg, err := repo.Create(aggWithCreator)
	if err != nil {
		t.Error(err.Error())
	}

	// Act
	result, err := repo.GetById(created_agg.Appointment().Id)
	if err != nil {
		t.Error(err.Error())
	}

	assert.Equal(t, created_agg.Appointment().Id, result.Appointment().Id)
	assert.Equal(t, created_agg.Appointment().Name, result.Appointment().Name)
}
