package cardhandler

import (
	"github.com/tonet-me/tonet-core/service/auth"
	cardservice "github.com/tonet-me/tonet-core/service/card"
	cardvalidator "github.com/tonet-me/tonet-core/validator/card"
)

type Handler struct {
	cardSvc    cardservice.Service
	cardVld    cardvalidator.Validator
	authSvc    auth.Service
	authConfig auth.Config
}

func New(cardSvc cardservice.Service, cardVld cardvalidator.Validator, authSvc auth.Service, authCfg auth.Config) Handler {
	return Handler{
		cardSvc:    cardSvc,
		cardVld:    cardVld,
		authSvc:    authSvc,
		authConfig: authCfg,
	}
}
