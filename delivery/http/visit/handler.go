package visithandler

import (
	visitservice "github.com/tonet-me/tonet-core/service/visit"
)

type Handler struct {
	visitSvc visitservice.Service
}

func New(visitSvc visitservice.Service) Handler {
	return Handler{
		visitSvc: visitSvc,
	}
}
