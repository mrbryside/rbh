package appointment

import (
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/mrbryside/rbh/pkg/mhttp"
)

func (h Handler) GetById(c echo.Context) error {
	id := c.Param("id")
	parsedId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return mhttp.BadRequest(c, err.Error())
	}

	result, err := h.appointmentService.GetById(uint(parsedId))
	if err != nil {
		return mhttp.InternalError(c, err.Error())
	}
	return mhttp.SuccessWithBody(c, toAppointmentResp(result))
}
