package comment

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/mrbryside/rbh/domain/interview/appointment/service"
)

const (
	basePath = "/comments"
)

type Handler struct {
	commentService service.CommentServicer
}

func NewHandler(cs service.CommentServicer) Handler {
	return Handler{
		commentService: cs,
	}
}

func (h Handler) RegisterRoutes(g *echo.Group) {
	g.POST(basePath, h.Create)
	g.GET(fmt.Sprintf("%s/appointment/:id", basePath), h.GetAll)
	g.PUT(fmt.Sprintf("%s/:id", basePath), h.UpdateById)
	g.DELETE(fmt.Sprintf("%s/:id", basePath), h.DeleteById)
}
