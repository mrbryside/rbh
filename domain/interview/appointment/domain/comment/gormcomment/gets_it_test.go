//go:build integration

package gormcomment_test

import (
	"testing"

	"github.com/mrbryside/rbh/domain/interview/appointment/domain/comment/gormcomment"
	"github.com/mrbryside/rbh/pkg/db/mygorm"
	"github.com/stretchr/testify/assert"
)

func TestGetAllSuccess(t *testing.T) {
	// Arrange
	db, close, err := mygorm.BasicConnection()
	if err != nil {
		t.Error(err)
	}
	db.Exec("DELETE FROM appointments")
	db.Exec("DELETE FROM comments")
	defer close()
	_, appId := createComment()
	createComment()

	// Act
	underTest := gormcomment.NewRepository(db)
	result, err := underTest.GetAllByAppointmentId(appId)

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, 1, len(result))

}
