package cardhandler

import (
	"fmt"
	"github.com/labstack/echo/v4"
	cardparam "github.com/tonet-me/tonet-core/param/card"
	"net/http"
)

func (h Handler) createNewCard(ctx echo.Context) error {

	createCardParam := cardparam.CreateNewRequest{}
	bErr := ctx.Bind(&createCardParam)
	fmt.Println("req body", createCardParam)
	if bErr != nil {
		return ctx.JSON(http.StatusBadRequest, "invalid json format")
	}
	res, cErr := h.cardSvc.CreateNew(ctx.Request().Context(), createCardParam)
	if cErr != nil {
		return ctx.JSON(http.StatusForbidden, cErr.Error())
	}
	return ctx.JSON(http.StatusOK, res)
}
