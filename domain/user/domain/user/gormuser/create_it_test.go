//go:build integration

package gormuser_test

import (
	"testing"

	"github.com/mrbryside/rbh/domain/user/domain/user"
	"github.com/mrbryside/rbh/domain/user/domain/user/gormuser"
	"github.com/mrbryside/rbh/domain/user/types/myrole"
	"github.com/mrbryside/rbh/pkg/db/mygorm"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	// Arrange
	db, close, err := mygorm.BasicConnection()
	if err != nil {
		t.Error(err)
	}
	defer close()
	db.Exec("DELETE FROM users WHERE id > 1")

	repo := gormuser.NewRepository(db)
	agg, err := user.New("Bryan", "sirawat.create@ku.th", "12345678")
	if err != nil {
		t.Error(err)
	}
	aggWithRole, err := agg.SetRole(myrole.PeopleTeam)
	if err != nil {
		t.Error(err)
	}

	// Act
	userAggResp, err := repo.Create(aggWithRole)
	if err != nil {
		t.Error(err)
	}

	// Assert
	assert.Equal(t, userAggResp.Person().Name, "Bryan")
	assert.Equal(t, userAggResp.Person().Email.Value, "sirawat.create@ku.th")
	assert.True(t, userAggResp.Person().Password.ComparePassword("12345678"))
	assert.Equal(t, userAggResp.Person().Role.Value, myrole.PeopleTeam)
}
