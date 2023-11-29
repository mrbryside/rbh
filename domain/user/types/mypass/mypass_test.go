//go:build unit

package mypass_test

import (
	"testing"

	"github.com/mrbryside/rbh/domain/user/types/mypass"
	"github.com/stretchr/testify/assert"
)

func TestVerifyPassword(t *testing.T) {
	// Arrange
	mpass, err := mypass.NewType("12345678")
	if err != nil {
		t.Error(err)
	}

	// Act && Assert
	assert.True(t, mpass.ComparePassword("12345678"))
}
