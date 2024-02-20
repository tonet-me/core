package miniohandler

import (
	"github.com/labstack/echo/v4"
)

func (h Handler) SetRoutes(e *echo.Echo) {
	r := e.Group("/files")
	r.POST("/profile", h.uploadUserProfile)
	r.GET("/profile/:id", h.downloadUserProfile)
}
