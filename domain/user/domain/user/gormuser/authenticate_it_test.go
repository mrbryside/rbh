//go:build integration

package gormuser

import (
	"testing"

	"github.com/mrbryside/rbh/domain/user/domain/user"
	"github.com/mrbryside/rbh/pkg/db/mygorm"
	"github.com/stretchr/testify/assert"
)

func TestAuthentication(t *testing.T) {

	// Arrange
	db, close, err := mygorm.BasicConnection()
	if err != nil {
		t.Error(err)
	}
	defer close()
	db.Exec("DELETE FROM users WHERE id > 1")

	repo := NewRepository(db)
	userAgg, err := user.New("Bryan", "sirawat.i@ku.th", "12345678")
	if err != nil {
		t.Error(err)
	}
	userAggWithRole, err := userAgg.SetRole("PeopleTeam")
	if err != nil {
		t.Error(err)
	}
	createdUserAgg, err := repo.Create(userAggWithRole)
	if err != nil {
		t.Error(err)
	}

	// Act
	email := createdUserAgg.Person().Email.Value
	gotIsAuthen, _ := repo.Authenticate(email, "12345678")

	// Assert
	assert.True(t, gotIsAuthen)
}
