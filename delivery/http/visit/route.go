package visithandler

import (
	"github.com/labstack/echo/v4"
)

func (h Handler) SetRoutes(e *echo.Echo) {
	r := e.Group("/v1/visits")

	r.GET("/:card-name", h.visit)
}
