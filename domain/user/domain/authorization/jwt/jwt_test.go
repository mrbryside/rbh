//go:build unit

package jwt_test

import (
	"fmt"
	"testing"

	"github.com/mrbryside/rbh/domain/user/domain/authorization/jwt"
	"github.com/stretchr/testify/assert"
)

func TestGenerateAndParseToken(t *testing.T) {
	// Arrange
	jwt := jwt.NewJwt("my-secret")
	userId := uint(1)
	role := "People"
	var err error
	token, err := jwt.GenerateToken(userId, role)
	if err != nil {
		t.Error(err.Error())
	}
	wantId := userId
	wantRole := role

	// Act
	claim, err := jwt.ParseToken(token)
	if err != nil {
		t.Error(err.Error())
	}

	gotId := claim.UserID
	gotRole := claim.Role

	// Assert
	assert.Equal(t, wantId, gotId, fmt.Sprintf("The userId should be %v", wantId))
	assert.Equal(t, wantRole, gotRole, fmt.Sprintf("The role should be %v", wantRole))
}
