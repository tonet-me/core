package cardhandler

import (
	"github.com/labstack/echo/v4"
	cardparam "github.com/tonet-me/tonet-core/param/card"
	"github.com/tonet-me/tonet-core/pkg/claim"
	errmsg "github.com/tonet-me/tonet-core/pkg/err_msg"
	httpmsg "github.com/tonet-me/tonet-core/pkg/http_msg"
	"net/http"
)

func (h Handler) createNewCard(ctx echo.Context) error {
	req := cardparam.CreateNewRequest{}
	bErr := ctx.Bind(&req)
	if bErr != nil {
		return ctx.JSON(http.StatusBadRequest, errmsg.ErrorMsgInvalidJson)
	}

	claims := claim.GetClaimsFromEchoContext(ctx)
	req.AuthenticatedUserID = claims.UserID
	if fieldErrors, err := h.cardVld.CreateRequest(req); err != nil {
		msg, code := httpmsg.Error(err)

		return ctx.JSON(code, echo.Map{
			"message": msg,
			"errors":  fieldErrors,
		})
	}

	res, cErr := h.cardSvc.CreateNew(ctx.Request().Context(), req)
	if cErr != nil {
		msg, code := httpmsg.Error(cErr)

		return echo.NewHTTPError(code, msg)
	}

	return ctx.JSON(http.StatusOK, res)
}
