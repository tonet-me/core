package cardservice

import (
	"context"
	cardparam "github.com/tonet-me/tonet-core/param/card"
	userparam "github.com/tonet-me/tonet-core/param/user"
	richerror "github.com/tonet-me/tonet-core/pkg/rich_error"
)

func (s Service) Active(ctx context.Context, req cardparam.ActiveRequest) (*userparam.ActiveResponse, error) {
	const op = richerror.OP("cardservice.Active")

	success, dErr := s.repo.ActiveCard(ctx, req.AuthenticatedUserID)
	if dErr != nil {
		return nil, richerror.New(richerror.WithOp(op),
			richerror.WithInnerError(dErr))
	}
	return &userparam.ActiveResponse{Success: success}, nil
}
