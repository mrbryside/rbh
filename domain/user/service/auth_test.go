//go:build unit

package service

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mrbryside/rbh/domain/user/domain/user"
	"github.com/mrbryside/rbh/domain/user/pkg/generated/mockgen"
	"github.com/stretchr/testify/assert"
)

// -- Test Register --//
func TestRegisterSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockJwtService := mockgen.NewMockJwtServicer(ctrl)
	mockUserDomain := mockgen.NewMockRepository(ctrl)

	userAgg, err := user.New("name", "email@mail.com", "password")
	if err != nil {
		t.Error(err.Error())
	}
	userAggWithId := userAgg.SetId(1)

	mockUserDomain.EXPECT().Create(gomock.Any()).Return(userAggWithId, nil)
	mockJwtService.EXPECT().GenerateToken(gomock.Any(), gomock.Any()).Return("token", nil)
	underTest := NewAuthService(mockUserDomain, mockJwtService)
	token, err := underTest.Register("name", "email@mail.com", "password")

	assert.Nil(t, err)
	assert.Equal(t, "token", token)
}

func TestRegisterCreateError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockJwtService := mockgen.NewMockJwtServicer(ctrl)
	mockUserDomain := mockgen.NewMockRepository(ctrl)

	mockUserDomain.EXPECT().Create(gomock.Any()).Return(user.Aggregate{}, errors.New("error create user"))
	underTest := NewAuthService(mockUserDomain, mockJwtService)
	token, err := underTest.Register("name", "email@mail.com", "password")

	assert.NotNil(t, err)
	assert.Empty(t, token)
}

func TestRegisterGenerateTokenError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockJwtService := mockgen.NewMockJwtServicer(ctrl)
	mockUserDomain := mockgen.NewMockRepository(ctrl)

	userAgg, err := user.New("name", "email@mail.com", "password")
	if err != nil {
		t.Error(err.Error())
	}
	userAggWithId := userAgg.SetId(1)

	mockUserDomain.EXPECT().Create(gomock.Any()).Return(userAggWithId, nil)
	mockJwtService.EXPECT().GenerateToken(gomock.Any(), gomock.Any()).Return("", errors.New("error generate token"))
	underTest := NewAuthService(mockUserDomain, mockJwtService)
	token, err := underTest.Register("name", "email@mail.com", "password")

	assert.NotNil(t, err)
	assert.Empty(t, token)
}

func TestRegisterAggregateError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockJwtService := mockgen.NewMockJwtServicer(ctrl)
	mockUserDomain := mockgen.NewMockRepository(ctrl)

	underTest := NewAuthService(mockUserDomain, mockJwtService)
	token, err := underTest.Register("name", "emailmail.com", "password")

	assert.NotNil(t, err)
	assert.Empty(t, token)
}

// -- Test Login --//
func TestLoginSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockJwtService := mockgen.NewMockJwtServicer(ctrl)
	mockUserDomain := mockgen.NewMockRepository(ctrl)

	userAgg, err := user.New("name", "email@mail.com", "password")
	if err != nil {
		t.Error(err.Error())
	}
	userAggWithId := userAgg.SetId(1)

	mockUserDomain.EXPECT().Authenticate(gomock.Any(), gomock.Any()).Return(true, userAggWithId)
	mockJwtService.EXPECT().GenerateToken(gomock.Any(), gomock.Any()).Return("token", nil)
	underTest := NewAuthService(mockUserDomain, mockJwtService)
	token, err := underTest.Login("email@mail.com", "password")

	assert.Nil(t, err)
	assert.Equal(t, "token", token)
}

func TestLoginAuthenticationError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockJwtService := mockgen.NewMockJwtServicer(ctrl)
	mockUserDomain := mockgen.NewMockRepository(ctrl)

	mockUserDomain.EXPECT().Authenticate(gomock.Any(), gomock.Any()).Return(false, user.Aggregate{})
	underTest := NewAuthService(mockUserDomain, mockJwtService)
	token, err := underTest.Login("email@mail.com", "password")

	assert.NotNil(t, err)
	assert.Empty(t, token)

}

func TestLoginGenerateTokenFailed(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockJwtService := mockgen.NewMockJwtServicer(ctrl)
	mockUserDomain := mockgen.NewMockRepository(ctrl)

	userAgg, err := user.New("name", "email@mail.com", "password")
	if err != nil {
		t.Error(err.Error())
	}
	userAggWithId := userAgg.SetId(1)

	mockUserDomain.EXPECT().Authenticate(gomock.Any(), gomock.Any()).Return(true, userAggWithId)
	mockJwtService.EXPECT().GenerateToken(gomock.Any(), gomock.Any()).Return("", errors.New("error generate token"))
	underTest := NewAuthService(mockUserDomain, mockJwtService)
	token, err := underTest.Login("email@mail.com", "password")

	assert.NotNil(t, err)
	assert.Empty(t, token)
}

func TestGetByIdSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockJwtService := mockgen.NewMockJwtServicer(ctrl)
	mockUserDomain := mockgen.NewMockRepository(ctrl)

	userAgg, err := user.New("name", "email@mail.com", "password")
	if err != nil {
		t.Error(err.Error())
	}
	userAggWithId := userAgg.SetId(1)

	mockUserDomain.EXPECT().GetById(gomock.Any()).Return(userAggWithId, nil)
	underTest := NewAuthService(mockUserDomain, mockJwtService)
	userAggResp, err := underTest.GetById(1)

	assert.Nil(t, err)
	assert.Equal(t, userAggWithId.Person().Id, userAggResp.Person().Id)
	assert.Equal(t, userAggWithId.Person().Name, userAggResp.Person().Name)
}

func TestGetByIdError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockJwtService := mockgen.NewMockJwtServicer(ctrl)
	mockUserDomain := mockgen.NewMockRepository(ctrl)

	mockUserDomain.EXPECT().GetById(gomock.Any()).Return(user.Aggregate{}, errors.New("error get user"))
	underTest := NewAuthService(mockUserDomain, mockJwtService)
	_, err := underTest.GetById(1)

	assert.NotNil(t, err)
}
