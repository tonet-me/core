package visitservice

import (
	"context"
	"github.com/tonet-me/tonet-core/entity"
	cardparam "github.com/tonet-me/tonet-core/param/card"
)

type Repository interface {
	AddVisitToCard(ctx context.Context, visit entity.Visit) error
}

type CardService interface {
	GetOnlyActiveCardInfoByName(ctx context.Context, req cardparam.GetInfoByNameRequest) (*cardparam.GetInfoByNameResponse, error)
}
type Service struct {
	repo    Repository
	cardSvc CardService
}

func New(repo Repository, cardSvc CardService) Service {
	return Service{
		repo:    repo,
		cardSvc: cardSvc,
	}
}
