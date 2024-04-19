package visithandler

import (
	"github.com/labstack/echo/v4"
	visitparam "github.com/tonet-me/tonet-core/param/visit"
	httpmsg "github.com/tonet-me/tonet-core/pkg/http_msg"
	"net/http"
	"strings"
)

func (h Handler) visit(ctx echo.Context) error {
	cardNameFromClient := ctx.Param("card-name")
	cardName := strings.Split(cardNameFromClient, `/`)[0]

	//userAgent:=ctx.Request().UserAgent() // todo: update to map[string]string from other pkg
	var req = visitparam.AddNewCardVisitRequest{
		CardName: cardName,
	}

	res, aErr := h.visitSvc.AddNewVisitToCard(ctx.Request().Context(), req) // TODO: pass user agent
	if aErr != nil {
		msg, code := httpmsg.Error(aErr)

		return echo.NewHTTPError(code, msg)
	}

	return ctx.JSON(http.StatusOK, res)
}
