package cardservice

import (
	"context"
	cardparam "github.com/tonet-me/tonet-core/param/card"
	richerror "github.com/tonet-me/tonet-core/pkg/rich_error"
)

func (s Service) GetAllUserCards(ctx context.Context, req cardparam.GetAllUserCardsRequest) (*cardparam.GetAllUserCardsResponse, error) {
	const op = richerror.OP("cardservice.GetAllUserCards")

	cards, gErr := s.repo.GetAllCardsByUserID(ctx, req.AuthenticatedUserID)
	if gErr != nil {
		return nil, richerror.New(richerror.WithOp(op),
			richerror.WithInnerError(gErr))
	}

	return &cardparam.GetAllUserCardsResponse{Cards: cards}, nil
}
