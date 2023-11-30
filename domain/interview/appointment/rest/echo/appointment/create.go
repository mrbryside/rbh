package appointment

import (
	"github.com/labstack/echo/v4"
	"github.com/mrbryside/rbh/domain/interview/appointment/service"
	"github.com/mrbryside/rbh/pkg/claim"
	"github.com/mrbryside/rbh/pkg/mhttp"
)

func (h Handler) Create(c echo.Context) error {
	var ap CreatePayload
	err := c.Bind(&ap)
	if err != nil {
		return mhttp.BadRequest(c, err.Error())
	}
	creatorId := c.Get(claim.UserId).(uint)
	result, err := h.appointmentService.Create(service.CreateAppointmentDto{
		Name:        ap.Name,
		Description: ap.Description,
		CreatorId:   creatorId,
	})
	if err != nil {
		return mhttp.InternalError(c, err.Error())
	}
	return mhttp.SuccessCreatedWithBody(c, toAppointmentResp(result))
}
