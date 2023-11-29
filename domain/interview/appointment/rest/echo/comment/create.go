package comment

import (
	"fmt"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/mrbryside/rbh/domain/interview/appointment/service"
	"github.com/mrbryside/rbh/pkg/claim"
	"github.com/mrbryside/rbh/pkg/mhttp"
)

func (h Handler) Create(c echo.Context) error {
	var cp CreatePayload
	err := c.Bind(&cp)
	if err != nil {
		return mhttp.BadRequest(c, err.Error())
	}

	validate := validator.New()
	if err := validate.Struct(cp); err != nil {
		return mhttp.BadRequest(c, fmt.Sprintf("validation error: %s", err.Error()))
	}

	creatorId := c.Get(claim.UserId).(uint)
	agg, err := h.commentService.Create(service.CreateCommentDto{
		Message:       cp.Message,
		AppointmentId: cp.AppointmentId,
		CreatorId:     creatorId,
	})
	if err != nil {
		return mhttp.InternalError(c, err.Error())
	}

	return mhttp.SuccessWithBody(c, toCommentResponse(agg))
}
