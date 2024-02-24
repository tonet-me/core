package userhandler

import (
	"github.com/labstack/echo/v4"
	userparam "github.com/tonet-me/tonet-core/param/user"
	"github.com/tonet-me/tonet-core/pkg/claim"
	errmsg "github.com/tonet-me/tonet-core/pkg/err_msg"
	httpmsg "github.com/tonet-me/tonet-core/pkg/http_msg"
	"net/http"
)

func (h Handler) updateUser(ctx echo.Context) error {
	req := userparam.UpdateRequest{}
	if bErr := ctx.Bind(&req); bErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errmsg.ErrorMsgInvalidJson)
	}

	claims := claim.GetClaimsFromEchoContext(ctx)
	req.AuthenticatedUserID = claims.UserID

	if fieldErrors, err := h.userVld.UpdateRequest(req); err != nil {
		msg, code := httpmsg.Error(err)

		return ctx.JSON(code, echo.Map{
			"message": msg,
			"errors":  fieldErrors,
		})
	}

	res, sErr := h.userSvc.Update(ctx.Request().Context(), req)
	if sErr != nil {
		msg, code := httpmsg.Error(sErr)

		return echo.NewHTTPError(code, msg)
	}

	return ctx.JSON(http.StatusOK, res)
}
