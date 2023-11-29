//go:build integration

package gormcomment_test

import (
	"testing"

	"github.com/mrbryside/rbh/domain/interview/appointment/domain/comment/gormcomment"
	"github.com/mrbryside/rbh/pkg/db/mygorm"
	"github.com/stretchr/testify/assert"
)

func TestDeleteCommentSuccess(t *testing.T) {
	// Arrange
	commentId, _ := createComment()
	db, close, err := mygorm.BasicConnection()
	if err != nil {
		t.Error(err)
	}
	defer close()

	repo := gormcomment.NewRepository(db)

	// Act
	err = repo.DeleteById(commentId)

	// Assert
	assert.Nil(t, err)
}
