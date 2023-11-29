package comment

import (
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/mrbryside/rbh/pkg/claim"
	"github.com/mrbryside/rbh/pkg/mhttp"
)

func (h Handler) DeleteById(c echo.Context) error {
	id := c.Param(paramId)
	parsedId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return mhttp.BadRequest(c, "invalid id")
	}

	creatorId := c.Get(claim.UserId).(uint)
	err = h.commentService.DeleteById(uint(parsedId), creatorId)
	if err != nil {
		return mhttp.InternalError(c, err.Error())
	}

	return mhttp.SuccessWithMessage(c, "deleted comment")
}
