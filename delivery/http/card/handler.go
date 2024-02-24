package cardhandler

import (
	"github.com/tonet-me/tonet-core/service/auth"
	cardservice "github.com/tonet-me/tonet-core/service/card"
)

type Handler struct {
	cardSvc    cardservice.Service
	authSvc    auth.Service
	authConfig auth.Config
}

func New(cardSvc cardservice.Service, authSvc auth.Service, authCfg auth.Config) Handler {
	return Handler{
		cardSvc:    cardSvc,
		authSvc:    authSvc,
		authConfig: authCfg,
	}
}
