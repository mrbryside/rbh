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
	getAllSuccessBody = `{
		"code":2000,
		"message":"success",
		"data":{
			"results":[
				{
					"id":1,
					"name":"test",
					"description":"test",
					"status":"To Do",
					"creator":{
						"id":1,
						"name":"Bryan",
						"email":"bryan@mail.com"
					},
					"created_at":"2013"
				}
			],
			"next":false
		}
	}`
	getAllInternalErrorBody = `{
		"code":5000,
		"message":"internal server error: internal error"
	}`
)

func TestGetAllAppointment(t *testing.T) {

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
			desc: "get all success - should return 200 with response",
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

				appAggs := app.Aggregates{
					Aggregates: []app.Aggregate{appAgg},
					Next:       false,
				}
				as.EXPECT().GetAll(gomock.Any()).Return(
					appAggs,
					nil,
				)
				return as
			},
			payload:    "",
			creatorId:  1,
			expectCode: http.StatusOK,
			expectBody: getAllSuccessBody,
			echoContext: func(payload string, creatorId uint, method string) (*echo.Echo, echo.Context, *httptest.ResponseRecorder) {
				e := echo.New()
				req := httptest.NewRequest(method, "/?page=1&page_size=1", strings.NewReader(payload))
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				rec := httptest.NewRecorder()
				ctx := e.NewContext(req, rec)
				return e, ctx, rec
			},
		},
		{
			desc: "get all internal error - should return 500",
			asMock: func(t *testing.T) *appointmentmock.MockAppointmentServicer {
				ctrl := gomock.NewController(t)
				as := appointmentmock.NewMockAppointmentServicer(ctrl)
				as.EXPECT().GetAll(gomock.Any()).Return(
					app.Aggregates{},
					errors.New("internal error"),
				)
				return as
			},
			payload:    "",
			creatorId:  1,
			expectCode: http.StatusInternalServerError,
			expectBody: getAllInternalErrorBody,
			echoContext: func(payload string, creatorId uint, method string) (*echo.Echo, echo.Context, *httptest.ResponseRecorder) {
				e := echo.New()
				req := httptest.NewRequest(method, "/", strings.NewReader(payload))
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				rec := httptest.NewRecorder()
				ctx := e.NewContext(req, rec)
				return e, ctx, rec
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			h := appointment.NewHandler(tC.asMock(t))
			e, ctx, rec := tC.echoContext(tC.payload, tC.creatorId, http.MethodGet)
			defer e.Close()
			err := h.GetAll(ctx)

			gotCode := rec.Code
			gotBody := rec.Body.String()

			assert.NoError(t, err)
			assert.Equal(t, tC.expectCode, gotCode)
			assert.JSONEq(t, tC.expectBody, gotBody)
		})
	}
}
