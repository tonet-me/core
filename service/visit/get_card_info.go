package visitservice

import (
	"context"
	cardparam "github.com/tonet-me/tonet-core/param/card"
	visitparam "github.com/tonet-me/tonet-core/param/visit"
	richerror "github.com/tonet-me/tonet-core/pkg/rich_error"
)

func (s Service) getCardInfoByName(ctx context.Context, req visitparam.GetCardInfoByNameRequest) (*cardparam.GetInfoByNameResponse, error) {
	const op = richerror.OP("visitservice.GetCardInfoByName")

	res, gErr := s.cardSvc.GetOnlyActiveCardInfoByName(ctx, cardparam.GetInfoByNameRequest{Name: req.Name})
	if gErr != nil {
		return nil, richerror.New(richerror.WithOp(op),
			richerror.WithInnerError(gErr))
	}

	return res, nil
}
