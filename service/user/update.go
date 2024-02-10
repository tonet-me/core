package userservice

import (
	"context"
	"github.com/tonet-me/tonet-core/entity"
	userparam "github.com/tonet-me/tonet-core/param/user"
)

func (s Service) Update(ctx context.Context, req userparam.UpdateRequest) (*userparam.UpdateResponse, error) {
	user, gErr := s.repo.GetUserByID(ctx, req.AuthenticatedUserID)
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
	updatedUser, uErr := s.repo.UpdateUser(ctx, req.AuthenticatedUserID, userDataUpdate)
	if uErr != nil {
		return nil, uErr
	}

	return &userparam.UpdateResponse{User: updatedUser}, nil
}
