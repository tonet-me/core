package userservice

import (
	"context"
	userparam "github.com/tonet-me/tonet-core/param/user"
	richerror "github.com/tonet-me/tonet-core/pkg/rich_error"
)

func (s Service) Active(ctx context.Context, req userparam.ActiveRequest) (*userparam.ActiveResponse, error) {
	const op = richerror.OP("userservice.Active")

	success, dErr := s.repo.ActiveUser(ctx, req.AuthenticatedUserID)
	if dErr != nil {
		return nil, richerror.New(richerror.WithOp(op),
			richerror.WithInnerError(dErr))
	}
	return &userparam.ActiveResponse{Success: success}, nil
}
