package miniohandler

import (
	"github.com/labstack/echo/v4"
	fileparam "github.com/tonet-me/tonet-core/param/file"
	errmsg "github.com/tonet-me/tonet-core/pkg/err_msg"
	httpmsg "github.com/tonet-me/tonet-core/pkg/http_msg"
	"net/http"
)

func (h Handler) uploadCardProfile(ctx echo.Context) error {
	req := fileparam.CardProfileUploadRequest{}
	bErr := ctx.Bind(&req)
	if bErr != nil {
		return ctx.JSON(http.StatusBadRequest, errmsg.ErrorMsgInvalidJson)
	}

	if fieldErrors, err := h.fileVld.UploadCardProfileRequest(req); err != nil {
		msg, code := httpmsg.Error(err)

		return echo.NewHTTPError(code, echo.Map{
			"message": msg,
			"errors":  fieldErrors,
		})
	}

	fileFromClient, fErr := ctx.FormFile("card-profile-photo")
	if fErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid file parameter")
	}

	bufferFile, oErr := fileFromClient.Open()
	fileSize := fileFromClient.Size
	if oErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid file")
	}
	defer bufferFile.Close()

	res, cErr := h.client.UploadCardProfile(ctx.Request().Context(), req.CardID, &bufferFile, fileSize)
	if cErr != nil {
		return echo.NewHTTPError(http.StatusForbidden, cErr.Error())
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{"file-name": res})
}
