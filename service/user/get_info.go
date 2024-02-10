package userservice

import userparam "tonet-core/param/user"

func (s Service) GetInfo(req userparam.GetInfoRequest) (*userparam.GetInfoResponse, error) {
	user, gErr := s.repo.GetUserByID(req.AuthenticatedUserID)
	if gErr != nil {
		return nil, gErr
	}

	return &userparam.GetInfoResponse{User: user}, nil
}
