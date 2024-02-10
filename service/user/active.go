package userservice

import (
	"context"
	userparam "github.com/tonet-me/tonet-core/param/user"
)

func (s Service) Active(ctx context.Context, req userparam.ActiveRequest) (*userparam.ActiveResponse, error) {
	success, dErr := s.repo.ActiveUser(ctx, req.AuthenticatedUserID)
	if dErr != nil {
		return nil, dErr
	}
	return &userparam.ActiveResponse{Success: success}, nil
}
