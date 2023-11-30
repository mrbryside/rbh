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
	"github.com/mrbryside/rbh/domain/interview/appointment/pkg/generated/commentmock"
	"github.com/mrbryside/rbh/domain/interview/appointment/rest/echo/comment"
	"github.com/mrbryside/rbh/pkg/claim"
	"github.com/stretchr/testify/assert"
)

const (
	commentDeleteSuccessBody = `{
		"code":2000,
		"message":"deleted comment"
	}`

	commentParamsErrorBody = `{
		"code":4000,
		"message":"bad request: invalid id"
	}`
	commentDeleteInternalErrorBody = `{
		"code":5000,
		"message":"internal server error: internal error"
	}`
)

func TestDeleteComment(t *testing.T) {
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
			desc: "delete success - should return 200",
			csMock: func(t *testing.T) *commentmock.MockCommentServicer {
				ctrl := gomock.NewController(t)
				as := commentmock.NewMockCommentServicer(ctrl)
				as.EXPECT().DeleteById(gomock.Any(), gomock.Any()).Return(
					nil,
				)
				return as
			},
			payload:    "",
			creatorId:  1,
			expectCode: http.StatusOK,
			expectBody: commentDeleteSuccessBody,
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
			desc: "delete internal error - should return 500",
			csMock: func(t *testing.T) *commentmock.MockCommentServicer {
				ctrl := gomock.NewController(t)
				as := commentmock.NewMockCommentServicer(ctrl)
				as.EXPECT().DeleteById(gomock.Any(), gomock.Any()).Return(
					errors.New("internal error"),
				)
				return as
			},
			payload:    "",
			creatorId:  1,
			expectCode: http.StatusInternalServerError,
			expectBody: commentDeleteInternalErrorBody,
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
			desc: "delete params error - should return 400",
			csMock: func(t *testing.T) *commentmock.MockCommentServicer {
				ctrl := gomock.NewController(t)
				as := commentmock.NewMockCommentServicer(ctrl)
				return as
			},
			payload:    "",
			creatorId:  1,
			expectCode: http.StatusBadRequest,
			expectBody: commentParamsErrorBody,
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
			e, ctx, rec := tC.echoContext(tC.payload, tC.creatorId, http.MethodDelete)
			defer e.Close()
			err := h.DeleteById(ctx)

			gotCode := rec.Code
			gotBody := rec.Body.String()

			assert.NoError(t, err)
			assert.Equal(t, tC.expectCode, gotCode)
			assert.JSONEq(t, tC.expectBody, gotBody)
		})
	}
}
