package miniohandler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h Handler) uploadUserProfile(ctx echo.Context) error {

	//createCardParam := cardparam.CreateNewRequest{}
	//bErr := ctx.Bind(&createCardParam)
	//fmt.Println("req body", createCardParam)
	//if bErr != nil {
	//	return ctx.JSON(http.StatusBadRequest, "invalid json format")
	//}

	fileFromClient, fErr := ctx.FormFile("profile-photo")
	if fErr != nil {
		return ctx.JSON(http.StatusBadRequest, "invalid file parameter")
	}

	// Get Buffer from file
	bufferFile, oErr := fileFromClient.Open()
	fileSize := fileFromClient.Size
	if oErr != nil {
		return ctx.JSON(http.StatusBadRequest, "invalid file")
	}
	defer bufferFile.Close()

	res, cErr := h.client.UploadUserProfilePhoto(ctx.Request().Context(), "test", &bufferFile, fileSize)
	if cErr != nil {
		return ctx.JSON(http.StatusForbidden, cErr.Error())
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{"file-name": res})
}
