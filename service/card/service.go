package cardservice

import (
	"context"
	"github.com/tonet-me/tonet-core/entity"
)

type Repository interface {
	CreateNewCard(ctx context.Context, card entity.Card) (entity.Card, error)
	UpdateCard(ctx context.Context, cardID string, card entity.Card) (bool, error)
	//ActiveCard(ctx context.Context, cardID string) (bool, error)
	//DeActiveCard(ctx context.Context, cardID string) (bool, error)
	GetCardByID(ctx context.Context, cardID string) (entity.Card, error)
	GetCardByName(ctx context.Context, name string) (entity.Card, error)
	GetAllCardsByUserID(ctx context.Context, userID string) ([]entity.Card, error)
	IsCardExistByName(ctx context.Context, name string) (bool, error)
}

type Service struct {
	repo Repository
}

func New(repo Repository) Service {
	return Service{repo: repo}
}
