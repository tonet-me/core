package cardhandler

import (
	"github.com/labstack/echo/v4"
	cardparam "github.com/tonet-me/tonet-core/param/card"
	"github.com/tonet-me/tonet-core/pkg/claim"
	httpmsg "github.com/tonet-me/tonet-core/pkg/http_msg"
	"net/http"
)

func (h Handler) deActiveCardByID(ctx echo.Context) error {
	req := cardparam.DeActiveRequest{}

	cardIdFromClient := ctx.Param("id")
	req.CardID = cardIdFromClient

	claims := claim.GetClaimsFromEchoContext(ctx)
	req.AuthenticatedUserID = claims.UserID

	if fieldErrors, err := h.cardVld.CheckIfCardDeleted(cardparam.DeleteRequest{
		CardID: req.CardID,
	}); err != nil {
		msg, code := httpmsg.Error(err)

		return echo.NewHTTPError(code, echo.Map{
			"message": msg,
			"errors":  fieldErrors,
		})
	}

	res, gErr := h.cardSvc.DeActive(ctx.Request().Context(), req)
	if gErr != nil {
		msg, code := httpmsg.Error(gErr)

		return echo.NewHTTPError(code, msg)
	}

	return ctx.JSON(http.StatusOK, res)
}
