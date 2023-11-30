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
	"github.com/mrbryside/rbh/pkg/claim"
	"github.com/stretchr/testify/assert"
)

const (
	createSuccessPayload = `{
		"name": "test",
		"description": "test"
	}`
	createBindErrorPayload = `{
		"name": "test",
		"description": 1234 
	}`

	createSuccessBody = `{
		"code": 2001,
		"message": "created",
		"data": {
			"id": 1,
			"name": "test",
			"description": "test",
			"status": "To Do",
			"enabled": true,
			"creator": {
				"id": 1,
				"name": "Bryan",
				"email": "bryan@mail.com"
			},
			"comments": [],
			"histories": [],
			"created_at": "2013"
		}
	}`
	createBindErrorBody = `{
		"code":4000,
		"message":"bad request: code=400, message=Unmarshal type error: expected=string, got=number, field=description, offset=41, internal=json: cannot unmarshal number into Go struct field CreatePayload.description of type string"
		 
	}`
	createInternalErrorBody = `{
		"code":5000,
		"message":"internal server error: internal error"
	}`
)

func TestCreateAppointment(t *testing.T) {
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
			desc: "create success - should return 201 with response",
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
				as.EXPECT().Create(gomock.Any()).Return(
					appAgg,
					nil,
				)
				return as
			},
			payload:    createSuccessPayload,
			creatorId:  1,
			expectCode: http.StatusCreated,
			expectBody: createSuccessBody,
			echoContext: func(payload string, creatorId uint, method string) (*echo.Echo, echo.Context, *httptest.ResponseRecorder) {
				e := echo.New()
				req := httptest.NewRequest(method, "/", strings.NewReader(payload))
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				rec := httptest.NewRecorder()
				ctx := e.NewContext(req, rec)
				ctx.Set(claim.UserId, creatorId)
				return e, ctx, rec
			},
		},
		{
			desc: "create bad request with wrong payload - should return 400",
			asMock: func(t *testing.T) *appointmentmock.MockAppointmentServicer {
				ctrl := gomock.NewController(t)
				as := appointmentmock.NewMockAppointmentServicer(ctrl)
				return as
			},
			payload:    createBindErrorPayload,
			creatorId:  1,
			expectCode: http.StatusBadRequest,
			expectBody: createBindErrorBody,
			echoContext: func(payload string, creatorId uint, method string) (*echo.Echo, echo.Context, *httptest.ResponseRecorder) {
				e := echo.New()
				req := httptest.NewRequest(method, "/", strings.NewReader(payload))
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				rec := httptest.NewRecorder()
				ctx := e.NewContext(req, rec)
				ctx.Set(claim.UserId, creatorId)
				return e, ctx, rec
			},
		},
		{
			desc: "create internal error - should return 500",
			asMock: func(t *testing.T) *appointmentmock.MockAppointmentServicer {
				ctrl := gomock.NewController(t)
				as := appointmentmock.NewMockAppointmentServicer(ctrl)
				as.EXPECT().Create(gomock.Any()).Return(app.Aggregate{}, errors.New("internal error"))
				return as
			},
			payload:    createSuccessBody,
			creatorId:  1,
			expectCode: http.StatusInternalServerError,
			expectBody: createInternalErrorBody,
			echoContext: func(payload string, creatorId uint, method string) (*echo.Echo, echo.Context, *httptest.ResponseRecorder) {
				e := echo.New()
				req := httptest.NewRequest(method, "/", strings.NewReader(payload))
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				rec := httptest.NewRecorder()
				ctx := e.NewContext(req, rec)
				ctx.Set(claim.UserId, creatorId)
				return e, ctx, rec
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			h := appointment.NewHandler(tC.asMock(t))
			e, ctx, rec := tC.echoContext(tC.payload, tC.creatorId, http.MethodPost)
			defer e.Close()
			err := h.Create(ctx)

			gotCode := rec.Code
			gotBody := rec.Body.String()

			assert.NoError(t, err)
			assert.Equal(t, tC.expectCode, gotCode)
			assert.JSONEq(t, tC.expectBody, gotBody)
		})
	}
}
