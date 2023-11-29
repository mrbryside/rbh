package appointment

import (
	"github.com/labstack/echo/v4"
	"github.com/mrbryside/rbh/domain/interview/appointment/service"
	"github.com/mrbryside/rbh/pkg/mhttp"
)

var (
	defaultPage     = uint(1)
	defaultPageSize = uint(10)
)

type getAllQueryParams struct {
	Page     uint `json:"page" query:"page"`
	PageSize uint `json:"page_size" query:"page_size"`
}

func newGetAllQueryParams() getAllQueryParams {
	return getAllQueryParams{
		Page:     defaultPage,
		PageSize: defaultPageSize,
	}
}

func (ah Handler) GetAll(c echo.Context) error {
	qp := newGetAllQueryParams()
	c.Bind(&qp)

	results, err := ah.appointmentService.GetAll(service.GetAllAppointmentDto{
		Page:     qp.Page,
		PageSize: qp.PageSize,
	})
	if err != nil {
		return mhttp.InternalError(c, err.Error())
	}
	return mhttp.SuccessWithBody(c, toAppointmentPaginateResp(results))
}
