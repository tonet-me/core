package userhandler

import (
	"fmt"
	"github.com/labstack/echo/v4"
)

func (h Handler) SetRoutes(e *echo.Echo) {
	r := e.Group("/users")
	fmt.Println("ok set route")
	r.POST("/login-or-register", h.loginOriRegister)
}
