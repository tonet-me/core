package miniohandler

import (
	"github.com/labstack/echo/v4"
	"github.com/tonet-me/tonet-core/delivery/http/middleware"
)

func (h Handler) SetRoutes(e *echo.Echo) {
	r := e.Group("/v1/files")

	r.POST("/profile", h.uploadUserProfile, middleware.Authentication(h.authSvc, h.authConfig))
	r.GET("/profile/:id", h.downloadUserProfile, middleware.Authentication(h.authSvc, h.authConfig))
}
