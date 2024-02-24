package userhandler

import (
	"github.com/labstack/echo/v4"
	userparam "github.com/tonet-me/tonet-core/param/user"
	errmsg "github.com/tonet-me/tonet-core/pkg/err_msg"
	httpmsg "github.com/tonet-me/tonet-core/pkg/http_msg"
	"net/http"
)

func (h Handler) loginOriRegister(ctx echo.Context) error {
	req := userparam.LoginOrRegisterRequest{}
	if bErr := ctx.Bind(&req); bErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errmsg.ErrorMsgInvalidJson)
	}

	if fieldErrors, err := h.userVld.LoginRegisterRequest(req); err != nil {
		msg, code := httpmsg.Error(err)

		return ctx.JSON(code, echo.Map{
			"message": msg,
			"errors":  fieldErrors,
		})
	}

	res, sErr := h.userSvc.LoginOrRegister(ctx.Request().Context(), req)
	if sErr != nil {
		msg, code := httpmsg.Error(sErr)

		return echo.NewHTTPError(code, msg)
	}

	return ctx.JSON(http.StatusOK, res)
}
