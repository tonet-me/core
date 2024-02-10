package userservice

import userparam "tonet-core/param/user"

func (s Service) Active(req userparam.ActiveRequest) (*userparam.ActiveResponse, error) {
	success, dErr := s.repo.ActiveUser(req.AuthenticatedUserID)
	if dErr != nil {
		return nil, dErr
	}
	return &userparam.ActiveResponse{Success: success}, nil
}
