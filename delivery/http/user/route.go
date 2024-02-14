package userhandler

import "github.com/labstack/echo/v4"

func (h Handler) SetRoutes(e *echo.Echo) {
	r := e.Group("/users")

	r.POST("/login-or-register", h.LoginOriRegister)
}
