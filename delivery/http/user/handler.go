package userhandler

import userservice "github.com/tonet-me/tonet-core/service/user"

type Handler struct {
	userSvc userservice.Service
}

func New(userSvc userservice.Service) Handler {
	return Handler{userSvc: userSvc}
}
