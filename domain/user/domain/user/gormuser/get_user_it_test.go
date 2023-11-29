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

func TestGetUserById(t *testing.T) {
	// Arrange
	db, close, err := mygorm.BasicConnection()
	if err != nil {
		t.Error(err)
	}
	defer close()
	db.Exec("DELETE FROM users WHERE id > 1")

	repo := gormuser.NewRepository(db)
	agg, err := user.New("Bryan", "sirawat.get@ku.th", "12345678")
	if err != nil {
		t.Error(err)
	}
	aggWithRole, err := agg.SetRole(myrole.PeopleTeam)
	if err != nil {
		t.Error(err)
	}
	userAggResp, err := repo.Create(aggWithRole)
	if err != nil {
		t.Error(err)
	}
	wantName := userAggResp.Person().Name
	wantEmail := userAggResp.Person().Email.Value
	wantId := userAggResp.Person().Id
	wantRole := userAggResp.Person().Role.Value

	// Act
	uResp, err := repo.GetById(userAggResp.Person().Id)
	if err != nil {
		t.Error(err)
	}
	gotName := uResp.Person().Name
	gotEmail := uResp.Person().Email.Value
	gotId := uResp.Person().Id
	gotRole := uResp.Person().Role.Value

	// Assert
	assert.Equal(t, wantId, gotId)
	assert.Equal(t, wantName, gotName)
	assert.Equal(t, wantEmail, gotEmail)
	assert.Equal(t, wantRole, gotRole)
}

func TestGetUserByEmail(t *testing.T) {
	// Arrange
	db, close, err := mygorm.BasicConnection()
	if err != nil {
		t.Error(err)
	}
	defer close()
	db.Exec("DELETE FROM users WHERE id > 1")

	repo := gormuser.NewRepository(db)
	agg, err := user.New("Bryan", "sirawat.getEmail@ku.th", "12345678")
	if err != nil {
		t.Error(err)
	}
	aggWithRole, err := agg.SetRole(myrole.PeopleTeam)
	if err != nil {
		t.Error(err)
	}
	userAggResp, err := repo.Create(aggWithRole)
	if err != nil {
		t.Error(err)
	}
	wantName := userAggResp.Person().Name
	wantEmail := userAggResp.Person().Email.Value
	wantId := userAggResp.Person().Id
	wantRole := userAggResp.Person().Role.Value

	// Act
	email := userAggResp.Person().Email.Value
	uResp, err := repo.GetByEmail(email)
	if err != nil {
		t.Error(err)
	}
	gotName := uResp.Person().Name
	gotEmail := uResp.Person().Email.Value
	gotId := uResp.Person().Id
	gotRole := uResp.Person().Role.Value

	// Assert
	assert.Equal(t, wantId, gotId)
	assert.Equal(t, wantName, gotName)
	assert.Equal(t, wantEmail, gotEmail)
	assert.Equal(t, wantRole, gotRole)
}
