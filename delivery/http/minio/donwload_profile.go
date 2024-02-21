package miniohandler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h Handler) downloadUserProfile(ctx echo.Context) error {
	fileNameFromClient := ctx.Param("id")

	res, cErr := h.client.DownloadUserProfilePhoto(ctx.Request().Context(), fileNameFromClient)
	if cErr != nil {
		return ctx.JSON(http.StatusForbidden, cErr.Error())
	}

	return ctx.Stream(http.StatusOK, "image/jpeg", res)
}
