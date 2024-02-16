package userhandler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h Handler) loginOriRegister(ctx echo.Context) error {

	//panic("implement")
	return ctx.JSON(http.StatusOK, map[string]interface{}{"msg": "ok"})
}
