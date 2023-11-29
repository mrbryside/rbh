//go:build integration

package gormappointment_test

import (
	"testing"

	"github.com/mrbryside/rbh/domain/interview/appointment/domain/appointment"
	"github.com/mrbryside/rbh/domain/interview/appointment/domain/appointment/gormappointment"
	"github.com/mrbryside/rbh/pkg/db/mygorm"
	"github.com/stretchr/testify/assert"
)

func createAppointment() {
	db, close, _ := mygorm.BasicConnection()
	defer close()

	repo := gormappointment.NewRepository(db)
	agg := appointment.New("name", "description")
	aggWithCreator := agg.SetCreatorId(1)

	repo.Create(aggWithCreator)
}

func TestCreateAppointment(t *testing.T) {
	// Arrange
	db, close, err := mygorm.BasicConnection()
	if err != nil {
		t.Error(err)
	}
	db.Exec("DELETE FROM appointments")
	defer close()

	repo := gormappointment.NewRepository(db)
	agg := appointment.New("name", "description")
	if err != nil {
		t.Error(err.Error())
	}
	aggWithCreator := agg.SetCreatorId(1)

	// Act
	created_agg, err := repo.Create(aggWithCreator)

	// Assert
	assert.Nil(t, err)
	assert.NotZero(t, created_agg.Appointment().Id)
	assert.Equal(t, "description", created_agg.Appointment().Description)
	assert.Equal(t, "name", created_agg.Appointment().Name)
}
