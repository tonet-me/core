package userservice

import (
	"context"
	userparam "github.com/tonet-me/tonet-core/param/user"
)

func (s Service) GetInfo(ctx context.Context, req userparam.GetInfoRequest) (*userparam.GetInfoResponse, error) {
	user, gErr := s.repo.GetUserByID(ctx, req.AuthenticatedUserID)
	if gErr != nil {
		return nil, gErr
	}

	return &userparam.GetInfoResponse{User: user}, nil
}
