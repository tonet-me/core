package cardhandler

import (
	"github.com/labstack/echo/v4"
	cardparam "github.com/tonet-me/tonet-core/param/card"
	"github.com/tonet-me/tonet-core/pkg/claim"
	errmsg "github.com/tonet-me/tonet-core/pkg/err_msg"
	httpmsg "github.com/tonet-me/tonet-core/pkg/http_msg"
	"net/http"
)

func (h Handler) updateCard(ctx echo.Context) error {
	req := cardparam.UpdateRequest{}
	bErr := ctx.Bind(&req)
	if bErr != nil {
		return ctx.JSON(http.StatusBadRequest, errmsg.ErrorMsgInvalidJson)
	}

	claims := claim.GetClaimsFromEchoContext(ctx)
	req.AuthenticatedUserID = claims.UserID

	res, uErr := h.cardSvc.Update(ctx.Request().Context(), req)
	if uErr != nil {
		msg, code := httpmsg.Error(uErr)

		return echo.NewHTTPError(code, msg)
	}

	return ctx.JSON(http.StatusOK, res)
}
