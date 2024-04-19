package miniohandler

import (
	"github.com/labstack/echo/v4"
	"github.com/tonet-me/tonet-core/pkg/claim"
	"net/http"
)

func (h Handler) uploadCardProfile(ctx echo.Context) error {
	fileFromClient, fErr := ctx.FormFile("card-profile-photo")
	if fErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid file parameter")
	}

	claims := claim.GetClaimsFromEchoContext(ctx)

	bufferFile, oErr := fileFromClient.Open()
	fileSize := fileFromClient.Size
	if oErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid file")
	}
	defer bufferFile.Close()

	res, cErr := h.client.UploadCardProfile(ctx.Request().Context(), claims.UserID, &bufferFile, fileSize)
	if cErr != nil {
		return echo.NewHTTPError(http.StatusForbidden, cErr.Error())
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{"file-name": res})
}
