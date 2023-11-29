//go:build integration

package gormcomment_test

import (
	"testing"

	"github.com/mrbryside/rbh/domain/interview/appointment/domain/appointment"
	"github.com/mrbryside/rbh/domain/interview/appointment/domain/appointment/gormappointment"
	"github.com/mrbryside/rbh/domain/interview/appointment/domain/comment"
	"github.com/mrbryside/rbh/domain/interview/appointment/domain/comment/gormcomment"
	"github.com/mrbryside/rbh/pkg/db/mygorm"
	"github.com/stretchr/testify/assert"
)

func createComment() (uint, uint) {
	// Arrange
	db, close, _ := mygorm.BasicConnection()
	defer close()

	aggApp := appointment.New("name", "description")
	aggAppWithCreator := aggApp.SetCreatorId(1)

	repoApp := gormappointment.NewRepository(db)
	createdApp, _ := repoApp.Create(aggAppWithCreator)

	commentAgg := comment.New("message")
	commentAggWithAppId := commentAgg.SetAppointmentId(createdApp.Appointment().Id)
	commentAggWithCreatorId := commentAggWithAppId.SetCreatorId(1)

	repo := gormcomment.NewRepository(db)
	result, _ := repo.Create(commentAggWithCreatorId)
	return result.Comment().Id, result.Appointment().Id
}

func TestCreateCommentSuccess(t *testing.T) {
	// Arrange
	db, close, err := mygorm.BasicConnection()
	if err != nil {
		t.Error(err)
	}
	db.Exec("DELETE FROM appointments")
	db.Exec("DELETE FROM comments")
	defer close()

	aggApp := appointment.New("name", "description")
	if err != nil {
		t.Error(err.Error())
	}
	aggAppWithCreator := aggApp.SetCreatorId(1)

	repoApp := gormappointment.NewRepository(db)
	createdApp, err := repoApp.Create(aggAppWithCreator)
	if err != nil {
		t.Error(err.Error())
	}

	commentAgg := comment.New("message")
	commentAggWithAppId := commentAgg.SetAppointmentId(createdApp.Appointment().Id)
	commentAggWithCreatorId := commentAggWithAppId.SetCreatorId(1)

	underTest := gormcomment.NewRepository(db)
	result, err := underTest.Create(commentAggWithCreatorId)

	// Act
	assert.NoError(t, err)
	assert.Equal(t, "message", result.Comment().Message)
	assert.Equal(t, createdApp.Appointment().Id, result.Appointment().Id)
	assert.Equal(t, uint(1), result.Comment().Creator.Id)
}
