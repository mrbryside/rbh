//go:build unit

package appointment_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	app "github.com/mrbryside/rbh/domain/interview/appointment/domain/appointment"
	"github.com/mrbryside/rbh/domain/interview/appointment/domain/creator"
	"github.com/mrbryside/rbh/domain/interview/appointment/pkg/generated/appointmentmock"
	"github.com/mrbryside/rbh/domain/interview/appointment/rest/echo/appointment"
	"github.com/stretchr/testify/assert"
)

const (
	updateSuccessPayload = `{
		"name": "edited5",
		"description": "edited5",
		"status": "In Progress",
		"enabled": true
	}`
	updateBindErrorPayload = `{
		"name": "edited5",
		"description": "edited5",
		"status": "In Progress",
		"enabled": "invalid"
	}`
	updateValidatorErrorPayload = `{
		"name": "edited5",
		"description": "edited5",
		"status": "In Progress"
	}`

	updateSuccessBody = `{
		"code":2000,
		"message":"success",
		"data":{
			"id":1,
			"name":"test",
			"description":"test",
			"status":"To Do",
			"enabled":true,
			"creator":{
				"id":1,
				"name":"Bryan",
				"email":"bryan@mail.com"
			},
			"created_at":"2013"
		}
	}`
	updateInternalErrorBody = `{
		"code":5000,
		"message":"internal server error: internal error"
	}`
	updateBindErrorBody = `{
		"code":4000,
		"message":"bad request: code=400, message=Unmarshal type error: expected=bool, got=string, field=enabled, offset=100, internal=json: cannot unmarshal string into Go struct field UpdatePayload.enabled of type bool"
	}`
	updateParamsErrorBody = `{
		"code":4000,
		"message":"bad request: invalid id"
	}`
	updateValidatorErrorBody = `{
		"code":4000,
		"message":"bad request: validation error: Key: 'UpdatePayload.Enabled' Error:Field validation for 'Enabled' failed on the 'required' tag"
	}`
)

func TestUpdateAppointment(t *testing.T) {

	testCases := []struct {
		desc        string
		asMock      func(t *testing.T) *appointmentmock.MockAppointmentServicer
		echoContext func(string, uint, string) (*echo.Echo, echo.Context, *httptest.ResponseRecorder)
		creatorId   uint
		payload     string
		expectCode  int
		expectBody  string
	}{
		{
			desc: "update success - should return 200 with response",
			asMock: func(t *testing.T) *appointmentmock.MockAppointmentServicer {
				ctrl := gomock.NewController(t)
				as := appointmentmock.NewMockAppointmentServicer(ctrl)
				c, err := creator.New(1, "Bryan", "bryan@mail.com")
				if err != nil {
					t.Error(err.Error())
				}
				appAgg := app.New("test", "test").
					SetCreator(c.Creator()).
					SetId(1).
					SetTimeStamps("2013")

				as.EXPECT().UpdateById(gomock.Any()).Return(
					appAgg,
					nil,
				)
				return as
			},
			payload:    updateSuccessPayload,
			creatorId:  1,
			expectCode: http.StatusOK,
			expectBody: updateSuccessBody,
			echoContext: func(payload string, creatorId uint, method string) (*echo.Echo, echo.Context, *httptest.ResponseRecorder) {
				e := echo.New()
				req := httptest.NewRequest(method, "/", strings.NewReader(payload))
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				rec := httptest.NewRecorder()
				ctx := e.NewContext(req, rec)
				ctx.SetParamNames("id")
				ctx.SetParamValues("1")
				return e, ctx, rec
			},
		},
		{
			desc: "update internal error - should return 500",
			asMock: func(t *testing.T) *appointmentmock.MockAppointmentServicer {
				ctrl := gomock.NewController(t)
				as := appointmentmock.NewMockAppointmentServicer(ctrl)
				as.EXPECT().UpdateById(gomock.Any()).Return(
					app.Aggregate{},
					errors.New("internal error"),
				)
				return as
			},
			payload:    updateSuccessPayload,
			creatorId:  1,
			expectCode: http.StatusInternalServerError,
			expectBody: updateInternalErrorBody,
			echoContext: func(payload string, creatorId uint, method string) (*echo.Echo, echo.Context, *httptest.ResponseRecorder) {
				e := echo.New()
				req := httptest.NewRequest(method, "/", strings.NewReader(payload))
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				rec := httptest.NewRecorder()
				ctx := e.NewContext(req, rec)
				ctx.SetParamNames("id")
				ctx.SetParamValues("1")
				return e, ctx, rec
			},
		},
		{
			desc: "update bind error - should return 400",
			asMock: func(t *testing.T) *appointmentmock.MockAppointmentServicer {
				ctrl := gomock.NewController(t)
				as := appointmentmock.NewMockAppointmentServicer(ctrl)
				return as
			},
			payload:    updateBindErrorPayload,
			creatorId:  1,
			expectCode: http.StatusBadRequest,
			expectBody: updateBindErrorBody,
			echoContext: func(payload string, creatorId uint, method string) (*echo.Echo, echo.Context, *httptest.ResponseRecorder) {
				e := echo.New()
				req := httptest.NewRequest(method, "/", strings.NewReader(payload))
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				rec := httptest.NewRecorder()
				ctx := e.NewContext(req, rec)
				ctx.SetParamNames("id")
				ctx.SetParamValues("1")
				return e, ctx, rec
			},
		},
		{
			desc: "update params error - should return 400",
			asMock: func(t *testing.T) *appointmentmock.MockAppointmentServicer {
				ctrl := gomock.NewController(t)
				as := appointmentmock.NewMockAppointmentServicer(ctrl)
				return as
			},
			payload:    updateSuccessPayload,
			creatorId:  1,
			expectCode: http.StatusBadRequest,
			expectBody: updateParamsErrorBody,
			echoContext: func(payload string, creatorId uint, method string) (*echo.Echo, echo.Context, *httptest.ResponseRecorder) {
				e := echo.New()
				req := httptest.NewRequest(method, "/", strings.NewReader(payload))
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				rec := httptest.NewRecorder()
				ctx := e.NewContext(req, rec)
				ctx.SetParamNames("id")
				ctx.SetParamValues("invalid")
				return e, ctx, rec
			},
		},
		{
			desc: "update validator body error - should return 400",
			asMock: func(t *testing.T) *appointmentmock.MockAppointmentServicer {
				ctrl := gomock.NewController(t)
				as := appointmentmock.NewMockAppointmentServicer(ctrl)
				return as
			},
			payload:    updateValidatorErrorPayload,
			creatorId:  1,
			expectCode: http.StatusBadRequest,
			expectBody: updateValidatorErrorBody,
			echoContext: func(payload string, creatorId uint, method string) (*echo.Echo, echo.Context, *httptest.ResponseRecorder) {
				e := echo.New()
				req := httptest.NewRequest(method, "/", strings.NewReader(payload))
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				rec := httptest.NewRecorder()
				ctx := e.NewContext(req, rec)
				ctx.SetParamNames("id")
				ctx.SetParamValues("1")
				return e, ctx, rec
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			h := appointment.NewHandler(tC.asMock(t))
			e, ctx, rec := tC.echoContext(tC.payload, tC.creatorId, http.MethodGet)
			defer e.Close()
			err := h.UpdateById(ctx)

			gotCode := rec.Code
			gotBody := rec.Body.String()

			assert.NoError(t, err)
			assert.Equal(t, tC.expectCode, gotCode)
			assert.JSONEq(t, tC.expectBody, gotBody)
		})
	}
}
