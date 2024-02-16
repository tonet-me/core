package cardhandler

import (
	cardservice "github.com/tonet-me/tonet-core/service/card"
)

type Handler struct {
	cardSvc cardservice.Service
}

func New(cardSvc cardservice.Service) Handler {
	return Handler{cardSvc: cardSvc}
}
