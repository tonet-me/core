package httpserver

import "github.com/labstack/echo/v4"

type Handler interface {
	SetRoutes(e *echo.Echo)
}
