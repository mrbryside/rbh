//go:build integration

package gormappointment_test

import (
	"testing"

	"github.com/mrbryside/rbh/domain/interview/appointment/domain/appointment"
	"github.com/mrbryside/rbh/domain/interview/appointment/domain/appointment/gormappointment"
	"github.com/mrbryside/rbh/domain/interview/appointment/types/mystatus"
	"github.com/mrbryside/rbh/pkg/db/mygorm"
	"github.com/stretchr/testify/assert"
)

func TestUpdateAppointment(t *testing.T) {
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

	updateAgg := appointment.New("new name", "new description")
	updateAggWithId := updateAgg.SetId(created_agg.Appointment().Id)
	updateAggWithEnabled := updateAggWithId.SetEnabled(false)
	updateAggWithStatus, err := updateAggWithEnabled.SetStatus(mystatus.Done)
	if err != nil {
		t.Error(err.Error())
	}

	// Act
	result, err := repo.UpdateById(updateAggWithStatus)
	if err != nil {
		t.Error(err.Error())
	}

	// Assert
	assert.Equal(t, false, result.Appointment().Enabled)
	assert.Equal(t, "new name", result.Appointment().Name)
	assert.Equal(t, "new description", result.Appointment().Description)
	assert.Equal(t, mystatus.Done, result.Appointment().Status.Value)

}
