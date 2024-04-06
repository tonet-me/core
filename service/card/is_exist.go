package cardservice

import (
	"context"
	cardparam "github.com/tonet-me/tonet-core/param/card"
	richerror "github.com/tonet-me/tonet-core/pkg/rich_error"
)

func (s Service) IsExist(ctx context.Context, req cardparam.IsExistRequest) (*cardparam.IsExistResponse, error) {
	const op = richerror.OP("cardservice.IsExist")

	isExist, gErr := s.repo.IsCardExistByName(ctx, req.Name)
	if gErr != nil {
		return nil, richerror.New(richerror.WithOp(op),
			richerror.WithInnerError(gErr))
	}

	return &cardparam.IsExistResponse{IsExist: isExist}, nil
}
