package cardservice

import (
	"context"
	cardparam "github.com/tonet-me/tonet-core/param/card"
	userparam "github.com/tonet-me/tonet-core/param/user"
	richerror "github.com/tonet-me/tonet-core/pkg/rich_error"
)

func (s Service) DeActive(ctx context.Context, req cardparam.DeActiveRequest) (*userparam.DeActiveResponse, error) {
	const op = richerror.OP("cardservice.DeActive")

	success, dErr := s.repo.DeActiveCard(ctx, req.AuthenticatedUserID)
	if dErr != nil {
		return nil, richerror.New(richerror.WithOp(op),
			richerror.WithInnerError(dErr))
	}
	return &userparam.DeActiveResponse{Success: success}, nil
}
