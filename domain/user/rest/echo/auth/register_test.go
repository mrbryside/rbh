//go:build unit

package auth_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/mrbryside/rbh/domain/user/pkg/generated/mockgen"
	"github.com/mrbryside/rbh/domain/user/rest/echo/auth"
	"github.com/stretchr/testify/assert"
)

const (
	// register success
	registerCorrectPayload = `{
		"name": "test",
		"email": "test@mail.com",
		"password": "password"
	}`
	registerSuccessResponse = `{
		"code": 2000,
		"message": "success",
		"data": {
			"access_token": "mock_token",
			"refresh_token": "coming soon"
		}
	}`
	// register bind error
	registerBindErrorPayload = `{
		"name": "test",
		"email": "test@mail.com",
		"password": 123456
	}`
)

func TestRegisterSuccess(t *testing.T) {
	// Arrange
	e := echo.New()
	defer e.Close()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(registerCorrectPayload))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	ctrl := gomock.NewController(t)
	mock := mockgen.NewMockAuthServicer(ctrl)
	mock.EXPECT().Register(gomock.Any(), gomock.Any(), gomock.Any()).Return("mock_token", nil)

	wantBody := registerSuccessResponse
	wantCode := http.StatusOK

	// Act
	underTest := auth.NewAuthHandler(mock)
	err := underTest.RegisterHandler(ctx)
	gotBody := rec.Body.String()
	gotCode := rec.Code

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, wantCode, gotCode)
	assert.JSONEq(t, wantBody, gotBody)

}

func TestRegisterBindError(t *testing.T) {
	// Arrange
	e := echo.New()
	defer e.Close()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(registerBindErrorPayload))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	ctrl := gomock.NewController(t)
	mock := mockgen.NewMockAuthServicer(ctrl)

	wantCode := http.StatusBadRequest

	// Act
	underTest := auth.NewAuthHandler(mock)
	err := underTest.RegisterHandler(ctx)
	gotCode := rec.Code

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, wantCode, gotCode)

}

func TestRegisterServiceError(t *testing.T) {
	// Arrange
	e := echo.New()
	defer e.Close()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(registerCorrectPayload))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	ctrl := gomock.NewController(t)
	mock := mockgen.NewMockAuthServicer(ctrl)
	mock.EXPECT().Register(gomock.Any(), gomock.Any(), gomock.Any()).Return("", errors.New("error register"))

	wantCode := http.StatusInternalServerError

	// Act
	underTest := auth.NewAuthHandler(mock)
	err := underTest.RegisterHandler(ctx)
	gotCode := rec.Code

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, wantCode, gotCode)

}
