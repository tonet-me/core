package miniohandler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h Handler) downloadCardProfile(ctx echo.Context) error {
	fileNameFromClient := ctx.Param("id")

	res, cErr := h.client.DownloadCardProfile(ctx.Request().Context(), fileNameFromClient)
	if cErr != nil {
		return echo.NewHTTPError(http.StatusForbidden, cErr.Error())
	}

	return ctx.Stream(http.StatusOK, "image/jpeg", res)
}
