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
	getByIdSuccessBody = `{
		"code":2000,
		"message":"success",
		"data":{
			"id":1,
			"name":"test",
			"description":"test",
			"status":"To Do",
			"enabled":true,
			"creator":{"id":1,"name":"Bryan","email":"bryan@mail.com"},
			"comments":[],
			"histories":[],
			"created_at":"2013"
		}
	}`
	getByIdParseErrorBody = `{
		"code":4000,
		"message":"bad request: strconv.ParseUint: parsing \"wrong\": invalid syntax"
	}`
	getByIdInternalErrorBody = `{
		"code":5000, 
		"message":"internal server error: internal error"
	}`
)

func TestGetAppointmentById(t *testing.T) {
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
			desc: "get by id success - should return 200 with response",
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
				as.EXPECT().GetById(gomock.Any()).Return(
					appAgg,
					nil,
				)
				return as
			},
			payload:    "",
			creatorId:  1,
			expectCode: http.StatusOK,
			expectBody: getByIdSuccessBody,
			echoContext: func(payload string, creatorId uint, method string) (*echo.Echo, echo.Context, *httptest.ResponseRecorder) {
				e := echo.New()
				req := httptest.NewRequest(method, "/", strings.NewReader(payload))
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				rec := httptest.NewRecorder()
				ctx := e.NewContext(req, rec)
				ctx.Set(claim.UserId, creatorId)
				ctx.SetParamNames("id")
				ctx.SetParamValues("1")
				return e, ctx, rec
			},
		},
		{
			desc: "get by id error parse params id - should return 400",
			asMock: func(t *testing.T) *appointmentmock.MockAppointmentServicer {
				ctrl := gomock.NewController(t)
				as := appointmentmock.NewMockAppointmentServicer(ctrl)
				return as
			},
			payload:    "",
			creatorId:  1,
			expectCode: http.StatusBadRequest,
			expectBody: getByIdParseErrorBody,
			echoContext: func(payload string, creatorId uint, method string) (*echo.Echo, echo.Context, *httptest.ResponseRecorder) {
				e := echo.New()
				req := httptest.NewRequest(method, "/", strings.NewReader(payload))
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				rec := httptest.NewRecorder()
				ctx := e.NewContext(req, rec)
				ctx.Set(claim.UserId, creatorId)
				ctx.SetParamNames("id")
				ctx.SetParamValues("wrong")
				return e, ctx, rec
			},
		},
		{
			desc: "get by id internal error - should return 500",
			asMock: func(t *testing.T) *appointmentmock.MockAppointmentServicer {
				ctrl := gomock.NewController(t)
				as := appointmentmock.NewMockAppointmentServicer(ctrl)
				as.EXPECT().GetById(gomock.Any()).Return(
					app.Aggregate{},
					errors.New("internal error"),
				)
				return as
			},
			payload:    "",
			creatorId:  1,
			expectCode: http.StatusInternalServerError,
			expectBody: getByIdInternalErrorBody,
			echoContext: func(payload string, creatorId uint, method string) (*echo.Echo, echo.Context, *httptest.ResponseRecorder) {
				e := echo.New()
				req := httptest.NewRequest(method, "/", strings.NewReader(payload))
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				rec := httptest.NewRecorder()
				ctx := e.NewContext(req, rec)
				ctx.Set(claim.UserId, creatorId)
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
			err := h.GetById(ctx)

			gotCode := rec.Code
			gotBody := rec.Body.String()

			assert.NoError(t, err)
			assert.Equal(t, tC.expectCode, gotCode)
			assert.JSONEq(t, tC.expectBody, gotBody)
		})
	}
}
