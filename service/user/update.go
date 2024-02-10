package userservice

import (
	"tonet-core/entity"
	userparam "tonet-core/param/user"
)

func (s Service) Update(req userparam.UpdateRequest) (*userparam.UpdateResponse, error) {
	user, gErr := s.repo.GetUserByID(req.AuthenticatedUserID)
	if gErr != nil {
		return nil, gErr
	}

	userDataUpdate := entity.User{
		FirstName:       req.UpdateData.FirstName,
		LastName:        req.UpdateData.LastName,
		Email:           user.Email,
		PhoneNumber:     req.UpdateData.PhoneNumber,
		ProfilePhotoURL: req.UpdateData.ProfilePhotoURL,
		Status:          user.Status,
	}
	updatedUser, uErr := s.repo.UpdateUser(userDataUpdate)
	if uErr != nil {
		return nil, uErr
	}

	return &userparam.UpdateResponse{User: updatedUser}, nil
}
