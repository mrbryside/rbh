package comment

import (
	"fmt"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/mrbryside/rbh/domain/interview/appointment/service"
	"github.com/mrbryside/rbh/pkg/claim"
	"github.com/mrbryside/rbh/pkg/mhttp"
)

const (
	paramId = "id"
)

func (h Handler) UpdateById(c echo.Context) error {
	var up UpdatePayload
	err := c.Bind(&up)
	if err != nil {
		return mhttp.BadRequest(c, err.Error())
	}

	id := c.Param(paramId)
	parsedId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return mhttp.BadRequest(c, "invalid id")
	}

	validate := validator.New()
	if err := validate.Struct(up); err != nil {
		return mhttp.BadRequest(c, fmt.Sprintf("validation error: %s", err.Error()))
	}

	creatorId := c.Get(claim.UserId).(uint)
	agg, err := h.commentService.UpdateById(service.UpdateCommentDto{
		Id:        uint(parsedId),
		CreatorId: creatorId,
		Message:   up.Message,
	})
	if err != nil {
		return mhttp.InternalError(c, err.Error())
	}

	return mhttp.SuccessWithBody(c, toCommentResponse(agg))
}
