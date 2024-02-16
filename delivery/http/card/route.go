package cardhandler

import (
	"github.com/labstack/echo/v4"
)

func (h Handler) SetRoutes(e *echo.Echo) {
	r := e.Group("/cards")
	r.POST("/", h.createNewCard)
}
