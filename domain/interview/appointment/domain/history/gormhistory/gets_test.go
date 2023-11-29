//go:build integration

package gormhistory_test

import (
	"testing"

	"github.com/mrbryside/rbh/domain/interview/appointment/domain/history/gormhistory"
	"github.com/mrbryside/rbh/pkg/db/mygorm"
	"github.com/stretchr/testify/assert"
)

func TestGetAllHistory(t *testing.T) {
	// Arrange
	db, close, err := mygorm.BasicConnection()
	if err != nil {
		t.Error(err)
	}
	defer close()
	db.Exec("DELETE FROM histories")
	appId := createHistory()

	underTest := gormhistory.NewRepository(db)
	// Act
	responses, err := underTest.GetAllByAppointmentId(appId)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, 1, len(responses))
}
