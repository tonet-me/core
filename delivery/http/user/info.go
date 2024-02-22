package userhandler

import (
	"github.com/labstack/echo/v4"
	userparam "github.com/tonet-me/tonet-core/param/user"
	"github.com/tonet-me/tonet-core/pkg/claim"
	httpmsg "github.com/tonet-me/tonet-core/pkg/http_msg"
	"net/http"
)

func (h Handler) getUserInfo(ctx echo.Context) error {
	req := userparam.GetInfoRequest{}
	if bErr := ctx.Bind(&req); bErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	claims := claim.GetClaimsFromEchoContext(ctx)
	req.AuthenticatedUserID = claims.UserID

	res, sErr := h.userSvc.GetInfo(ctx.Request().Context(), req)
	if sErr != nil {
		msg, code := httpmsg.Error(sErr)

		return echo.NewHTTPError(code, msg)
	}

	return ctx.JSON(http.StatusOK, res)
}
