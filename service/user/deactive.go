package userservice

import (
	"context"
	userparam "github.com/tonet-me/tonet-core/param/user"
	richerror "github.com/tonet-me/tonet-core/pkg/rich_error"
)

func (s Service) DeActive(ctx context.Context, req userparam.DeActiveRequest) (*userparam.DeActiveResponse, error) {
	const op = richerror.OP("userservice.DeActive")

	success, dErr := s.repo.DeActiveUser(ctx, req.AuthenticatedUserID)
	if dErr != nil {
		return nil, richerror.New(richerror.WithOp(op),
			richerror.WithInnerError(dErr))
	}
	return &userparam.DeActiveResponse{Success: success}, nil
}
