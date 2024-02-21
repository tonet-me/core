package userservice

import (
	"context"
	userparam "github.com/tonet-me/tonet-core/param/user"
	richerror "github.com/tonet-me/tonet-core/pkg/rich_error"
)

func (s Service) GetInfo(ctx context.Context, req userparam.GetInfoRequest) (*userparam.GetInfoResponse, error) {
	const op = richerror.OP("userservice.GetInfo")

	user, gErr := s.repo.GetUserByID(ctx, req.AuthenticatedUserID)
	if gErr != nil {
		return nil, richerror.New(richerror.WithOp(op),
			richerror.WithInnerError(gErr))
	}

	return &userparam.GetInfoResponse{User: user}, nil
}
