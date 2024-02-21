package claim

import (
	"github.com/labstack/echo/v4"
	"github.com/tonet-me/tonet-core/config"
	"github.com/tonet-me/tonet-core/service/auth"
)

func GetClaimsFromEchoContext(c echo.Context) *auth.Claims {
	//defensive programming vs let it crash - log-metric-recover ,...
	return c.Get(config.AuthMiddlewareContextKey).(*auth.Claims)
}
