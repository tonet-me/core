package userhandler

import (
	"github.com/labstack/echo/v4"
	"github.com/tonet-me/tonet-core/entity"
	userparam "github.com/tonet-me/tonet-core/param/user"
	errmsg "github.com/tonet-me/tonet-core/pkg/err_msg"
	httpmsg "github.com/tonet-me/tonet-core/pkg/http_msg"
	"net/http"
)

func (h Handler) getTokenFromRefreshToken(ctx echo.Context) error {
	req := userparam.GetRefreshTokenRequest{}
	if bErr := ctx.Bind(&req); bErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errmsg.ErrorMsgInvalidJson)
	}

	if fieldErrors, err := h.userVld.RefreshTokenRequest(req); err != nil {
		msg, code := httpmsg.Error(err)

		return echo.NewHTTPError(code, echo.Map{
			"message": msg,
			"errors":  fieldErrors,
		})
	}

	claims, pErr := h.authSvc.ParseToken(req.RefreshToken)
	if pErr != nil {
		msg, code := httpmsg.Error(pErr)

		return echo.NewHTTPError(code, msg)
	}

	if claims.Subject != h.authConfig.RefreshSubject {
		return echo.NewHTTPError(http.StatusBadRequest, errmsg.ErrorMsgInvalidRefreshToken)
	}

	authenticate := entity.Authenticable{
		ID: claims.UserID,
	}

	res, gErr := h.userSvc.GenerateTokens(authenticate)
	if gErr != nil {
		msg, code := httpmsg.Error(pErr)

		return echo.NewHTTPError(code, msg)
	}

	return ctx.JSON(http.StatusOK, res)

}
