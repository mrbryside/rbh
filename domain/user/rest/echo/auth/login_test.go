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
	// login_success
	loginCorrectPayload = `{
		"email": "test@mail.com",
		"password": "password"
	}`
	loginSuccessResponse = `{
		"code": 2000,
		"message": "success",
		"data": {
			"access_token": "mock_token",
			"refresh_token": "coming soon"
		}
	}`
	// login_bind_error
	loginBindErrorPayload = `{
		"email": "test@mail.com",
		"password": 1234,
	}`
	loginBindErrorResponse = `{
		"code": 4000,
		"message": "bad request: error binding payload: code=400, message=Unmarshal type error: expected=string, got=number, field=password, offset=56, internal=json: cannot unmarshal number into Go struct field loginPayload.password of type string"
	}`
)

func TestLoginSuccess(t *testing.T) {
	// Arrange
	e := echo.New()
	defer e.Close()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(loginCorrectPayload))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	ctrl := gomock.NewController(t)
	mock := mockgen.NewMockAuthServicer(ctrl)
	mock.EXPECT().Login(gomock.Any(), gomock.Any()).Return("mock_token", nil)

	wantBody := loginSuccessResponse
	wantCode := http.StatusOK

	// Act
	underTest := auth.NewAuthHandler(mock)
	err := underTest.LoginHandler(ctx)
	gotBody := rec.Body.String()
	gotCode := rec.Code

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, wantCode, gotCode)
	assert.JSONEq(t, wantBody, gotBody)

}

func TestLoginBindError(t *testing.T) {
	// Arrange
	e := echo.New()
	defer e.Close()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(loginBindErrorPayload))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	ctrl := gomock.NewController(t)
	mock := mockgen.NewMockAuthServicer(ctrl)

	wantCode := http.StatusBadRequest

	// Act
	underTest := auth.NewAuthHandler(mock)
	err := underTest.LoginHandler(ctx)
	gotCode := rec.Code

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, wantCode, gotCode)
}

func TestLoginServiceError(t *testing.T) {
	// Arrange
	e := echo.New()
	defer e.Close()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(loginCorrectPayload))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	ctrl := gomock.NewController(t)
	mock := mockgen.NewMockAuthServicer(ctrl)
	mock.EXPECT().Login(gomock.Any(), gomock.Any()).Return("", errors.New("login service internal error"))

	wantCode := http.StatusInternalServerError

	// Act
	underTest := auth.NewAuthHandler(mock)
	err := underTest.LoginHandler(ctx)
	gotCode := rec.Code

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, wantCode, gotCode)
}
