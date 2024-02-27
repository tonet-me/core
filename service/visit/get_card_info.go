package visitservice

import (
	"context"
	cardparam "github.com/tonet-me/tonet-core/param/card"
	richerror "github.com/tonet-me/tonet-core/pkg/rich_error"
)

func (s Service) GetCardInfoByName(ctx context.Context, req cardparam.GetInfoByNameRequest) (*cardparam.GetInfoByNameResponse, error) {
	const op = richerror.OP("visitservice.GetCardInfoByName")

	res, gErr := s.cardSvc.GetInfoByName(ctx, req)
	if gErr != nil {
		return nil, richerror.New(richerror.WithOp(op),
			richerror.WithInnerError(gErr))
	}

	return res, nil
}
