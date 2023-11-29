package appointment

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/mrbryside/rbh/domain/interview/appointment/service"
)

const (
	basePath = "/appointments"
)

type Handler struct {
	appointmentService service.AppointmentServicer
}

func NewHandler(as service.AppointmentServicer) Handler {
	return Handler{
		appointmentService: as,
	}
}

func (h Handler) RegisterRoutes(g *echo.Group) {
	g.POST(basePath, h.Create)
	g.GET(basePath, h.GetAll)
	g.GET(fmt.Sprintf("%s/:id", basePath), h.GetById)
	g.PUT(fmt.Sprintf("%s/:id", basePath), h.UpdateById)
}
