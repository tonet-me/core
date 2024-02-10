package userservice

import userparam "tonet-core/param/user"

func (s Service) DeActive(req userparam.DeActiveRequest) (*userparam.DeActiveResponse, error) {
	success, dErr := s.repo.DeActiveUser(req.AuthenticatedUserID)
	if dErr != nil {
		return nil, dErr
	}
	return &userparam.DeActiveResponse{Success: success}, nil
}
