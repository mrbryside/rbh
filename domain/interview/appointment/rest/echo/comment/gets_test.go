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
	commentGetAllSuccessBody = `{
		"code":2000,
		"message":"success",
		"data":[
			{
				"id":1,
				"message":"test message",
				"creator":{
					"id":1,
					"name":"Bryan",
					"email":"bryan@mail.com"
				}
			}
		]
	}`
	commentGetAllInternalErrorBody = `{
		"code":5000,
		"message":"internal server error: internal error"
	}`
	commentGetAllParamsErrorBody = `{
		"code":4000,
		"message":"bad request: invalid id"
	}`
)

func TestGetAllComment(t *testing.T) {
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
			desc: "get all success - should return 200",
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
				comAggs := []com.Aggregate{comAgg}
				as.EXPECT().GetAllByAppointmentId(gomock.Any()).Return(
					comAggs,
					nil,
				)
				return as
			},
			payload:    "",
			creatorId:  1,
			expectCode: http.StatusOK,
			expectBody: commentGetAllSuccessBody,
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
			desc: "get all internal error - should return 500",
			csMock: func(t *testing.T) *commentmock.MockCommentServicer {
				ctrl := gomock.NewController(t)
				as := commentmock.NewMockCommentServicer(ctrl)
				as.EXPECT().GetAllByAppointmentId(gomock.Any()).Return(
					[]com.Aggregate{},
					errors.New("internal error"),
				)
				return as
			},
			payload:    "",
			creatorId:  1,
			expectCode: http.StatusInternalServerError,
			expectBody: commentGetAllInternalErrorBody,
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
			desc: "get all params error - should return 400",
			csMock: func(t *testing.T) *commentmock.MockCommentServicer {
				ctrl := gomock.NewController(t)
				as := commentmock.NewMockCommentServicer(ctrl)
				return as
			},
			payload:    "",
			creatorId:  1,
			expectCode: http.StatusBadRequest,
			expectBody: commentGetAllParamsErrorBody,
			echoContext: func(payload string, creatorId uint, method string) (*echo.Echo, echo.Context, *httptest.ResponseRecorder) {
				e := echo.New()
				req := httptest.NewRequest(method, "/", strings.NewReader(payload))
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				rec := httptest.NewRecorder()
				ctx := e.NewContext(req, rec)
				ctx.Set(claim.UserId, creatorId)
				ctx.SetParamNames("id")
				ctx.SetParamValues("invalid")
				return e, ctx, rec
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			h := comment.NewHandler(tC.csMock(t))
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
