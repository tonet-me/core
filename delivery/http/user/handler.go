package userhandler

import (
	"github.com/tonet-me/tonet-core/service/auth"
	userservice "github.com/tonet-me/tonet-core/service/user"
	uservalidator "github.com/tonet-me/tonet-core/validator/user"
)

type Handler struct {
	userSvc    userservice.Service
	userVld    uservalidator.Validator
	authSvc    auth.Service
	authConfig auth.Config
}

func New(userSvc userservice.Service, userVld uservalidator.Validator, authSvc auth.Service, authCfg auth.Config) Handler {
	return Handler{
		userSvc:    userSvc,
		userVld:    userVld,
		authSvc:    authSvc,
		authConfig: authCfg,
	}
}
