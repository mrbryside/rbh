//go:build unit

package comment_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	com "github.com/mrbryside/rbh/domain/interview/appointment/domain/comment"
	"github.com/mrbryside/rbh/domain/interview/appointment/domain/creator"
	"github.com/mrbryside/rbh/domain/interview/appointment/pkg/generated/commentmock"
	"github.com/mrbryside/rbh/domain/interview/appointment/rest/echo/comment"
	"github.com/mrbryside/rbh/pkg/claim"
	"github.com/stretchr/testify/assert"
)

const (
	commentCreateSuccessPayload = `{
		"message": "test message",
		"appointment_id": 1
	}`
	commentCreateBindErrorPayload = `{
		"message": "test message",
		"appointment_id": "invalid"
	}`
	commentCreateValidateErrorPayload = `{
		"message": "test message"
	}`

	commentCreateSuccessBody = `{
		"code":2001,
		"message":"created",
		"data":{
			"id":1,
			"message":"test message",
			"creator":{
				"id":1,
				"name":"Bryan",
				"email":"bryan@mail.com"
			}
		}
	}`
	commentCreateInternalErrorBody = `{
		"code":5000,
		"message":"internal server error: internal error"
	}`
	commentCreateBindErrorBody = `{
		"code":4000,
		"message":"bad request: code=400, message=Unmarshal type error: expected=uint, got=string, field=appointment_id, offset=60, internal=json: cannot unmarshal string into Go struct field CreatePayload.appointment_id of type uint"
	}`
	commentCreateValidateErrorBody = `{
		"code":4000,
		"message":"bad request: validation error: Key: 'CreatePayload.AppointmentId' Error:Field validation for 'AppointmentId' failed on the 'required' tag"
	}`
)

func TestCreateComment(t *testing.T) {
	testCases := []struct {
		desc        string
		csMock      func(t *testing.T) *commentmock.MockCommentServicer
		echoContext func(string, uint, string) (*echo.Echo, echo.Context, *httptest.ResponseRecorder)
		creatorId   uint
		payload     string
		expectCode  int
		expectBody  string
	}{
		{
			desc: "create success - should return 201 with response",
			csMock: func(t *testing.T) *commentmock.MockCommentServicer {
				ctrl := gomock.NewController(t)
				as := commentmock.NewMockCommentServicer(ctrl)
				c, err := creator.New(1, "Bryan", "bryan@mail.com")
				if err != nil {
					t.Error(err.Error())
				}
				comAgg := com.New("test message").
					SetAppointmentId(1).
					SetCreator(c.Creator()).
					SetCommentId(1)
				as.EXPECT().Create(gomock.Any()).Return(
					comAgg,
					nil,
				)
				return as
			},
			payload:    commentCreateSuccessPayload,
			creatorId:  1,
			expectCode: http.StatusCreated,
			expectBody: commentCreateSuccessBody,
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
			csMock: func(t *testing.T) *commentmock.MockCommentServicer {
				ctrl := gomock.NewController(t)
				as := commentmock.NewMockCommentServicer(ctrl)
				as.EXPECT().Create(gomock.Any()).Return(
					com.Aggregate{},
					errors.New("internal error"),
				)
				return as
			},
			payload:    commentCreateSuccessPayload,
			creatorId:  1,
			expectCode: http.StatusInternalServerError,
			expectBody: commentCreateInternalErrorBody,
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
			desc: "create bind error - should return 400",
			csMock: func(t *testing.T) *commentmock.MockCommentServicer {
				ctrl := gomock.NewController(t)
				as := commentmock.NewMockCommentServicer(ctrl)
				return as
			},
			payload:    commentCreateBindErrorPayload,
			creatorId:  1,
			expectCode: http.StatusBadRequest,
			expectBody: commentCreateBindErrorBody,
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
			desc: "create validation error - should return 400",
			csMock: func(t *testing.T) *commentmock.MockCommentServicer {
				ctrl := gomock.NewController(t)
				as := commentmock.NewMockCommentServicer(ctrl)
				return as
			},
			payload:    commentCreateValidateErrorPayload,
			creatorId:  1,
			expectCode: http.StatusBadRequest,
			expectBody: commentCreateValidateErrorBody,
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
			h := comment.NewHandler(tC.csMock(t))
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
