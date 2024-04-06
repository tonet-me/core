package cardhandler

import (
	"github.com/labstack/echo/v4"
	cardparam "github.com/tonet-me/tonet-core/param/card"
	"github.com/tonet-me/tonet-core/pkg/claim"
	httpmsg "github.com/tonet-me/tonet-core/pkg/http_msg"
	"net/http"
)

func (h Handler) getAllUserCards(ctx echo.Context) error {
	req := cardparam.GetAllUserCardsRequest{}

	claims := claim.GetClaimsFromEchoContext(ctx)
	req.AuthenticatedUserID = claims.UserID

	res, gErr := h.cardSvc.GetAllUserCards(ctx.Request().Context(), req)
	if gErr != nil {
		msg, code := httpmsg.Error(gErr)

		return echo.NewHTTPError(code, msg)
	}

	return ctx.JSON(http.StatusOK, res)
}
