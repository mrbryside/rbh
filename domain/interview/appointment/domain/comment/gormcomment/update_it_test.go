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

func TestUpdateCommentSuccess(t *testing.T) {
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

	commentId, _ := createComment()
	commentAgg := comment.New("new message")
	commentAggWithAppId := commentAgg.SetAppointmentId(createdApp.Appointment().Id)
	commentAggWithCommentId := commentAggWithAppId.SetCommentId(commentId)
	commentAggWithCreatorId := commentAggWithCommentId.SetCreatorId(1)

	// Act
	underTest := gormcomment.NewRepository(db)
	result, err := underTest.UpdateById(commentAggWithCreatorId)

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, commentAggWithCreatorId.Comment().Message, result.Comment().Message)
}
