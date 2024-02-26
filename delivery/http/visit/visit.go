package visithandler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

func (h Handler) visit(ctx echo.Context) error {
	cardNameFromClient := ctx.Param("name")
	cardName := strings.Split(cardNameFromClient, `/`)[0]

	return ctx.JSON(http.StatusOK, echo.Map{"file_name": cardName})
}
