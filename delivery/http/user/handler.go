package userhandler

import (
	"github.com/tonet-me/tonet-core/service/auth"
	userservice "github.com/tonet-me/tonet-core/service/user"
)

type Handler struct {
	userSvc    userservice.Service
	authSvc    auth.Service
	authConfig auth.Config
}

func New(userSvc userservice.Service) Handler {
	return Handler{userSvc: userSvc}
}
