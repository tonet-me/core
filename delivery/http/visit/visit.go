package visithandler

import (
	"github.com/labstack/echo/v4"
	cardparam "github.com/tonet-me/tonet-core/param/card"
	httpmsg "github.com/tonet-me/tonet-core/pkg/http_msg"
	"net/http"
	"strings"
)

func (h Handler) visit(ctx echo.Context) error {
	cardNameFromClient := ctx.Param("name")
	cardName := strings.Split(cardNameFromClient, `/`)[0]

	//userAgent:=ctx.Request().UserAgent() // todo: update to map[string]string from other pkg
	var req = cardparam.GetInfoByNameRequest{
		Name: cardName,
	}

	res, gErr := h.visitSvc.GetCardInfoByName(ctx.Request().Context(), req)
	if gErr != nil {
		msg, code := httpmsg.Error(gErr)

		return echo.NewHTTPError(code, msg)
	}

	h.visitSvc.AddNewVisitToCard(ctx.Request().Context(), res.Card.ID, nil) // todo: pass user agent

	return ctx.JSON(http.StatusOK, echo.Map{"file_name": cardName})
}
