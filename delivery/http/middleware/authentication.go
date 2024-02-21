package middleware

import (
	mw "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/tonet-me/tonet-core/config"
	"github.com/tonet-me/tonet-core/service/auth"
)

func Authentication(service auth.Service, cfg auth.Config) echo.MiddlewareFunc {
	return mw.WithConfig(mw.Config{
		ContextKey: config.AuthMiddlewareContextKey,
		SigningKey: []byte(cfg.SignKey),
		// TODO  - as sign method string to config
		SigningMethod: "HS256",
		ParseTokenFunc: func(c echo.Context, auth string) (interface{}, error) {
			claims, err := service.ParseToken(auth)
			if err != nil {
				return nil, err
			}

			return claims, nil
		},
	})
}
