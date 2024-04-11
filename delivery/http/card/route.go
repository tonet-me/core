package cardhandler

import (
	"github.com/labstack/echo/v4"
	"github.com/tonet-me/tonet-core/delivery/http/middleware"
)

func (h Handler) SetRoutes(e *echo.Echo) {
	r := e.Group("/v1/cards")

	r.POST("", h.createNewCard, middleware.Authentication(h.authSvc, h.authConfig))
	r.PUT("", h.updateCard, middleware.Authentication(h.authSvc, h.authConfig))
	r.PUT("/active/:id", h.activeCardByID, middleware.Authentication(h.authSvc, h.authConfig))
	r.PUT("/de-active/:id", h.deActiveCardByID, middleware.Authentication(h.authSvc, h.authConfig))

	r.GET("", h.getAllUserCards, middleware.Authentication(h.authSvc, h.authConfig))
	r.GET("/:id", h.getCardInfoByID, middleware.Authentication(h.authSvc, h.authConfig))

	r.DELETE("/:id", h.deleteCardByID, middleware.Authentication(h.authSvc, h.authConfig))

	r.GET("/is-exist/:name", h.isCardExist)

}
