//go:build integration

package gormhistory_test

import (
	"testing"

	"github.com/mrbryside/rbh/domain/interview/appointment/domain/appointment"
	"github.com/mrbryside/rbh/domain/interview/appointment/domain/appointment/gormappointment"
	"github.com/mrbryside/rbh/domain/interview/appointment/domain/history"
	"github.com/mrbryside/rbh/domain/interview/appointment/domain/history/gormhistory"
	"github.com/mrbryside/rbh/pkg/db/mygorm"
	"github.com/stretchr/testify/assert"
)

func createHistory() uint {
	db, close, _ := mygorm.BasicConnection()
	defer close()
	db.Exec("DELETE FROM histories")

	repoApp := gormappointment.NewRepository(db)
	aggApp := appointment.New("name", "description")
	aggWithCreator := aggApp.SetCreatorId(1)

	resultApp, _ := repoApp.Create(aggWithCreator)
	underTest := gormhistory.NewRepository(db)
	aggHis, _ := history.New(resultApp.Appointment().Id, "new name", "new description", "Done")

	underTest.Create(aggHis)
	return resultApp.Appointment().Id

}
func TestCreateHistory(t *testing.T) {
	// Arrange
	db, close, err := mygorm.BasicConnection()
	if err != nil {
		t.Error(err)
	}
	defer close()
	db.Exec("DELETE FROM histories")

	aggApp := appointment.New("name", "description")
	aggWithCreator := aggApp.SetCreatorId(1)

	repoApp := gormappointment.NewRepository(db)
	resultApp, err := repoApp.Create(aggWithCreator)
	if err != nil {
		t.Error(err.Error())
	}

	underTest := gormhistory.NewRepository(db)
	aggHis, err := history.New(resultApp.Appointment().Id, "new name", "new description", "Done")
	if err != nil {
		t.Error(err.Error())
	}

	// Act
	hisAggResp, err := underTest.Create(aggHis)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, hisAggResp.Appointment().Name, "new name")
	assert.Equal(t, hisAggResp.Appointment().Description, "new description")
	assert.Equal(t, hisAggResp.Appointment().Status.Value, "Done")
}
