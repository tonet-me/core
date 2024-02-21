package userservice

import (
	"context"
	"github.com/tonet-me/tonet-core/entity"
	userparam "github.com/tonet-me/tonet-core/param/user"
	richerror "github.com/tonet-me/tonet-core/pkg/rich_error"
)

func (s Service) Update(ctx context.Context, req userparam.UpdateRequest) (*userparam.UpdateResponse, error) {
	const op = richerror.OP("userservice.Update")

	user, gErr := s.repo.GetUserByID(ctx, req.AuthenticatedUserID)
	if gErr != nil {
		return nil, richerror.New(richerror.WithOp(op),
			richerror.WithInnerError(gErr),
		)
	}

	userDataUpdate := entity.User{
		FirstName:       req.UpdateData.FirstName,
		LastName:        req.UpdateData.LastName,
		Email:           user.Email,
		PhoneNumber:     req.UpdateData.PhoneNumber,
		ProfilePhotoURL: req.UpdateData.ProfilePhotoURL,
		Status:          user.Status,
	}
	updated, uErr := s.repo.UpdateUser(ctx, req.AuthenticatedUserID, userDataUpdate)
	if uErr != nil {
		return nil, richerror.New(richerror.WithOp(op),
			richerror.WithInnerError(uErr),
		)
	}

	return &userparam.UpdateResponse{Updated: updated}, nil
}
