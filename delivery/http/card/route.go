package cardhandler

import (
	"github.com/labstack/echo/v4"
	"github.com/tonet-me/tonet-core/delivery/http/middleware"
)

func (h Handler) SetRoutes(e *echo.Echo) {
	r := e.Group("/cards")
	r.POST("/", h.createNewCard, middleware.Authentication(h.authSvc, h.authConfig))
	r.PUT("/", h.updateCard, middleware.Authentication(h.authSvc, h.authConfig))
}
