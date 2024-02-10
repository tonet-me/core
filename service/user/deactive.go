package userservice

import (
	"context"
	userparam "github.com/tonet-me/tonet-core/param/user"
)

func (s Service) DeActive(ctx context.Context, req userparam.DeActiveRequest) (*userparam.DeActiveResponse, error) {
	success, dErr := s.repo.DeActiveUser(ctx, req.AuthenticatedUserID)
	if dErr != nil {
		return nil, dErr
	}
	return &userparam.DeActiveResponse{Success: success}, nil
}
