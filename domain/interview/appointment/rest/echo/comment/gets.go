package comment

import (
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/mrbryside/rbh/pkg/mhttp"
)

func (h Handler) GetAll(c echo.Context) error {
	id := c.Param("id")
	parsedId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return mhttp.BadRequest(c, "invalid id")
	}

	result, err := h.commentService.GetAllByAppointmentId(uint(parsedId))
	if err != nil {
		return mhttp.InternalError(c, err.Error())
	}
	return mhttp.SuccessWithBody(c, toCommentGetAllResponse(result))
}
