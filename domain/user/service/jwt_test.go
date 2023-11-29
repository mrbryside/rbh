package service

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mrbryside/rbh/domain/user/domain/authorization/jwt"
	"github.com/mrbryside/rbh/domain/user/pkg/generated/mockgen"
	"github.com/stretchr/testify/assert"
)

func TestGenerateJwtTokenSuccess(t *testing.T) {
	myjwt := jwt.NewJwt("my-secret")
	service := NewJwtService(myjwt)
	token, err := service.GenerateToken(1, "People")
	if err != nil {
		t.Error(err.Error())
	}

	claim, err := myjwt.ParseToken(token)
	if err != nil {
		t.Error(err.Error())
	}

	assert.Equal(t, uint(1), claim.UserID)
	assert.Equal(t, "People", claim.Role)
}

func TestGenerateJwtTokenError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockJwt := mockgen.NewMockJwtAuthorizer(ctrl)
	mockJwt.EXPECT().GenerateToken(gomock.Any(), gomock.Any()).Return("", errors.New("error generate token"))
	service := NewJwtService(mockJwt)
	_, err := service.GenerateToken(1, "People")

	assert.NotNil(t, err)
}
